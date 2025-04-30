package main

import (
	"log"
	"os" // Untuk membuat direktori/file dummy
	"path/filepath" // Untuk path

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Fungsi untuk membuat direktori dan file dummy jika belum ada
func setupStaticDirs() {
	dirs := []string{"./public/css", "./public/js", "./public/images", "./assets", "./root_files"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Printf("Warning: Gagal membuat direktori %s: %v", dir, err)
		}
	}

	files := map[string]string{
		"./public/css/style.css": "body { background-color: #f0f0f0; }",
		"./public/js/script.js":  "console.log('Static script loaded!');",
		"./public/images/logo.png": "", // Buat file kosong, idealnya ganti dengan gambar asli
		"./assets/data.json":     `{"key": "value"}`,
		"./root_files/favicon.ico": "", // Buat file kosong, idealnya ganti dengan ikon asli
		"./root_files/index.html": "<html><body><h1>Root Index</h1><p>Served from /</p></body></html>",
	}

	for path, content := range files {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			errCreate := os.WriteFile(path, []byte(content), 0644)
			if errCreate != nil {
				log.Printf("Warning: Gagal membuat file dummy %s: %v", path, errCreate)
			} else {
				log.Printf("Membuat file dummy: %s", path)
			}
		}
	}
}


func main() {
	// Buat direktori dan file contoh jika belum ada
	setupStaticDirs()

	app := fiber.New()
	app.Use(logger.New())

	// 1. Sajikan dari direktori './public' di bawah URL '/static'
	log.Println("Serving static files from './public' under '/static' URL prefix")
	app.Static("/static", "./public", fiber.Static{
		Compress:  true, // Aktifkan kompresi
		ByteRange: true, // Izinkan byte range request
		Browse:    false, // Jangan izinkan browsing direktori
		MaxAge:    3600, // Cache 1 jam
	})
	// Contoh akses: http://localhost:3000/static/css/style.css
	// Contoh akses: http://localhost:3000/static/js/script.js
	// Contoh akses: http://localhost:3000/static/images/logo.png

	// 2. Sajikan dari direktori './assets' di bawah URL '/assets'
	log.Println("Serving static files from './assets' under '/assets' URL prefix")
	app.Static("/assets", "./assets") // Dengan config default
	// Contoh akses: http://localhost:3000/assets/data.json

	// 3. Sajikan dari root URL (prefix "/") dari direktori './root_files'
	//    Berguna untuk favicon.ico, index.html di root, dll.
	//    Middleware ini HARUS didaftarkan SETELAH rute API/aplikasi Anda
	//    jika ada potensi konflik path (misalnya jika Anda punya rute GET "/")
	//    Namun, untuk contoh ini kita letakkan di sini.
	log.Println("Serving static files from './root_files' under '/' URL prefix")
	app.Static("/", "./root_files", fiber.Static{
		Index: "index.html", // Sajikan index.html jika user akses "/"
	})
	// Contoh akses: http://localhost:3000/favicon.ico
	// Contoh akses: http://localhost:3000/ (akan menyajikan ./root_files/index.html)


	// Rute aplikasi Anda lainnya (contoh)
	app.Get("/api/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "API is running"})
	})

	// Halaman HTML sederhana yang mereferensikan file statis
	app.Get("/page", func(c *fiber.Ctx) error {
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>App Fiber Static Test</title>
			<link rel="stylesheet" href="/static/css/style.css">
			<link rel="icon" href="/favicon.ico">
		</head>
		<body>
			<h1>Selamat Datang!</h1>
			<p>Ini halaman yang memuat file statis.</p>
			<img src="/static/images/logo.png" alt="Logo Placeholder">
			<p><a href="/assets/data.json">Lihat data JSON</a></p>
			<script src="/static/js/script.js"></script>
		</body>
		</html>
		`
		// Kirim sebagai HTML
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendString(html)
	})


	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
