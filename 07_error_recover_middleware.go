package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})

	app.Get("/panic-now", func(c *fiber.Ctx) error {
		var myMap map[string]string
		log.Println("Trying to access nil map...")
		myMap["key"] = "value"
		return c.SendString("Tidak akan sampai sini")
	})

	app.Get("/panic-custom", func(c *fiber.Ctx) error {
		panic("Ini adalah panic message kustom")
	})


	log.Println("Starting server on port 3000...")
	log.Println("Access /panic-now or /panic-custom to test recovery.")
	log.Fatal(app.Listen(":3000"))
}
