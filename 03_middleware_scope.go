package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func rateLimitMiddleware(c *fiber.Ctx) error {
	log.Println("Rate Limit Middleware Dijalankan!")
	return c.Next()
}

func specificAuthMiddleware(c *fiber.Ctx) error {
	log.Println("Specific Auth Middleware Dijalankan!")
	token := c.Query("token")
	if token != "special-token" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token khusus tidak valid"})
	}
	return c.Next()
}

func finalHandler(c *fiber.Ctx) error {
	log.Println("Final Handler Dijalankan!")
	return c.JSON(fiber.Map{"message": "Data berhasil disubmit!"})
}

func requireAdminLogin(c *fiber.Ctx) error {
	log.Println("Admin Login Check Middleware Dijalankan!")
	isAdmin := c.Get("X-User-Role") == "admin"
	if !isAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Hanya admin yang boleh akses"})
	}
	return c.Next()
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	admin := app.Group("/admin")
	admin.Use(requireAdminLogin)
	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendString("Selamat datang di Dashboard Admin!")
	})

	app.Get("/public", func(c *fiber.Ctx) error {
		return c.SendString("Ini halaman publik")
	})

	app.Post("/submit", rateLimitMiddleware, specificAuthMiddleware, finalHandler)

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
