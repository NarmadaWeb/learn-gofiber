package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func TimerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		stop := time.Now()
		duration := stop.Sub(start)
		c.Set("X-Request-Time", duration.String())
		log.Printf("Request ke %s memakan waktu %s", c.Path(), duration)
		return err
	}
}

func APIKeyAuthMiddleware(apiKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Get("X-API-Key")

		if key == "" {
			log.Println("Middleware Auth: Header X-API-Key tidak ditemukan")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Header X-API-Key dibutuhkan",
			})
		}

		if key != apiKey {
			log.Printf("Middleware Auth: API Key tidak valid: %s", key)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "API Key tidak valid",
			})
		}

		log.Println("Middleware Auth: API Key valid")
		return c.Next()
	}
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(TimerMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(50 * time.Millisecond)
		return c.SendString("Halaman Publik")
	})

	api := app.Group("/api")
	api.Use(APIKeyAuthMiddleware("kunci-rahasia-123"))

	api.Get("/data", func(c *fiber.Ctx) error {
		time.Sleep(100 * time.Millisecond)
		return c.JSON(fiber.Map{"message": "Ini data rahasia Anda!"})
	})

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
