package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func HandleRequest(c *fiber.Ctx) error {
	userAgent := c.Get(fiber.HeaderUserAgent)
	apiKey := c.Get("X-API-Key", "default-key")
	acceptHeader := c.Get("Accept")

	log.Printf("User-Agent: %s", userAgent)
	log.Printf("API Key: %s", apiKey)
	log.Printf("Accept: %s", acceptHeader)

	bestAccept := c.Accepts("application/json", "text/html", "application/xml")
	log.Printf("Client best accepts: %s", bestAccept)

	switch bestAccept {
	case "application/json":
		return c.JSON(fiber.Map{"message": "Anda meminta JSON"})
	case "text/html":
		return c.Type("html").SendString("<h1>Anda meminta HTML</h1>")
	case "application/xml":
		return c.Type("xml").SendString("<response><message>Anda meminta XML</message></response>")
	default:
		if acceptHeader != "" && bestAccept == "" {
			log.Println("Client Accept header tidak cocok dengan yang ditawarkan")
			return c.Status(fiber.StatusNotAcceptable).SendString("Tipe konten yang diminta tidak didukung")
		}
		log.Println("Tidak ada Accept header spesifik atau cocok, mengirim default text")
		return c.SendString("Header diterima, mengirim response default (text/plain)")
	}
}

func main() {
	app := fiber.New()

	app.Get("/headers", HandleRequest)

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
