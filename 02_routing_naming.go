package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/users/:id/profile", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("Ini halaman profil untuk user ID: " + id)
	}).Name("user.profile")

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		profileURL, err := c.GetRouteURL("user.profile", fiber.Map{
			"id": "123",
		})
		if err != nil {
			log.Printf("Error getting route URL: %v", err)
			return err
		}

		return c.SendString("URL Profil Pengguna 123: " + profileURL)
	})

	app.Get("/generate-invalid", func(c *fiber.Ctx) error {
		_, err := c.GetRouteURL("route.tidak.ada", fiber.Map{})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Gagal generate URL: " + err.Error())
		}
		return c.SendString("Ini seharusnya tidak tercapai")
	})

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
