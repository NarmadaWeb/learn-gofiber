package main

import (
	"log"
	"sync" // Diperlukan untuk Mutex

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Struct untuk menyimpan koneksi client
type client struct {
	isClosing bool
	mu        sync.Mutex
}

// Map untuk menyimpan semua koneksi client yang aktif
var clients = make(map[*websocket.Conn]*client)
// Mutex untuk melindungi akses ke map 'clients' (read/write)
var clientsMu sync.RWMutex
// Channel untuk pesan broadcast
var broadcast = make(chan []byte)

// Hub function untuk mengelola koneksi dan broadcast
func runHub() {
	log.Println("Starting WebSocket Hub...")
	for {
		// Tunggu pesan baru di channel broadcast
		msg := <-broadcast
		log.Printf("Hub: Broadcasting message: %s", string(msg))

		// Lock map clients untuk iterasi (read lock)
		clientsMu.RLock()
		// Buat daftar koneksi yang akan dihapus (jika error)
		deadConnections := []*websocket.Conn{}

		// Iterasi semua client yang terhubung
		for conn, c := range clients {
			// Lock client individual untuk memeriksa status isClosing
			c.mu.Lock()
			isClosing := c.isClosing
			c.mu.Unlock()

			if !isClosing {
				// Kirim pesan ke client ini
				// Gunakan goroutine agar pengiriman ke satu client lambat/error
				// tidak memblokir pengiriman ke client lain
				go func(conn *websocket.Conn, msg []byte) {
					 if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
						log.Printf("Hub: Error writing to client %s: %v", conn.RemoteAddr(), err)
						// Tandai koneksi ini untuk dihapus
						clientsMu.Lock() // Lock map untuk menulis
						// Tandai client sebagai closing untuk mencegah write lebih lanjut
						if cl, ok := clients[conn]; ok {
							cl.mu.Lock()
							cl.isClosing = true
							cl.mu.Unlock()
						}
						// Tambahkan ke daftar penghapusan (dilakukan setelah RLock dilepas)
						// (Alternatif: kirim ke channel unregister)
						deadConnections = append(deadConnections, conn)
						clientsMu.Unlock() // Unlock map setelah modifikasi
					}
				}(conn, msg)
			}
		}
		clientsMu.RUnlock() // Lepas read lock setelah loop

		// Hapus koneksi yang error (jika ada)
		if len(deadConnections) > 0 {
			clientsMu.Lock() // Lock map untuk menulis (delete)
			for _, conn := range deadConnections {
				log.Printf("Hub: Removing dead connection: %s", conn.RemoteAddr())
				conn.Close() // Tutup koneksi
				delete(clients, conn) // Hapus dari map
			}
			clientsMu.Unlock()
		}
	}
}


func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Jalankan hub di goroutine terpisah
	go runHub()

	// Middleware upgrade WebSocket
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Handler WebSocket
	app.Get("/ws", websocket.New(func(conn *websocket.Conn) {
		remoteAddr := conn.RemoteAddr().String()
		log.Printf("WebSocket connected: %s", remoteAddr)

		// 1. Buat struct client baru dan daftarkan koneksi
		c := &client{}
		clientsMu.Lock()
		clients[conn] = c
		clientsMu.Unlock()
		log.Printf("Client %s registered. Total clients: %d", remoteAddr, len(clients))


		// 2. Pastikan koneksi dihapus saat handler selesai
		defer func() {
			log.Printf("Defer: Unregistering client %s", remoteAddr)
			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()
			conn.Close() // Pastikan koneksi ditutup
			log.Printf("Client %s unregistered and closed. Total clients: %d", remoteAddr, len(clients))
			// Kirim pesan bahwa user keluar (opsional)
			// broadcast <- []byte(fmt.Sprintf("User %s left.", remoteAddr))
		}()


		// 3. Kirim pesan selamat datang ke client yang baru connect
		if err := conn.WriteMessage(websocket.TextMessage, []byte("Selamat datang di Chat!")); err != nil {
			log.Printf("Error writing welcome message to %s: %v", remoteAddr, err)
			return // Keluar jika gagal kirim pesan pertama
		}
		// Kirim pesan ke semua bahwa user baru masuk (opsional)
		// broadcast <- []byte(fmt.Sprintf("User %s joined.", remoteAddr))


		// 4. Loop membaca pesan dari client ini
		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Read error from %s: %v", remoteAddr, err)
				} else {
					log.Printf("Connection closed by %s", remoteAddr)
				}
				break // Keluar loop jika error/tutup
			}

			if mt == websocket.TextMessage {
				// Kirim pesan yang diterima ke channel broadcast
				log.Printf("Message from %s: %s", remoteAddr, string(msg))
				// Tambahkan prefix nama pengirim (contoh)
				// broadcastMsg := fmt.Sprintf("[%s]: %s", remoteAddr, string(msg))
				broadcast <- msg // Kirim pesan asli ke hub
			} else {
				log.Printf("Received non-text message type %d from %s", mt, remoteAddr)
			}
		}
	}))


	log.Println("Starting server on port 3000...")
	log.Println("WebSocket endpoint: ws://localhost:3000/ws")
	log.Fatal(app.Listen(":3000"))
}
