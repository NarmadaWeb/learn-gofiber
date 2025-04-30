package main

import (
	"log"

	"github.com/gofiber/contrib/websocket" // Import websocket
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Middleware untuk memastikan request adalah upgrade WebSocket
	// Hanya berlaku untuk path yang dimulai dengan /ws
	app.Use("/ws", func(c *fiber.Ctx) error {
		// Periksa apakah header menunjukkan permintaan upgrade WebSocket
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true) // Tandai bahwa upgrade diizinkan
			log.Println("WebSocket upgrade request detected")
			return c.Next() // Lanjutkan ke handler WebSocket
		}
		log.Println("Request is not a WebSocket upgrade")
		// Jika bukan permintaan WebSocket, kirim 426 Upgrade Required
		return fiber.ErrUpgradeRequired
	})

	// Handler untuk koneksi WebSocket di path /ws/:id
	// Handler ini hanya akan dipanggil jika middleware di atas memanggil c.Next()
	app.Get("/ws/:id", websocket.New(func(conn *websocket.Conn) {
		// conn adalah *websocket.Conn yang membungkus koneksi

		// Dapatkan parameter dari URL asli (sebelum upgrade)
		// Data dari context HTTP asli masih bisa diakses di sini
		id := conn.Params("id")
		queryParam := conn.Query("token")
		isAllowed := conn.Locals("allowed") // Ambil data dari Locals middleware sebelumnya

		log.Printf("WebSocket connected for ID: %s from %s", id, conn.RemoteAddr())
		log.Printf("Query Param 'token': %s", queryParam)
		log.Printf("Is Upgrade Allowed (from Locals): %v", isAllowed)


		// Variabel untuk tipe pesan, pesan, dan error
		var (
			mt  int // messageType
			msg []byte
			err error
		)

		// Loop tak terbatas untuk membaca pesan dari client
		for {
			// Baca pesan dari client
			// conn.ReadMessage() adalah blocking call
			if mt, msg, err = conn.ReadMessage(); err != nil {
				// Jika ada error (koneksi ditutup, dll.), log dan keluar dari loop
				// Cek jenis error untuk logging yang lebih baik
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Error reading message (ID: %s): %v", id, err)
				} else {
					log.Printf("WebSocket connection closed normally (ID: %s): %v", id, err)
				}
				break // Keluar dari loop, menutup koneksi di sisi server
			}

			log.Printf("Message received from ID %s: %s (Type: %d)", id, msg, mt)

			// Kirim pesan kembali ke client (echo)
			// Anda bisa mengirim pesan teks (websocket.TextMessage) atau biner (websocket.BinaryMessage)
			if err = conn.WriteMessage(mt, msg); err != nil {
				log.Println("Error writing message:", err)
				break // Keluar jika gagal menulis
			}
			log.Printf("Echoed message back to ID %s", id)

			// Contoh kirim JSON
			// response := fiber.Map{"received": string(msg), "echoed": true, "clientId": id}
			// if err = conn.WriteJSON(response); err != nil {
			//     log.Println("Error writing JSON:", err)
			//     break
			// }
		}
		// Kode setelah loop akan dieksekusi saat koneksi ditutup (baik oleh client atau server)
		log.Printf("WebSocket disconnected for ID: %s", id)
		// Lakukan pembersihan jika perlu (misalnya, hapus user dari daftar online)
	}))

	// Handler biasa untuk menunjukkan middleware /ws tidak memblokir non-websocket request
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ini Halaman Biasa (Bukan WebSocket)")
	})


	log.Println("Starting server on port 3000...")
	log.Println("WebSocket endpoint available at ws://localhost:3000/ws/<some_id>?token=<your_token>")
	log.Fatal(app.Listen(":3000"))
}
