package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func myCustomErrorHandler(c *fiber.Ctx, err error) error {
	log.Printf("Custom error handler caught: %v", err)
	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
}

func main() {
	appDefault := fiber.New()
	log.Printf("Default App Name: %s", appDefault.Config().AppName)

	appWithConfig := fiber.New(fiber.Config{
		AppName:      "Aplikasi Keren Saya v1.0",
		Prefork:      false,
		ErrorHandler: myCustomErrorHandler,
	})
	log.Printf("Custom App Name: %s", appWithConfig.Config().AppName)

	appWithConfig.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Custom Config App!")
	})
	appWithConfig.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusConflict, "Ini adalah error tes")
	})

	log.Println("Starting server with custom config on port 3001...")
	log.Fatal(appWithConfig.Listen(":3001"))
}
