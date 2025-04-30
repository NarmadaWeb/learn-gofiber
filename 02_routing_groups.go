package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func adminAuthMiddleware(c *fiber.Ctx) error {
	log.Println("Middleware Admin Auth Dijalankan!")
	isAdmin := c.Get("X-Is-Admin") == "true"
	if !isAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Akses khusus admin"})
	}
	return c.Next()
}

func main() {
	app := fiber.New()

	apiV1 := app.Group("/api/v1")

	apiV1.Use(func(c *fiber.Ctx) error {
		log.Println("Middleware API V1 dijalankan!")
		apiKey := c.Get("X-API-Key")
		if apiKey != "secret-key" {
			log.Println("API Key tidak valid atau tidak ada, tetapi tetap melanjutkan (contoh)")
		}
		return c.Next()
	})

	apiV1.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "API v1 OK"})
	})

	apiV1.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON([]fiber.Map{{"id": 1, "name": "User Satu"}, {"id": 2, "name": "User Dua"}})
	})

	admin := apiV1.Group("/admin")
	admin.Use(adminAuthMiddleware)

	admin.Post("/settings", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Pengaturan admin disimpan"})
	})

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
