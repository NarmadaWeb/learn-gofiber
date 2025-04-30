package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/echo", func(c *fiber.Ctx) error {
		contentType := c.Get(fiber.HeaderContentType)
		log.Printf("Echo endpoint received Content-Type: %s", contentType)

		body := c.BodyRaw()

		c.Set(fiber.HeaderContentType, contentType)
		return c.Status(fiber.StatusOK).Send(body)
	})

	app.Get("/users/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.JSON(fiber.Map{"user": name})
	})
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	log.Println("Starting basic server for testing on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
