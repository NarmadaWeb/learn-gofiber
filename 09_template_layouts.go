package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2" // Import adapter
	"os" // Untuk cek direktori
)

func main() {
	viewDir := "./views"
	layoutDir := "./views/layouts"

	// Cek apakah direktori views dan layouts ada
	if _, err := os.Stat(viewDir); os.IsNotExist(err) {
		log.Fatalf("Direktori views '%s' tidak ditemukan.", viewDir)
	}
	if _, err := os.Stat(layoutDir); os.IsNotExist(err) {
		log.Fatalf("Direktori layout '%s' tidak ditemukan.", layoutDir)
	}
	// Cek file spesifik
	if _, err := os.Stat(filepath.Join(layoutDir, "main.html")); os.IsNotExist(err) { // Perlu import "path/filepath"
		log.Fatalf("File layout '%s' tidak ditemukan.", filepath.Join(layoutDir, "main.html"))
	}
	if _, err := os.Stat(filepath.Join(viewDir, "about.html")); os.IsNotExist(err) {
		log.Fatalf("File template '%s' tidak ditemukan.", filepath.Join(viewDir, "about.html"))
	}
	if _, err := os.Stat(filepath.Join(viewDir, "contact.html")); os.IsNotExist(err) {
		log.Fatalf("File template '%s' tidak ditemukan.", filepath.Join(viewDir, "contact.html"))
	}


	// Muat semua template .html dari direktori views
	// Penting: Berikan path ke direktori root views, bukan layouts
	engine := html.New(viewDir, ".html")
	engine.Reload(true) // Aktifkan reload saat development

	app := fiber.New(fiber.Config{
		Views: engine,
		// Anda BISA juga set default layout di sini:
		// ViewsLayout: "layouts/main",
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		log.Println("Rendering 'about' with layout 'layouts/main'")
		// Saat merender 'about', juga teruskan nama file layout 'main' relatif dari root views
		// Data akan tersedia di kedua template (layout dan content)
		// Jika ViewsLayout sudah diset di config, parameter ketiga tidak diperlukan.
		return c.Render("about", fiber.Map{"Email": "info-layout@example.com"}, "layouts/main")
	})

	// Rute lain bisa menggunakan layout yang sama
	app.Get("/contact", func(c *fiber.Ctx) error {
		log.Println("Rendering 'contact' with layout 'layouts/main'")
		// Buat views/contact.html serupa dengan about.html
		return c.Render("contact", fiber.Map{"Email": "kontak@example.com"}, "layouts/main")
	})

	// Rute tanpa layout eksplisit (akan menggunakan default jika diset, atau tanpa layout jika tidak)
	app.Get("/no-layout-test", func(c *fiber.Ctx) error {
		// Perlu file views/no_layout_test.html
		// {{define "title"}}Test Tanpa Layout Eksplisit{{end}}
		// {{define "content"}}Ini konten tanpa layout eksplisit.{{end}}
		log.Println("Rendering 'no_layout_test' (layout tergantung config ViewsLayout)")
		// return c.Render("no_layout_test", fiber.Map{"Email": "no-layout@example.com"})
		return c.Status(fiber.StatusNotImplemented).SendString("Buat file views/no_layout_test.html untuk tes ini")
	})


	log.Println("Starting server on port 3000...")
	log.Println("Pastikan direktori './views' dan './views/layouts' ada beserta file .html di dalamnya.")
	log.Println("Akses /about atau /contact")
	log.Fatal(app.Listen(":3000"))
}
