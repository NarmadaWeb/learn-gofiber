package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/order/:id?", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.SendString("Menampilkan semua order")
		}
		return c.SendString("Menampilkan detail order ID: " + id)
	})

	app.Get("/files/*", func(c *fiber.Ctx) error {
		filePath := c.Params("*")
		return c.SendString("Mengakses file di path: " + filePath)
	})

	app.Get("/user/+", func(c *fiber.Ctx) error {
		name := c.Params("+")
		return c.SendString("Profil pengguna: " + name)
	})

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
