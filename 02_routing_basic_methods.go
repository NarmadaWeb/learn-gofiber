package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("Mengambil daftar pengguna (GET)")
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		return c.SendString("Membuat pengguna baru (POST)")
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("Memperbarui pengguna dengan ID: " + id + " (PUT)")
	})

	app.Patch("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("Memperbarui sebagian data pengguna ID: " + id + " (PATCH)")
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("Menghapus pengguna dengan ID: " + id + " (DELETE)")
	})

	app.Options("/info", func(c *fiber.Ctx) error {
		c.Set("Allow", "GET, POST, OPTIONS")
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Head("/status", func(c *fiber.Ctx) error {
		c.Set("X-App-Status", "OK")
		return c.SendStatus(fiber.StatusOK)
	})

	app.All("/universal", func(c *fiber.Ctx) error {
		return c.SendString("Endpoint ini merespons semua metode HTTP.")
	})

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
