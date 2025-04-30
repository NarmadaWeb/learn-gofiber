package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2" // Import adapter handlebars
	"os" // Untuk cek direktori
)

func main() {
	viewDir := "./views"

	// Cek apakah direktori views ada
	if _, err := os.Stat(viewDir); os.IsNotExist(err) {
		log.Fatalf("Direktori views '%s' tidak ditemukan. Pastikan Anda membuatnya dan menaruh profile.hbs di dalamnya.", viewDir)
	}

	// Buat engine Handlebars
	engine := handlebars.New(viewDir, ".hbs") // Ekstensi .hbs
	engine.Reload(true) // Aktifkan reload (nonaktifkan di produksi)

	// (Opsional) Tambahkan helpers kustom
	engine.AddFunc("formatDate", func(t time.Time) string { // Perlu import "time"
		return t.Format("02 January 2006")
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/profile/:id", func(c *fiber.Ctx) error {
		// Data contoh
		userData := fiber.Map{
			"ID": c.Params("id"),
			"Name": "John Doe",
			"Email": "john.doe@example.com",
		}
		isAdmin := c.QueryBool("admin", false) // Cek query ?admin=true

		// Kirim data ke template profile.hbs
		log.Printf("Rendering template 'profile' for ID %s, isAdmin: %t", c.Params("id"), isAdmin)
		err := c.Render("profile", fiber.Map{
			"user": userData,
			"isAdmin": isAdmin,
		})
		if err != nil {
			log.Printf("Error rendering template: %v", err)
		}
		return err
	})

	log.Println("Starting server on port 3000...")
	log.Println("Pastikan direktori './views' ada dan berisi 'profile.hbs'")
	log.Println("Akses /profile/123 atau /profile/456?admin=true")
	log.Fatal(app.Listen(":3000"))
}
