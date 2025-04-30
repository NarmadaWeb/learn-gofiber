package main

import (
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp" // Import fasthttp untuk Stream
	"context" // Untuk Done()
)

func sseHandler(c *fiber.Ctx) error {
	// 1. Set header yang diperlukan untuk SSE
	c.Set(fiber.HeaderContentType, "text/event-stream")
	c.Set(fiber.HeaderCacheControl, "no-cache")
	c.Set(fiber.HeaderConnection, "keep-alive")
	// Transfer-Encoding: chunked biasanya otomatis ditangani oleh fasthttp/fiber
	// c.Set(fiber.HeaderTransferEncoding, "chunked")

	clientIP := c.IP() // Dapatkan IP client untuk logging
	log.Printf("SSE Client connected: %s", clientIP)

	// 2. Gunakan c.Context().SetBodyStreamWriter untuk streaming response
	// Ini memungkinkan kita menulis ke response secara bertahap.
	// Fungsi ini akan dijalankan dalam goroutine oleh fasthttp.
	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		defer log.Printf("SSE Stream writer finished for client: %s", clientIP)

		eventID := 0
		ticker := time.NewTicker(2 * time.Second) // Kirim event setiap 2 detik
		defer ticker.Stop() // Pastikan ticker berhenti saat fungsi selesai

		// Kirim komentar :ok untuk membuka koneksi (beberapa client membutuhkannya)
		fmt.Fprintln(w, ": ok")
		if err := w.Flush(); err != nil {
			log.Printf("SSE Error initial flush for %s: %v", clientIP, err)
			return
		}

		// Loop untuk mengirim event secara berkala
		for {
			select {
			// Cek apakah koneksi client masih aktif (context dibatalkan)
			case <-c.Context().Done(): // Context bawaan Fiber/Fasthttp
				log.Printf("SSE Client disconnected (context done): %s", clientIP)
				return // Hentikan loop jika client disconnect

			// Tunggu interval waktu tertentu dari ticker
			case <-ticker.C:
				eventID++
				now := time.Now()
				log.Printf("SSE Sending event %d to %s", eventID, clientIP)

				// Format dan kirim pesan SSE:
				// id: <unique_id>\n
				// event: <event_name>\n (opsional)
				// data: <your_data>\n\n (data bisa multiline jika diawali 'data: ')

				// Kirim event 'server-time' dengan data waktu saat ini (JSON)
				fmt.Fprintf(w, "id: %d\n", eventID)
				fmt.Fprintf(w, "event: server-time\n")
				// Pastikan data JSON valid
				jsonData := fmt.Sprintf("{\"time\": \"%s\", \"clientID\": \"%s\"}", now.Format(time.RFC3339), clientIP)
				fmt.Fprintf(w, "data: %s\n\n", jsonData)

				// Kirim event lain (contoh) setiap 5 detik
				if eventID%5 == 0 {
					fmt.Fprintf(w, "id: %d-ping\n", eventID)
					fmt.Fprintf(w, "event: ping\n")
					fmt.Fprintf(w, "data: Pong from server\n\n") // Data bisa string biasa
					// Kirim komentar (diabaikan oleh EventSource tapi bisa untuk debug)
					// fmt.Fprintln(w, ": This is a server comment")
				}

				// Flush buffer untuk memastikan data dikirim ke client
				if err := w.Flush(); err != nil {
					// Error flushing biasanya berarti client disconnect
					log.Printf("SSE Error flushing for %s: %v", clientIP, err)
					// Hentikan goroutine jika tidak bisa flush
					return
				}
				// log.Printf("SSE event %d flushed to %s", eventID, clientIP)
			}
		}
	})

	// Penting: Return nil di sini karena response ditulis oleh stream writer
	// Fiber/Fasthttp akan menangani penyelesaian response setelah stream writer selesai.
	// Jangan coba return c.SendStatus atau lainnya di sini.
	log.Printf("SSE Handler returning for %s (stream writer started)", clientIP)
	return nil
}

func main() {
	app := fiber.New()

	app.Get("/events", sseHandler)

	// Halaman HTML sederhana untuk menguji SSE dari browser
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(`
			<!DOCTYPE html>
			<html>
			<head><title>SSE Test</title>
			<style>
				body { font-family: sans-serif; }
				#events { list-style: none; padding-left: 0; }
				#events li { margin-bottom: 5px; padding: 8px; border: 1px solid #eee; border-radius: 4px; }
				#events li.ping { background-color: #e0ffe0; }
				#events li.error { background-color: #ffe0e0; color: red; }
			</style>
			</head>
			<body>
			<h1>Server-Sent Events Test</h1>
			<p>Status: <span id="status">Connecting...</span></p>
			<ul id="events"></ul>
			<script>
				const eventsList = document.getElementById('events');
				const statusElem = document.getElementById('status');
				let evtSource;

				function connectSSE() {
					if (evtSource) {
						evtSource.close();
					}
					statusElem.textContent = "Connecting...";
					eventsList.innerHTML = '<li>Attempting to connect to /events</li>';
					evtSource = new EventSource("/events"); // Hubungkan ke endpoint SSE

					evtSource.onopen = function() {
						console.log("SSE Connection opened.");
						statusElem.textContent = "Connected";
						const newItem = document.createElement("li");
						newItem.textContent = "Connection established!";
						newItem.style.backgroundColor = '#e0e0ff';
						eventsList.appendChild(newItem);
					};


					// Handler untuk event bernama 'server-time'
					evtSource.addEventListener("server-time", function(event) {
						console.log("Received server-time event:", event.data);
						try {
							const data = JSON.parse(event.data);
							const newItem = document.createElement("li");
							newItem.textContent = "Server Time: " + data.time + " (ID: " + event.lastEventId + ", Client: " + data.clientID + ")";
							eventsList.appendChild(newItem);
						} catch (e) {
							console.error("Failed to parse server-time data:", e, event.data);
						}
					});

					 // Handler untuk event bernama 'ping'
					evtSource.addEventListener("ping", function(event) {
						console.log("Received ping event (ID: " + event.lastEventId + ")");
						const newItem = document.createElement("li");
						newItem.classList.add('ping');
						newItem.textContent = "PING! Data: " + event.data + " (ID: " + event.lastEventId + ")";
						eventsList.appendChild(newItem);
					});

					// Handler untuk event 'message' default (jika 'event:' tidak diset)
					evtSource.onmessage = function(event) {
						console.log("Received generic message:", event.data);
						 const newItem = document.createElement("li");
						newItem.textContent = "Generic Message: " + event.data;
						eventsList.appendChild(newItem);
					};

					// Handler untuk error koneksi
					evtSource.onerror = function(err) {
						console.error("EventSource failed:", err);
						statusElem.textContent = "Connection Error (Retrying...)";
						const newItem = document.createElement("li");
						newItem.textContent = "Connection error! Browser might attempt to reconnect.";
						newItem.classList.add('error');
						eventsList.appendChild(newItem);
						// Browser akan mencoba reconnect otomatis by default
						// Jika status readyState adalah CONNECTING, biarkan saja
						if (evtSource.readyState === EventSource.CLOSED) {
							 statusElem.textContent = "Connection Closed Permanently";
							 // Mungkin coba reconnect manual setelah delay
							 // setTimeout(connectSSE, 5000);
						}
					};
					console.log("Connecting to SSE stream...");
				} // end connectSSE

				connectSSE(); // Panggil saat halaman dimuat
			</script>
			</body>
			</html>
		`)
	})

	log.Println("Starting server on port 3000...")
	log.Println("SSE endpoint: http://localhost:3000/events")
	log.Println("HTML test page: http://localhost:3000/")
	log.Fatal(app.Listen(":3000"))
}
