package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2" // Import adapter
	"os" // Untuk cek direktori
)

func main() {
	viewDir := "./views"

	// Cek apakah direktori views ada
	if _, err := os.Stat(viewDir); os.IsNotExist(err) {
		log.Fatalf("Direktori views '%s' tidak ditemukan. Pastikan Anda membuatnya dan menaruh index.html di dalamnya.", viewDir)
	}

	engine := html.New(viewDir, ".html") // Cari file .html di ./views
	engine.Reload(true) // Aktifkan reload saat development (nonaktifkan di produksi)
	engine.Debug(true) // Aktifkan debug logging saat development

	// (Opsional) Tambahkan fungsi kustom ke template
	engine.AddFunc("uppercase", func(s string) string {
		return strings.ToUpper(s) // Perlu import "strings"
	})


	// 2. Buat aplikasi Fiber dengan engine yang dikonfigurasi
	app := fiber.New(fiber.Config{
		Views: engine, // Beritahu Fiber untuk menggunakan engine ini
		ViewsLayout: "layouts/main", // Default layout (opsional)
	})

	// 3. Definisikan rute yang menggunakan c.Render()
	app.Get("/", func(c *fiber.Ctx) error {
		// Data yang akan dikirim ke template
		data := fiber.Map{
			"Title":  "Halaman Utama",
			"Header": "Halo dari Fiber!",
			"Items":  []string{"Apel", "Pisang", "Ceri"},
		}
		// Render template "index" (tanpa ekstensi .html)
		// dan kirim data ke dalamnya.
		// Jika ViewsLayout diset di config, tidak perlu sebutkan layout di sini
		// return c.Render("index", data, "layouts/main") // Contoh jika tidak set default layout
		log.Println("Rendering template 'index'")
		err := c.Render("index", data)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
		}
		return err
	})

	app.Get("/kosong", func(c *fiber.Ctx) error {
		// Contoh tanpa item
		data := fiber.Map{
			"Title":  "Halaman Kosong",
			"Header": "Tidak ada item",
			"Items":  nil, // Atau []string{}
		}
		log.Println("Rendering template 'index' (kosong)")
		// return c.Render("index", data, "layouts/main") // Contoh jika tidak set default layout
		return c.Render("index", data)
	})

	log.Println("Starting server on port 3000...")
	log.Println("Pastikan direktori './views' ada dan berisi 'index.html'")
	log.Fatal(app.Listen(":3000"))
}
