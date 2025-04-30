# Panduan Lengkap Belajar Go Fiber v2 🚀

![Go Fiber Logo](https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/gofiber/fiber/v2)](https://goreportcard.com/report/github.com/gofiber/fiber/v2)
[![GoDoc](https://godoc.org/github.com/gofiber/fiber/v2?status.svg)](https://pkg.go.dev/github.com/gofiber/fiber/v2)
[![Versi Rilis](https://img.shields.io/github/v/release/gofiber/fiber)](https://github.com/gofiber/fiber/releases)
[![Lisensi](https://img.shields.io/github/license/gofiber/fiber)](https://github.com/gofiber/fiber/blob/master/LICENSE)
[![Tes Kode](https://img.shields.io/github/actions/workflow/status/gofiber/fiber/test.yml?branch=master)](https://github.com/gofiber/fiber/actions/workflows/test.yml)
[![Cakupan Kode](https://coveralls.io/repos/github/gofiber/fiber/badge.svg?branch=master)](https://coveralls.io/github/gofiber/fiber?branch=master)

## Language [English🇬🇧](english.md) & [Indonesian🇮🇩](README.md)
---
Selamat datang di panduan lengkap untuk mempelajari **Go Fiber v2**, sebuah framework web Go yang terinspirasi oleh Express.js, dibangun di atas [Fasthttp](https://github.com/valyala/fasthttp), HTTP engine tercepat untuk Go. Fiber dirancang untuk **kemudahan pengembangan** dengan **kinerja tinggi** dan **penggunaan memori yang rendah**.

Panduan ini bertujuan untuk menjadi sumber daya komprehensif bagi pemula hingga pengembang tingkat menengah yang ingin membangun aplikasi web, API RESTful, atau microservices menggunakan Go Fiber v2.

---

## Daftar Isi 📖

1.  [Pendahuluan](#1-pendahuluan-)
	*   [Apa itu Go Fiber?](#apa-itu-go-fiber)
	*   [Mengapa Memilih Fiber?](#mengapa-memilih-fiber)
	*   [Fitur Utama](#fitur-utama)
	*   [Untuk Siapa Panduan Ini?](#untuk-siapa-panduan-ini)
2.  [Prasyarat](#2-prasyarat-)
3.  [Instalasi](#3-instalasi-)
4.  [Memulai: "Hello, World!" Klasik](#4-memulai-hello-world-klasik-)
	*   [Membuat Proyek Baru](#membuat-proyek-baru)
	*   [Kode Dasar](#kode-dasar)
	*   [Menjalankan Aplikasi](#menjalankan-aplikasi)
	*   [Menguji Aplikasi](#menguji-aplikasi)
5.  [Konsep Inti Fiber](#5-konsep-inti-fiber-)
	*   [Aplikasi Fiber (`fiber.App`)](#aplikasi-fiber-fiberapp)
	*   [Routing](#routing)
		*   [Metode HTTP Dasar](#metode-http-dasar)
		*   [Parameter Rute](#parameter-rute)
		*   [Parameter Opsional & Wildcard](#parameter-opsional--wildcard)
		*   [Grup Rute](#grup-rute)
		*   [Penamaan Rute](#penamaan-rute)
		*   [Melihat Daftar Rute](#melihat-daftar-rute)
	*   [Middleware](#middleware)
		*   [Apa itu Middleware?](#apa-itu-middleware)
		*   [Menggunakan Middleware Bawaan](#menggunakan-middleware-bawaan)
		*   [Membuat Middleware Kustom](#membuat-middleware-kustom)
		*   [Urutan Eksekusi Middleware](#urutan-eksekusi-middleware)
		*   [Melewatkan Middleware (`Next`)](#melewatkan-middleware-next)
		*   [Lingkup Middleware (Global, Grup, Rute)](#lingkup-middleware-global-grup-rute)
		*   [Middleware Pihak Ketiga](#middleware-pihak-ketiga)
	*   [Konteks (`fiber.Ctx`)](#konteks-fiberctx)
		*   [Akses Informasi Request](#akses-informasi-request)
		*   [Mengirim Response](#mengirim-response)
		*   [Meneruskan Data (Locals)](#meneruskan-data-locals)
		*   [Binding Data Request](#binding-data-request)
	*   [Penanganan Request](#penanganan-request-)
		*   [Membaca Headers](#membaca-headers)
		*   [Membaca Query Parameters](#membaca-query-parameters)
		*   [Membaca Route Parameters](#membaca-route-parameters)
		*   [Membaca Body Request](#membaca-body-request)
		*   [File Upload](#file-upload)
	*   [Penanganan Response](#penanganan-response-)
		*   [Mengatur Status Code](#mengatur-status-code)
		*   [Mengatur Headers](#mengatur-headers)
		*   [Mengirim Berbagai Tipe Data](#mengirim-berbagai-tipe-data)
	*   [Penanganan Error](#penanganan-error-)
		*   [Mengembalikan Error dari Handler](#mengembalikan-error-dari-handler)
		*   [Error Kustom (`fiber.NewError`)](#error-kustom-fibernewerror)
		*   [Custom Error Handler](#custom-error-handler)
		*   [Middleware Recover](#middleware-recover)
	*   [Konfigurasi (`fiber.Config`)](#konfigurasi-fiberconfig)
		*   [Opsi Konfigurasi Umum](#opsi-konfigurasi-umum)
		*   [Konfigurasi Prefork](#konfigurasi-prefork)
	*   [Template Engine](#template-engine-)
		*   [Konsep Template Engine](#konsep-template-engine)
		*   [Menggunakan Template Engine Bawaan (HTML)](#menggunakan-template-engine-bawaan-html)
		*   [Menggunakan Template Engine Lain](#menggunakan-template-engine-lain)
		*   [Layouts](#layouts)
	*   [Menyajikan File Statis](#menyajikan-file-statis-)
	*   [Validasi Request](#validasi-request-)
		*   [Pentingnya Validasi](#pentingnya-validasi)
		*   [Menggunakan Library Validator](#menggunakan-library-validator)
		*   [Contoh Implementasi](#contoh-implementasi)
6.  [Topik Lanjutan](#6-topik-lanjutan-)
	*   [WebSocket](#websocket)
	*   [Server-Sent Events (SSE)](#server-sent-events-sse)
	*   [Integrasi Database](#integrasi-database)
	*   [Autentikasi & Otorisasi (JWT, Sesi)](#autentikasi--otorisasi-jwt-sesi)
	*   [Pengujian (Testing)](#pengujian-testing)
	*   [Deployment](#deployment)
	*   [Kinerja & Optimasi](#kinerja--optimasi)
	*   [Struktur Proyek](#struktur-proyek)
	*   [Graceful Shutdown](#graceful-shutdown)
7.  [Contoh Aplikasi (CRUD Sederhana)](#7-contoh-aplikasi-crud-sederhana-)
8.  [Dokumentasi API](#8-dokumentasi-api-)
9.  [Praktik Terbaik (Best Practices)](#9-praktik-terbaik-best-practices-)
10. [Berkontribusi](#10-berkontribusi-)
11. [Lisensi](#11-lisensi-)
12. [Ucapan Terima Kasih](#12-ucapan-terima-kasih-)

---

## 1. Pendahuluan 🌟

### Apa itu Go Fiber?

Go Fiber adalah framework web untuk bahasa pemrograman Go yang terinspirasi kuat oleh [Express.js](https://expressjs.com/), framework Node.js yang sangat populer. Fiber dibangun di atas [Fasthttp](https://github.com/valyala/fasthttp), sebuah library HTTP performa tinggi untuk Go. Tujuannya adalah menyediakan antarmuka yang familiar dan mudah digunakan bagi pengembang yang mungkin sudah terbiasa dengan Express, sambil memanfaatkan kecepatan dan efisiensi Go serta Fasthttp.

Fiber fokus pada:

*   **Kinerja Tinggi:** Memanfaatkan Fasthttp untuk mencapai throughput yang sangat tinggi dan latensi rendah.
*   **Alokasi Memori Rendah:** Dirancang untuk meminimalkan alokasi memori selama pemrosesan request.
*   **Kemudahan Penggunaan:** API yang ekspresif dan mudah dipelajari, terutama jika Anda memiliki latar belakang Express.js.
*   **Fleksibilitas:** Ekosistem middleware yang kaya dan kemampuan untuk diperluas.

### Mengapa Memilih Fiber?

Ada beberapa alasan kuat untuk memilih Fiber untuk proyek web Go Anda berikutnya:

1.  **Kecepatan Luar Biasa:** Berkat Fasthttp, Fiber adalah salah satu framework Go tercepat yang tersedia. Ini sangat penting untuk aplikasi dengan lalu lintas tinggi atau yang membutuhkan latensi minimal.
2.  **Pengembangan Cepat:** API yang intuitif dan mirip Express memungkinkan Anda membangun aplikasi dengan cepat. Dokumentasi yang baik dan komunitas yang aktif juga membantu.
3.  **Penggunaan Memori Efisien:** Penting untuk aplikasi yang berjalan di lingkungan dengan sumber daya terbatas atau untuk mengurangi biaya hosting.
4.  **Ekosistem Middleware:** Fiber menyediakan banyak middleware bawaan (seperti logger, recovery, CORS) dan memudahkan integrasi atau pembuatan middleware kustom.
5.  **Routing Ekspresif:** Sistem routing yang kuat mendukung parameter, wildcard, grup, dan penamaan rute.
6.  **Dukungan Template Engine:** Mudah diintegrasikan dengan berbagai template engine Go.
7.  **Dukungan WebSocket & SSE:** Memiliki dukungan bawaan atau melalui middleware resmi untuk komunikasi real-time.
8.  **Fokus pada API:** Sangat cocok untuk membangun API RESTful modern.

### Fitur Utama

*   Routing yang kuat
*   Penyajian file statis
*   Middleware & dukungan Next()
*   API yang terinspirasi Express
*   Dukungan Template Engine (Go, Pug, Amber, dll.)
*   Ekosistem middleware yang berkembang pesat
*   Dukungan WebSocket
*   Server-Sent Events (SSE)
*   Rate Limiter
*   Dibangun di atas Fasthttp
*   Konfigurasi mudah
*   Dan banyak lagi...

### Untuk Siapa Panduan Ini?

*   **Pengembang Go Pemula:** Yang ingin belajar membangun aplikasi web dengan Go menggunakan framework modern.
*   **Pengembang Berpengalaman (dari bahasa lain):** Seperti Node.js (Express), Python (Flask/Django), Ruby (Rails) yang ingin beralih ke Go untuk pengembangan web.
*   **Pengembang Go Tingkat Menengah:** Yang ingin mendalami fitur-fitur Fiber dan praktik terbaiknya.

Diharapkan Anda memiliki pemahaman dasar tentang bahasa Go (sintaks, tipe data, fungsi, struct, interface) dan konsep dasar HTTP (request, response, metode, status code).

---

## 2. Prasyarat 🛠️

Sebelum Anda mulai belajar dan menggunakan Go Fiber, pastikan Anda memiliki:

1.  **Go Terinstal:** Versi Go 1.17 atau yang lebih baru direkomendasikan. Anda bisa mengunduh dan menginstal Go dari [situs resminya](https://go.dev/dl/).
	*   Verifikasi instalasi dengan membuka terminal dan menjalankan: `go version`
2.  **Pemahaman Dasar Go:**
	*   Sintaks dasar Go (deklarasi variabel, tipe data, kontrol alur seperti `if`, `for`).
	*   Fungsi dan metode.
	*   Struct dan Interface.
	*   Goroutine dan Channel (pemahaman dasar akan membantu, tetapi tidak wajib untuk memulai).
	*   Manajemen package dan module (`go mod`).
3.  **Pemahaman Dasar Web & HTTP:**
	*   Model Client-Server.
	*   Request & Response HTTP.
	*   Metode HTTP (GET, POST, PUT, DELETE, dll.).
	*   Status Code HTTP (200 OK, 404 Not Found, 500 Internal Server Error, dll.).
	*   Headers HTTP.
	*   JSON (JavaScript Object Notation) sebagai format pertukaran data umum.
4.  **Terminal/Command Prompt:** Anda akan sering menggunakan terminal untuk menjalankan perintah Go (build, run, test) dan alat bantu lainnya seperti `curl`.
5.  **Teks Editor atau IDE:** Pilihlah editor kode favorit Anda (VS Code, GoLand, Vim, Emacs, dll.) yang memiliki dukungan baik untuk Go.

---

## 3. Instalasi ⚙️

Menginstal Fiber sangat mudah menggunakan sistem modul Go.

1.  **Buat Direktori Proyek:**
	```bash
	mkdir proyek-fiber-saya
	cd proyek-fiber-saya
	```

2.  **Inisialisasi Go Modules:**
	Jika Anda memulai proyek baru, inisialisasi Go module. Ganti `nama-modul-anda` dengan path modul yang sesuai (misalnya, `github.com/username/proyek-fiber-saya`).
	```bash
	go mod init nama-modul-anda
	```
	Ini akan membuat file `go.mod` di direktori proyek Anda.

3.  **Tambahkan Fiber sebagai Dependensi:**
	Jalankan perintah berikut untuk mengunduh dan menambahkan Fiber v2 ke proyek Anda:
	```bash
	go get -u github.com/gofiber/fiber/v2
	```
	Perintah ini akan:
	*   Mengunduh package Fiber dan dependensinya.
	*   Menambahkan `github.com/gofiber/fiber/v2` sebagai `require` di file `go.mod` Anda.
	*   Membuat atau memperbarui file `go.sum` yang berisi checksum dependensi.

Sekarang Anda siap untuk mulai menggunakan Fiber dalam kode Go Anda!

---

## 4. Memulai: "Hello, World!" Klasik 👋

Mari kita buat aplikasi Fiber paling sederhana untuk memastikan semuanya bekerja dengan baik.

### Membuat Proyek Baru

Jika Anda belum melakukannya, ikuti langkah-langkah di bagian [Instalasi](#3-instalasi-).

### Kode Dasar

Buat file baru bernama `main.go` di direktori proyek Anda dan tambahkan kode berikut:

```go
// main.go
package main

import (
	"log" // Package untuk logging

	"github.com/gofiber/fiber/v2" // Import package Fiber v2
)

func main() {
	// 1. Membuat instance aplikasi Fiber baru
	//    Kita bisa memberikan konfigurasi kustom di sini,
	//    tapi untuk sekarang kita gunakan default.
	app := fiber.New()

	// 2. Mendefinisikan route untuk HTTP GET request ke path "/" (root)
	//    Ketika request GET datang ke "/", fungsi handler anonim akan dieksekusi.
	app.Get("/", func(c *fiber.Ctx) error {
		// 'c' adalah pointer ke Context Fiber, yang menyimpan informasi
		// request dan menyediakan metode untuk mengirim response.

		// Mengirim response string sederhana "Hello, World! 👋" ke client.
		// Metode SendString secara otomatis mengatur Content-Type ke text/plain.
		return c.SendString("Hello, World! 👋")
	})

	// 3. Menjalankan server HTTP pada port 3000
	//    Server akan mendengarkan koneksi masuk pada alamat "0.0.0.0:3000".
	//    app.Listen() adalah blocking call, artinya program akan berhenti di sini
	//    dan terus berjalan sampai server dihentikan (misalnya dengan Ctrl+C).
	log.Fatal(app.Listen(":3000"))
	// Kita menggunakan log.Fatal untuk menangkap error yang mungkin terjadi saat
	// memulai server (misalnya port sudah digunakan) dan menghentikan program.
}
```

**Penjelasan Kode:**

1.  **`import`**: Kita mengimpor package `log` standar Go untuk mencatat pesan (terutama error saat start server) dan package `fiber` itu sendiri.
2.  **`fiber.New()`**: Membuat instance baru dari aplikasi Fiber. Ini adalah titik awal dari setiap aplikasi Fiber. Anda bisa meneruskan struct `fiber.Config` ke fungsi ini untuk menyesuaikan perilaku aplikasi (akan dibahas nanti).
3.  **`app.Get("/", ...)`**: Mendefinisikan *route*. Ini memberi tahu Fiber bahwa ketika ada request HTTP dengan metode `GET` ke path `/`, fungsi *handler* yang diberikan (fungsi anonim dalam contoh ini) harus dieksekusi.
4.  **`func(c *fiber.Ctx) error`**: Ini adalah *handler* rute. Semua handler di Fiber menerima pointer ke `fiber.Ctx` (Konteks) dan mengembalikan `error`. Konteks (`c`) menyediakan akses ke data request (seperti header, parameter, body) dan metode untuk membangun serta mengirim response. Jika handler mengembalikan `nil`, berarti request berhasil diproses. Jika mengembalikan error, Fiber akan menanganinya (biasanya mengirim response error ke client).
5.  **`c.SendString(...)`**: Metode pada Konteks untuk mengirim response berupa string dengan status code `200 OK` dan `Content-Type: text/plain`.
6.  **`app.Listen(":3000")`**: Memulai server HTTP dan membuatnya mendengarkan koneksi pada port 3000 di semua network interface (`:3000` adalah singkatan dari `0.0.0.0:3000`). Fungsi ini akan memblokir eksekusi, menjaga server tetap berjalan untuk menerima request.
7.  **`log.Fatal(...)`**: Jika `app.Listen` mengembalikan error saat startup (misalnya port 3000 sudah digunakan), `log.Fatal` akan mencetak pesan error ke konsol dan menghentikan program.

### Menjalankan Aplikasi

Buka terminal Anda, arahkan ke direktori proyek (`proyek-fiber-saya`), dan jalankan perintah:

```bash
go run main.go
```

Anda akan melihat output yang mirip dengan ini (logo Fiber dan informasi port):

```
┌───────────────────────────────────────────────────┐
│                   Fiber v2.xx.x                   │
│               http://127.0.0.1:3000               │
│       (bound on host 0.0.0.0 and port 3000)       │
│                                                   │
│ Handlers ........... 1  Processes ........... 1   │
│ Prefork ....... Disabled  PID ............. xxxxx │
└───────────────────────────────────────────────────┘
```

Ini menunjukkan server Fiber Anda sedang berjalan dan siap menerima request di port 3000.

### Menguji Aplikasi

Ada dua cara mudah untuk menguji:

1.  **Menggunakan Web Browser:** Buka browser Anda dan navigasikan ke `http://localhost:3000` atau `http://127.0.0.1:3000`. Anda akan melihat teks "Hello, World! 👋" ditampilkan di halaman.

2.  **Menggunakan `curl` (dari terminal lain):** Buka jendela terminal baru (biarkan server tetap berjalan di terminal pertama) dan jalankan:
	```bash
	curl http://localhost:3000
	```
	Anda akan mendapatkan output:
	```
	Hello, World! 👋
	```

Selamat! Anda telah berhasil membuat dan menjalankan aplikasi Go Fiber pertama Anda.

---

## 5. Konsep Inti Fiber 🧠

Bagian ini akan membahas konsep-konsep fundamental yang perlu Anda pahami untuk bekerja secara efektif dengan Fiber.

### Aplikasi Fiber (`fiber.App`)

Objek `*fiber.App` yang Anda buat dengan `fiber.New()` adalah inti dari aplikasi Anda. Objek ini digunakan untuk:

*   Mendaftarkan rute (routes).
*   Menerapkan middleware.
*   Mengonfigurasi pengaturan aplikasi.
*   Memulai server HTTP.

```go
// Membuat instance dengan konfigurasi default
app := fiber.New()

// Membuat instance dengan konfigurasi kustom
appWithConfig := fiber.New(fiber.Config{
	AppName:      "Aplikasi Keren Saya v1.0",
	Prefork:      true, // Mengaktifkan mode prefork (akan dibahas nanti)
	ErrorHandler: myCustomErrorHandler, // Mengatur error handler kustom
})
```

### Routing

Routing adalah proses menentukan bagaimana aplikasi merespons request client ke endpoint tertentu (URI atau path) dan metode HTTP spesifik (GET, POST, dll.). Fiber menyediakan sistem routing yang sangat fleksibel dan cepat.

#### Metode HTTP Dasar

Anda dapat mendaftarkan rute untuk semua metode HTTP standar menggunakan metode yang sesuai pada instance `fiber.App`:

```go
app := fiber.New()

// GET: Mengambil data
app.Get("/users", func(c *fiber.Ctx) error {
	return c.SendString("Mengambil daftar pengguna (GET)")
})

// POST: Membuat data baru
app.Post("/users", func(c *fiber.Ctx) error {
	return c.SendString("Membuat pengguna baru (POST)")
})

// PUT: Memperbarui data secara keseluruhan
app.Put("/users/:id", func(c *fiber.Ctx) error {
	id := c.Params("id") // Mengambil parameter 'id' dari URL
	return c.SendString("Memperbarui pengguna dengan ID: " + id + " (PUT)")
})

// PATCH: Memperbarui sebagian data
app.Patch("/users/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Memperbarui sebagian data pengguna ID: " + id + " (PATCH)")
})

// DELETE: Menghapus data
app.Delete("/users/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Menghapus pengguna dengan ID: " + id + " (DELETE)")
})

// OPTIONS: Mendapatkan opsi komunikasi yang diizinkan
app.Options("/info", func(c *fiber.Ctx) error {
	c.Set("Allow", "GET, POST, OPTIONS")
	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
})

// HEAD: Sama seperti GET tapi tanpa body response (hanya headers)
app.Head("/status", func(c *fiber.Ctx) error {
	c.Set("X-App-Status", "OK")
	return c.SendStatus(fiber.StatusOK) // 200 OK (tanpa body)
})

// app.All() untuk mencocokkan semua metode HTTP
app.All("/universal", func(c *fiber.Ctx) error {
	return c.SendString("Endpoint ini merespons semua metode HTTP.")
})
```

#### Parameter Rute

Seringkali Anda perlu menangkap segmen dinamis dari URL, seperti ID pengguna atau slug artikel. Fiber memungkinkan ini dengan *parameter rute*, yang diawali dengan titik dua (`:`).

```go
// Rute: /users/:userId/books/:bookId
app.Get("/users/:userId/books/:bookId", func(c *fiber.Ctx) error {
	// Mengambil nilai parameter menggunakan c.Params("namaParameter")
	userId := c.Params("userId")
	bookId := c.Params("bookId")

	return c.SendString("User ID: " + userId + ", Book ID: " + bookId)
	// Contoh request: GET /users/123/books/abc
	// Output: User ID: 123, Book ID: abc
})
```

Fiber juga menyediakan cara yang sedikit lebih efisien untuk mengambil parameter jika Anda tahu tipe datanya:

```go
app.Get("/product/:id", func(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // Mencoba parse parameter "id" sebagai integer
	if err != nil {
		// Jika parameter bukan integer yang valid, kirim error 400
		return c.Status(fiber.StatusBadRequest).SendString("ID produk harus berupa angka")
	}
	// ... proses dengan id (tipe int) ...
	return c.JSON(fiber.Map{"product_id": id, "status": "found"})
})
```
Metode `ParamsInt`, `ParamsBool`, `ParamsFloat` tersedia.

#### Parameter Opsional & Wildcard

*   **Parameter Opsional:** Tandai parameter dengan tanda tanya (`?`) untuk membuatnya opsional. Handler Anda perlu memeriksa apakah parameter tersebut ada.
	```go
	// Rute: /order/:id?
	app.Get("/order/:id?", func(c *fiber.Ctx) error {
		id := c.Params("id") // Akan kosong jika tidak ada di URL
		if id == "" {
			return c.SendString("Menampilkan semua order")
			// Request: GET /order
		}
		return c.SendString("Menampilkan detail order ID: " + id)
		// Request: GET /order/55
	})
	```

*   **Wildcard (`*`)**: Mencocokkan apa saja (termasuk `/`). Berguna untuk menangkap path yang panjang atau tidak tentu. Nilainya diambil dengan `c.Params("*")`.
	```go
	// Rute: /files/*
	app.Get("/files/*", func(c *fiber.Ctx) error {
		filePath := c.Params("*") // Mengambil semua setelah /files/
		return c.SendString("Mengakses file di path: " + filePath)
		// Request: GET /files/images/logo.png -> Output: Mengakses file di path: images/logo.png
		// Request: GET /files/docs/report.pdf -> Output: Mengakses file di path: docs/report.pdf
	})
	```
	*   **Penting:** Wildcard `*` harus berada di akhir path rute.

*   **Wildcard Parameter (`+`)**: Mirip dengan `*`, tetapi *harus* mencocokkan setidaknya satu karakter.
	```go
	// Rute: /user/+
	app.Get("/user/+", func(c *fiber.Ctx) error {
		name := c.Params("+") // Mengambil semua setelah /user/
		return c.SendString("Profil pengguna: " + name)
		// Request: GET /user/johndoe -> Output: Profil pengguna: johndoe
		// Request: GET /user/jane/doe (tidak cocok jika StrictRouting aktif, cocok jika tidak)
		// Request: GET /user/ (tidak cocok karena + butuh minimal 1 karakter)
	})
	```
	*   **Perbedaan `*` dan `+`**: `*` bisa cocok dengan string kosong (jika di akhir URL), `+` tidak bisa. `*` mencocokkan segalanya termasuk `/`, `+` biasanya digunakan untuk satu segmen path non-kosong.

#### Grup Rute

Grup rute sangat berguna untuk mengorganisir rute yang memiliki prefix path atau middleware yang sama.

```go
app := fiber.New()

// Membuat grup untuk semua rute di bawah /api/v1
apiV1 := app.Group("/api/v1")

// Menambahkan middleware spesifik untuk grup ini (misalnya, autentikasi)
apiV1.Use(func(c *fiber.Ctx) error {
	log.Println("Middleware API V1 dijalankan!")
	// Cek header autentikasi di sini...
	return c.Next() // Lanjutkan ke handler/middleware berikutnya
})

// Rute di dalam grup (path relatif terhadap prefix grup)
// Handler untuk GET /api/v1/status
apiV1.Get("/status", func(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "API v1 OK"})
})

// Handler untuk GET /api/v1/users
apiV1.Get("/users", func(c *fiber.Ctx) error {
	// ... logika mengambil pengguna ...
	return c.JSON([]fiber.Map{{"id": 1, "name": "User Satu"}, {"id": 2, "name": "User Dua"}})
})

// Anda bisa membuat grup di dalam grup (nested group)
admin := apiV1.Group("/admin")
admin.Use(adminAuthMiddleware) // Middleware khusus admin

// Handler untuk POST /api/v1/admin/settings
admin.Post("/settings", func(c *fiber.Ctx) error {
	// ... logika menyimpan pengaturan admin ...
	return c.JSON(fiber.Map{"message": "Pengaturan admin disimpan"})
})
```
Menggunakan `app.Group()` membuat kode lebih terstruktur, mudah dibaca, dan mengurangi redundansi (misalnya, tidak perlu mengetik `/api/v1` berulang kali atau menerapkan middleware yang sama ke banyak rute secara manual).

#### Penamaan Rute

Anda dapat memberi nama pada rute Anda. Ini berguna jika Anda perlu membuat URL secara dinamis di tempat lain dalam aplikasi Anda (misalnya dalam template atau redirect).

```go
// Memberi nama "user.profile" pada rute
app.Get("/users/:id/profile", func(c *fiber.Ctx) error {
	// ...
	return c.SendString("Ini halaman profil")
}).Name("user.profile")

// Di handler lain, membuat URL untuk rute bernama
app.Get("/dashboard", func(c *fiber.Ctx) error {
	// Membuat URL untuk user dengan ID 123
	profileURL, err := c.GetRouteURL("user.profile", fiber.Map{
		"id": "123", // Menyediakan nilai untuk parameter :id
	})
	if err != nil {
		return err // Tangani error jika rute tidak ditemukan atau parameter salah
	}
	// profileURL akan berisi "/users/123/profile"

	// Redirect ke halaman profil pengguna
	// return c.Redirect(profileURL)

	return c.SendString("URL Profil Pengguna 123: " + profileURL)
})
```
Penamaan rute meningkatkan maintainabilitas kode karena Anda tidak perlu hardcode URL di banyak tempat. Jika path rute berubah, Anda hanya perlu mengubahnya di satu tempat (definisi rute), dan semua pemanggilan `GetRouteURL` akan otomatis menghasilkan URL yang benar.

#### Melihat Daftar Rute

Untuk keperluan debugging atau introspeksi, Anda bisa mendapatkan daftar semua rute yang telah terdaftar di aplikasi Fiber.

```go
app.Get("/debug/routes", func(c *fiber.Ctx) error {
	// Mendapatkan slice dari semua rute
	routes := app.GetRoutes(true) // true untuk menyertakan rute internal Fiber

	// Format output (misalnya sebagai JSON)
	var routeList []fiber.Map
	for _, route := range routes {
		routeList = append(routeList, fiber.Map{
			"method": route.Method,
			"name":   route.Name,
			"path":   route.Path,
			"params": route.Params,
		})
	}
	return c.JSON(routeList)
})

// Jalankan server dan akses /debug/routes untuk melihat hasilnya
```

### Middleware

Middleware adalah salah satu konsep paling kuat dalam framework web modern seperti Fiber. Middleware adalah fungsi yang memiliki akses ke objek **Konteks (`fiber.Ctx`)** dan fungsi **`Next()`** dalam siklus request-response aplikasi.

#### Apa itu Middleware?

Middleware adalah "penjaga gerbang" atau "pemroses perantara" yang dieksekusi **sebelum** atau **setelah** handler rute utama Anda. Tugasnya bisa bermacam-macam:

*   **Logging:** Mencatat detail setiap request yang masuk.
*   **Authentication/Authorization:** Memeriksa kredensial pengguna atau izin akses.
*   **Data Validation/Sanitization:** Memvalidasi data input atau membersihkannya.
*   **Compression:** Mengompresi body response (misalnya dengan Gzip).
*   **CORS (Cross-Origin Resource Sharing):** Mengatur header agar API bisa diakses dari domain berbeda.
*   **Rate Limiting:** Membatasi jumlah request dari satu IP dalam periode waktu tertentu.
*   **Header Manipulation:** Menambah, mengubah, atau menghapus header request/response.
*   **Error Handling:** Menangkap error yang terjadi di handler berikutnya.
*   **Caching:** Menyimpan response untuk request yang sama.

#### Menggunakan Middleware Bawaan

Fiber hadir dengan banyak middleware siap pakai yang umum digunakan. Anda dapat menemukannya di package `github.com/gofiber/fiber/v2/middleware/...`.

Cara menggunakannya adalah dengan memanggil `app.Use()` atau metode `Use()` pada grup rute.

```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // Import middleware logger
	"github.com/gofiber/fiber/v2/middleware/recover" // Import middleware recover
	"github.com/gofiber/fiber/v2/middleware/cors"    // Import middleware CORS
)

func main() {
	app := fiber.New()

	// 1. Menggunakan Middleware Secara Global (berlaku untuk semua rute)
	app.Use(recover.New()) // Middleware untuk menangkap panic dan mengirim 500 Internal Server Error
	app.Use(logger.New(logger.Config{ // Middleware untuk logging request
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New()) // Middleware untuk mengaktifkan CORS dengan konfigurasi default

	// Rute Anda
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello dengan Middleware!")
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Ups, terjadi panic!") // Middleware recover akan menangkap ini
	})

	log.Fatal(app.Listen(":3000"))
}
```

Ketika Anda menjalankan aplikasi ini dan membuat request (misalnya `curl http://localhost:3000`), Anda akan melihat log request di terminal Anda karena `logger.New()` aktif. Jika Anda mengakses `/panic`, Anda tidak akan melihat server crash, melainkan mendapat response `500 Internal Server Error` karena `recover.New()` menangkap panic tersebut.

Middleware bawaan lainnya yang populer:

*   `basicauth`: Autentikasi HTTP Basic.
*   `compress`: Kompresi response (Gzip, Deflate, Brotli).
*   `etag`: Generasi ETag header untuk caching.
*   `limiter`: Rate limiting.
*   `monitor`: Menampilkan metrik performa aplikasi.
*   `pprof`: Profiling aplikasi Go.
*   `requestid`: Menambahkan ID unik ke setiap request.
*   `session`: Manajemen sesi.
*   Dan masih banyak lagi (lihat dokumentasi Fiber).

#### Membuat Middleware Kustom

Anda dapat dengan mudah membuat middleware Anda sendiri. Middleware kustom hanyalah sebuah fungsi yang mengikuti signature `func(*fiber.Ctx) error`.

```go
package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Contoh Middleware Kustom: Menambahkan header X-Request-Time
func TimerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now() // Catat waktu mulai

		// Panggil handler/middleware berikutnya dalam rantai
		// Ini penting! Tanpa c.Next(), request tidak akan sampai ke handler rute.
		err := c.Next()

		// Kode di sini dieksekusi SETELAH handler rute selesai
		stop := time.Now()
		duration := stop.Sub(start)

		// Tambahkan header kustom ke response
		c.Set("X-Request-Time", duration.String())
		log.Printf("Request ke %s memakan waktu %s", c.Path(), duration)

		// Kembalikan error jika ada dari handler/middleware berikutnya
		return err
	}
}

// Contoh Middleware Kustom: Pemeriksaan Header API Key Sederhana
func APIKeyAuthMiddleware(apiKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil nilai header 'X-API-Key' dari request
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

		// Jika API Key valid, lanjutkan ke handler berikutnya
		log.Println("Middleware Auth: API Key valid")
		return c.Next()
	}
}

func main() {
	app := fiber.New()

	// Gunakan middleware logger bawaan
	app.Use(logger.New())

	// Gunakan middleware timer kustom (global)
	app.Use(TimerMiddleware())

	// Rute publik
	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(50 * time.Millisecond) // Simulasi kerja
		return c.SendString("Halaman Publik")
	})

	// Grup rute yang memerlukan API Key
	api := app.Group("/api")
	// Terapkan middleware API Key hanya untuk grup /api
	api.Use(APIKeyAuthMiddleware("kunci-rahasia-123"))

	api.Get("/data", func(c *fiber.Ctx) error {
		time.Sleep(100 * time.Millisecond) // Simulasi kerja
		return c.JSON(fiber.Map{"message": "Ini data rahasia Anda!"})
	})

	log.Fatal(app.Listen(":3000"))
}
```

**Penjelasan:**

*   **`TimerMiddleware`**: Mencatat waktu sebelum memanggil `c.Next()` dan setelahnya, lalu menghitung durasi dan menambahkannya sebagai header response.
*   **`APIKeyAuthMiddleware`**: Menerima `apiKey` yang valid saat dibuat. Di dalamnya, ia memeriksa header `X-API-Key`. Jika tidak ada atau tidak cocok, ia mengirim response `401 Unauthorized` dan *tidak* memanggil `c.Next()`, menghentikan request. Jika cocok, ia memanggil `c.Next()` agar request bisa diproses oleh handler rute `/api/data`.
*   Perhatikan bagaimana `TimerMiddleware` diterapkan secara global (`app.Use`), sedangkan `APIKeyAuthMiddleware` hanya diterapkan pada grup `/api` (`api.Use`).

#### Urutan Eksekusi Middleware

Middleware dieksekusi dalam urutan di mana mereka ditambahkan menggunakan `app.Use()` atau `group.Use()`.

```go
app.Use(Middleware1)
app.Use(Middleware2)

api := app.Group("/api")
api.Use(Middleware3)

api.Get("/test", Handler)

// Urutan eksekusi untuk GET /api/test:
// 1. Middleware1
// 2. Middleware2
// 3. Middleware3
// 4. Handler
```

#### Melewatkan Middleware (`Next`)

Fungsi `c.Next()` sangat penting dalam middleware. Fungsinya adalah untuk **meneruskan kontrol ke handler atau middleware berikutnya** dalam rantai eksekusi.

*   Jika middleware **memanggil `c.Next()`**, eksekusi akan berlanjut ke fungsi berikutnya. Kode setelah `c.Next()` dalam middleware akan dieksekusi *setelah* handler/middleware berikutnya selesai.
*   Jika middleware **tidak memanggil `c.Next()`**, siklus request-response akan berhenti di middleware tersebut. Middleware tersebut bertanggung jawab penuh untuk mengirim response ke client. Ini berguna untuk kasus seperti autentikasi gagal, validasi gagal, atau caching.

#### Lingkup Middleware (Global, Grup, Rute)

Anda dapat menerapkan middleware pada level yang berbeda:

1.  **Global:** Menggunakan `app.Use()`. Middleware ini akan berjalan untuk **setiap** request yang masuk ke aplikasi, sebelum pencocokan rute dilakukan (kecuali untuk middleware seperti `Static`).
	```go
	app.Use(logger.New()) // Berjalan untuk semua request
	```

2.  **Grup:** Menggunakan `group.Use()`. Middleware ini hanya akan berjalan untuk request yang cocok dengan prefix grup tersebut.
	```go
	admin := app.Group("/admin")
	admin.Use(requireAdminLogin) // Hanya berjalan untuk rute di bawah /admin
	admin.Get("/dashboard", ...)
	```

3.  **Rute Spesifik:** Anda dapat meneruskan satu atau lebih fungsi middleware *sebelum* handler rute utama saat mendefinisikan rute.
	```go
	func rateLimitMiddleware(c *fiber.Ctx) error { ... return c.Next() }
	func specificAuthMiddleware(c *fiber.Ctx) error { ... return c.Next() }
	func finalHandler(c *fiber.Ctx) error { ... }

	// Middleware rateLimit dan specificAuth akan berjalan sebelum finalHandler
	// hanya untuk request POST ke /submit
	app.Post("/submit", rateLimitMiddleware, specificAuthMiddleware, finalHandler)
	```

#### Middleware Pihak Ketiga

Selain middleware bawaan, ada banyak middleware yang dikembangkan oleh komunitas Fiber atau pengembang Go lainnya. Anda bisa menemukannya di GitHub atau sumber lain. Cara menggunakannya biasanya sama: impor package-nya dan gunakan dengan `app.Use()`.

Pastikan untuk membaca dokumentasi middleware pihak ketiga untuk memahami cara kerja dan konfigurasinya.

### Konteks (`fiber.Ctx`)

Objek `*fiber.Ctx` adalah jantung dari penanganan request di Fiber. Setiap fungsi handler dan middleware menerima pointer ke objek ini. `Ctx` menyediakan semua yang Anda butuhkan untuk:

*   Mengakses informasi tentang request yang masuk (metode, path, headers, query params, route params, body, IP client, dll.).
*   Mengirim response kembali ke client (mengatur status code, headers, mengirim body dalam berbagai format).
*   Meneruskan data antara middleware dan handler dalam satu siklus request.
*   Mengelola siklus hidup request (misalnya, memanggil `Next()`).

Mari kita lihat beberapa metode `Ctx` yang paling penting:

#### Akses Informasi Request

*   **`c.Method()`**: Mendapatkan metode HTTP (string, e.g., "GET", "POST").
*   **`c.Path()`**: Mendapatkan path request (string, e.g., "/users/123").
*   **`c.BaseURL()`**: Mendapatkan base URL (e.g., "http://example.com").
*   **`c.OriginalURL()`**: Mendapatkan URL asli termasuk query string.
*   **`c.Hostname()`**: Mendapatkan hostname dari header `Host`.
*   **`c.IP()`**: Mendapatkan alamat IP client (mempertimbangkan header proxy seperti `X-Forwarded-For`).
*   **`c.IPs()`**: Mendapatkan daftar IP jika ada proxy (dari `X-Forwarded-For`).
*   **`c.Protocol()`**: Mendapatkan protokol request (string, e.g., "http", "https").
*   **`c.Secure()`**: Mengecek apakah koneksi menggunakan HTTPS (boolean).
*   **`c.Get(key string, defaultValue ...string)`**: Mendapatkan nilai header request. `key` tidak case-sensitive.
	```go
	ua := c.Get("User-Agent")
	auth := c.Get("Authorization", "default_value_if_not_found")
	```
*   **`c.Params(key string, defaultValue ...string)`**: Mendapatkan nilai parameter rute.
	```go
	userID := c.Params("id")
	```
*   **`c.ParamsInt(key string)`**, **`c.ParamsFloat(key string)`**, **`c.ParamsBool(key string)`**: Mendapatkan parameter rute dan mengonversinya ke tipe yang sesuai. Mengembalikan error jika konversi gagal.
*   **`c.Query(key string, defaultValue ...string)`**: Mendapatkan nilai query parameter dari URL.
	```go
	// URL: /search?q=fiber&page=2
	searchTerm := c.Query("q") // "fiber"
	page := c.Query("page", "1") // "2" (default "1" jika tidak ada)
	```
*   **`c.QueryParser(out interface{}) error`**: Mem-parse query string ke dalam struct Go. Berguna untuk parameter pencarian/filter yang kompleks.
	```go
	type SearchQuery struct {
		Query string `query:"q"`
		Limit int    `query:"limit"`
		Page  int    `query:"page"`
	}
	var sq SearchQuery
	if err := c.QueryParser(&sq); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Query tidak valid")
	}
	// sq.Query, sq.Limit, sq.Page akan terisi
	```
*   **`c.Body()`**: Mendapatkan body request sebagai `[]byte`. Gunakan ini jika Anda perlu mengakses body mentah.
	```go
	rawBody := c.Body()
	// Hati-hati: Membaca body bisa menghabiskan memori jika body besar.
	// Pertimbangkan BodyParser atau batasan ukuran body.
	```
*   **`c.BodyParser(out interface{}) error`**: Mem-parse body request ke dalam struct Go. Secara otomatis mendeteksi `Content-Type` (JSON, XML, form) dan melakukan unmarshaling. Ini adalah cara yang paling umum dan direkomendasikan untuk menangani data input.
	```go
	type CreateUserInput struct {
		Name  string `json:"name" xml:"name" form:"name"`
		Email string `json:"email" xml:"email" form:"email"`
	}
	var input CreateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body tidak valid: " + err.Error()})
	}
	// input.Name dan input.Email akan terisi
	```
*   **`c.FormValue(key string, defaultValue ...string)`**: Mendapatkan nilai dari form (application/x-www-form-urlencoded atau multipart/form-data).
*   **`c.FormFile(key string)`**: Mendapatkan file yang diunggah dari form multipart/form-data. Mengembalikan `*multipart.FileHeader`.
*   **`c.SaveFile(fileheader *multipart.FileHeader, path string)`**: Menyimpan file yang diunggah ke path yang ditentukan.
*   **`c.Is(contentType string)`**: Memeriksa apakah `Content-Type` request cocok (misalnya, `c.Is("json")`).
*   **`c.Accepts(offers ...string)`**: Memeriksa header `Accept` client dan menentukan tipe konten terbaik yang didukung (misalnya, `c.Accepts("json", "html")`).

#### Mengirim Response

*   **`c.SendStatus(statusCode int)`**: Mengirim response hanya dengan status code (tanpa body).
	```go
	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
	```
*   **`c.Status(statusCode int)`**: Mengatur status code untuk response berikutnya. Berguna untuk dirangkai dengan metode pengirim body.
	```go
	return c.Status(fiber.StatusCreated).JSON(newUser) // 201 Created
	```
*   **`c.Set(key string, val string)`**: Mengatur header response.
	```go
	c.Set("X-Custom-Header", "Nilai Saya")
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON) // Cara lain set Content-Type
	```
*   **`c.Append(key string, values ...string)`**: Menambahkan nilai ke header yang sudah ada (misalnya `Link` atau `Set-Cookie`).
*   **`c.SendString(body string)`**: Mengirim response string dengan `Content-Type: text/plain`.
*   **`c.Send(body []byte)`**: Mengirim response body berupa slice byte. `Content-Type` akan dideteksi secara otomatis (jika memungkinkan) atau default ke `application/octet-stream`.
*   **`c.JSON(data interface{})`**: Mengirim response JSON. Mengonversi `data` (struct, map, slice) ke JSON dan mengatur `Content-Type: application/json`. Ini sangat umum digunakan untuk API.
	```go
	user := User{ID: 1, Name: "Fiber"}
	return c.JSON(user)
	// atau
	return c.JSON(fiber.Map{"status": "success", "data": user})
	```
*   **`c.XML(data interface{})`**: Mengirim response XML.
*   **`c.Render(name string, bind interface{}, layouts ...string)`**: Merender template HTML (memerlukan konfigurasi template engine).
*   **`c.SendFile(filepath string, compress ...bool)`**: Mengirim file sebagai response. `Content-Type` biasanya dideteksi dari ekstensi file. Opsi `compress` (default true) memungkinkan kompresi Gzip jika client mendukung.
	```go
	return c.SendFile("./public/images/logo.png")
	```
*   **`c.Download(filepath string, filename ...string)`**: Mirip `SendFile`, tetapi menambahkan header `Content-Disposition: attachment`, yang memberitahu browser untuk mengunduh file alih-alih menampilkannya. Anda bisa memberikan nama file unduhan kustom.
	```go
	return c.Download("./private/report.pdf", "Laporan Bulanan.pdf")
	```
*   **`c.Redirect(location string, status ...int)`**: Mengirim response redirect (default status 302 Found).
	```go
	return c.Redirect("/login", fiber.StatusTemporaryRedirect) // 307
	```
*   **`c.Cookie(cookie *fiber.Cookie)`**: Mengatur cookie response.
	```go
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    "random-session-string",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true, // Hanya kirim via HTTPS
		SameSite: "Lax",
	})
	```
*   **`c.ClearCookie(key ...string)`**: Menghapus cookie dari browser client.

#### Meneruskan Data (Locals)

Terkadang Anda perlu meneruskan data dari satu middleware ke middleware lain atau ke handler rute utama dalam *siklus request yang sama*. Misalnya, middleware autentikasi bisa memverifikasi pengguna dan kemudian meneruskan ID pengguna atau objek pengguna ke handler. `c.Locals()` adalah cara untuk melakukan ini.

`c.Locals()` bekerja seperti map `string` ke `interface{}` yang terikat pada konteks request tersebut.

```go
// Middleware: Mendapatkan data pengguna (misalnya dari token)
func UserAuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	// ... validasi token dan dapatkan info pengguna ...
	user := User{ID: 123, Role: "admin"} // Contoh data pengguna

	// Simpan data pengguna di Locals
	c.Locals("currentUser", user)
	c.Locals("requestID", "xyz-789") // Bisa simpan tipe data lain

	log.Println("UserAuthMiddleware: Pengguna ditemukan dan disimpan di Locals")
	return c.Next() // Lanjutkan ke handler berikutnya
}

// Handler Rute: Menggunakan data dari Locals
func GetUserProfile(c *fiber.Ctx) error {
	// Ambil data dari Locals
	reqID := c.Locals("requestID").(string) // Perlu type assertion
	user, ok := c.Locals("currentUser").(User) // Gunakan type assertion dengan check 'ok'

	log.Printf("GetUserProfile: Request ID = %s", reqID)

	if !ok {
		log.Println("GetUserProfile: Data pengguna tidak ditemukan di Locals!")
		// Ini seharusnya tidak terjadi jika middleware UserAuthMiddleware selalu berjalan sebelumnya
		return c.Status(fiber.StatusInternalServerError).SendString("Error internal: data pengguna hilang")
	}

	// Gunakan data pengguna
	log.Printf("GetUserProfile: Mengambil profil untuk pengguna ID %d (%s)", user.ID, user.Role)
	return c.JSON(fiber.Map{
		"message":   "Profil Pengguna",
		"user_id":   user.ID,
		"user_role": user.Role,
		"request_id": reqID,
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Terapkan middleware auth sebelum handler profil
	app.Get("/profile", UserAuthMiddleware, GetUserProfile)

	log.Fatal(app.Listen(":3000"))
}
```

**Penting:**
*   Data di `c.Locals()` hanya ada selama siklus hidup satu request. Request berikutnya akan memiliki `Locals` yang kosong.
*   Saat mengambil data dari `Locals()`, Anda perlu melakukan *type assertion* (e.g., `.(string)`, `.(User)`) karena `Locals()` menyimpan nilai sebagai `interface{}`. Selalu periksa boolean `ok` kedua dari type assertion untuk menangani kasus di mana kunci tidak ada atau tipe datanya tidak sesuai.

#### Binding Data Request

Seperti yang disebutkan sebelumnya, `c.BodyParser()`, `c.QueryParser()`, dan `c.ParamsParser()` (untuk route params) adalah cara mudah untuk secara otomatis mem-parse data request masuk dan mengisinya ke dalam struct Go. Ini sangat mengurangi kode boilerplate untuk ekstraksi dan konversi data manual.

```go
type ProductFilter struct {
	Category string `query:"category"` // Dari query string ?category=...
	MaxPrice int    `query:"max_price"`// Dari query string ?max_price=...
	SortBy   string `query:"sort"`     // Dari query string ?sort=...
}

type UpdateProductInput struct {
	Name        string  `json:"name" form:"name"` // Dari JSON body atau Form data
	Description *string `json:"description" form:"description"` // Pointer untuk nilai opsional
	Price       float64 `json:"price" form:"price" validate:"required,gt=0"` // validasi
	IsActive    bool    `json:"is_active" form:"is_active"`
}

type ProductRouteParams struct {
	ProductID int `params:"id"` // Dari route parameter /products/:id
}

func SearchProducts(c *fiber.Ctx) error {
	var filter ProductFilter
	// Bind query params ke struct filter
	if err := c.QueryParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Filter tidak valid"})
	}
	// Gunakan filter.Category, filter.MaxPrice, filter.SortBy
	// ... logika pencarian ...
	return c.JSON(fiber.Map{"message": "Hasil pencarian", "filters": filter})
}

func UpdateProduct(c *fiber.Ctx) error {
	var params ProductRouteParams
	// Bind route params ke struct params
	if err := c.ParamsParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID Produk tidak valid"})
	}

	var input UpdateProductInput
	// Bind JSON/Form body ke struct input
	if err := c.BodyParser(&input); err != nil {
		// Cek jika error karena body kosong (jika diperbolehkan)
		if err == fiber.ErrUnprocessableEntity {
			 return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body request kosong atau format salah"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Data input tidak valid: " + err.Error()})
	}

	// Di sini Anda bisa menambahkan validasi tambahan (lihat bagian Validasi)
	// validate := validator.New()
	// if err := validate.Struct(input); err != nil { ... }

	// Gunakan params.ProductID dan input.Name, input.Description, dll.
	// ... logika pembaruan produk ...
	log.Printf("Memperbarui produk ID %d dengan data: %+v", params.ProductID, input)
	return c.JSON(fiber.Map{"message": "Produk berhasil diperbarui", "id": params.ProductID})
}

func main() {
	app := fiber.New()
	app.Get("/products/search", SearchProducts) // e.g., /products/search?category=books&max_price=50
	app.Put("/products/:id", UpdateProduct)     // e.g., PUT /products/123 dengan body JSON
	log.Fatal(app.Listen(":3000"))
}

```
Menggunakan tag struct (`query:`, `json:`, `form:`, `params:`, `xml:`, `header:`) memberi tahu parser cara memetakan nama field di request ke field struct Go Anda.

### Penanganan Request 📥

Bagian ini merangkum cara-cara spesifik untuk mendapatkan berbagai jenis data dari request client menggunakan `fiber.Ctx`.

#### Membaca Headers

Gunakan `c.Get("Header-Name")`. Nama header tidak case-sensitive.

```go
func HandleRequest(c *fiber.Ctx) error {
	userAgent := c.Get(fiber.HeaderUserAgent) // Konstanta Fiber untuk nama header umum
	apiKey := c.Get("X-API-Key")
	acceptHeader := c.Get("Accept")

	log.Printf("User-Agent: %s", userAgent)
	log.Printf("API Key: %s", apiKey)
	log.Printf("Accept: %s", acceptHeader)

	// Cek apakah request mengharapkan JSON
	if c.Accepts("application/json") != "" || c.Accepts("json") != "" {
		return c.JSON(fiber.Map{"message": "Anda meminta JSON"})
	}

	return c.SendString("Header diterima")
}
```

#### Membaca Query Parameters

Query parameter adalah bagian dari URL setelah tanda tanya (`?`), misalnya `/search?q=term&page=1`. Gunakan `c.Query("key")` atau `c.QueryParser(&struct)`.

```go
func SearchHandler(c *fiber.Ctx) error {
	// Cara manual
	searchTerm := c.Query("q")
	page := c.Query("page", "1") // Dengan nilai default "1"
	limit, err := c.QueryInt("limit", 10) // Parse ke int, default 10
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Parameter 'limit' harus angka")
	}

	log.Printf("Mencari '%s', Halaman: %s, Limit: %d", searchTerm, page, limit)

	// Cara dengan struct binding
	type SearchParams struct {
		Query    string `query:"q"`
		Page     int    `query:"page" default:"1"`
		Limit    int    `query:"limit" default:"10"`
		Sort     string `query:"sort"`
		ShowMeta bool   `query:"show_meta"`
	}
	var params SearchParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Query parameter tidak valid"})
	}
	log.Printf("Struct Binding - Mencari '%s', Halaman: %d, Limit: %d, Sort: '%s', Meta: %t",
		params.Query, params.Page, params.Limit, params.Sort, params.ShowMeta)

	// ... logika pencarian ...
	return c.JSON(fiber.Map{"results": "...", "params_used": params})
}
```
*Perhatikan:* Tag `default:"value"` dapat digunakan dalam struct `QueryParser` untuk memberikan nilai default jika query parameter tidak ada di URL.

#### Membaca Route Parameters

Parameter yang didefinisikan dalam path rute (e.g., `/users/:id`). Gunakan `c.Params("key")` atau `c.ParamsParser(&struct)`.

```go
type UserRouteParams struct {
	UserID int `params:"userId"` // Nama field harus cocok dengan :userId di definisi rute
}

// Rute: /users/:userId/orders/:orderId
func GetUserOrder(c *fiber.Ctx) error {
	// Cara manual
	userIdStr := c.Params("userId")
	orderIdStr := c.Params("orderId")
	log.Printf("Manual - User ID: %s, Order ID: %s", userIdStr, orderIdStr)

	// Cara dengan struct binding (hanya untuk UserID)
	var params UserRouteParams
	if err := c.ParamsParser(&params); err != nil {
		 return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID tidak valid"})
	}
	log.Printf("Struct Binding - User ID: %d", params.UserID)
	// Anda masih perlu mengambil orderId secara manual jika tidak di-bind
	orderId, err := c.ParamsInt("orderId")
	 if err != nil {
		 return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Order ID tidak valid"})
	}
	log.Printf("Struct Binding - Order ID: %d", orderId)


	// ... logika mengambil data order ...
	return c.JSON(fiber.Map{"user_id": params.UserID, "order_id": orderId})
}
```

#### Membaca Body Request

Untuk request seperti POST, PUT, PATCH yang membawa data di body. Cara paling umum adalah menggunakan `c.BodyParser(&struct)`.

```go
type CreatePostInput struct {
	Title   string   `json:"title" form:"title" validate:"required"`
	Content string   `json:"content" form:"content" validate:"required"`
	Tags    []string `json:"tags" form:"tags"` // Bisa array/slice
}

func CreatePostHandler(c *fiber.Ctx) error {
	var input CreatePostInput

	// BodyParser menangani JSON, Form (urlencoded/multipart), XML
	if err := c.BodyParser(&input); err != nil {
		log.Printf("Error parsing body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Gagal memproses body request",
			"details": err.Error(),
		})
	}

	// (Opsional tapi direkomendasikan) Lakukan validasi
	// validate := validator.New()
	// if err := validate.Struct(input); err != nil { ... return validation errors ...}

	log.Printf("Membuat post baru: Title='%s', Content='%s', Tags=%v", input.Title, input.Content, input.Tags)
	// ... logika menyimpan post ke database ...

	// Kembalikan data post yang baru dibuat (misalnya dengan ID)
	newPost := fiber.Map{
		"id": 1, // ID contoh dari DB
		"title": input.Title,
		"content": input.Content,
		"tags": input.Tags,
	}
	return c.Status(fiber.StatusCreated).JSON(newPost) // 201 Created
}
```

Jika Anda perlu body mentah (misalnya, untuk memproses webhook tanda tangan):

```go
func WebhookHandler(c *fiber.Ctx) error {
	rawBody := c.Body() // Mendapatkan []byte

	// Lakukan sesuatu dengan rawBody, misalnya verifikasi signature
	signature := c.Get("X-Webhook-Signature")
	if !verifySignature(rawBody, signature) {
		return c.Status(fiber.StatusUnauthorized).SendString("Signature tidak valid")
	}

	// Jika signature valid, Anda mungkin masih ingin parse body-nya
	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		// Gunakan fiber.Unmarshal jika body sudah dibaca
		// if err := fiber.Unmarshal(rawBody, &payload); err != nil {
		//     return c.Status(fiber.StatusBadRequest).SendString("Gagal parse payload JSON")
		// }
		// atau jika BodyParser gagal setelah c.Body() dipanggil
		return c.Status(fiber.StatusBadRequest).SendString("Gagal parse payload JSON setelah membaca body mentah")
	}

	log.Printf("Webhook diterima: %+v", payload)
	// ... proses event webhook ...

	return c.SendStatus(fiber.StatusOK) // Kirim 200 OK
}
```
**Perhatian:** Memanggil `c.Body()` akan membaca seluruh body ke memori. Jika Anda kemudian memanggil `c.BodyParser()`, parser mungkin tidak bisa membaca ulang body (tergantung implementasi internal Fiber/Fasthttp). Jika Anda perlu body mentah *dan* parsing, baca body mentah terlebih dahulu, lalu gunakan fungsi unmarshal yang sesuai (misalnya `fiber.Unmarshal` untuk JSON) pada byte mentah tersebut.

#### File Upload

File biasanya diunggah menggunakan `multipart/form-data`. Fiber memudahkan penanganan ini.

```go
func UploadFileHandler(c *fiber.Ctx) error {
	// 1. Dapatkan file header dari form field bernama "file_upload"
	fileHeader, err := c.FormFile("file_upload")
	if err != nil {
		log.Printf("Error mendapatkan file: %v", err)
		// Cek jika error karena field tidak ditemukan
		if err.Error() == "multipart: no such file" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Field 'file_upload' tidak ditemukan atau kosong"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses upload file"})
	}

	// 2. (Opsional) Dapatkan field form lain jika ada
	description := c.FormValue("description", "Tidak ada deskripsi")

	// 3. (Opsional) Validasi file (ukuran, tipe MIME)
	maxSize := int64(5 * 1024 * 1024) // 5 MB
	if fileHeader.Size > maxSize {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{"error": "Ukuran file melebihi batas 5MB"})
	}

	allowedMIMETypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"application/pdf": true,
	}
	// Deteksi tipe MIME dari header file (lebih aman daripada ekstensi)
	file, err := fileHeader.Open()
	if err != nil {
		 return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuka file"})
	}
	defer file.Close()
	buffer := make([]byte, 512) // Buffer untuk deteksi MIME
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		 return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membaca file"})
	}
	mimeType := http.DetectContentType(buffer)

	if !allowedMIMETypes[mimeType] {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"error": "Tipe file tidak didukung",
			"detected_mime": mimeType,
		})
	}

	// Reset file pointer ke awal sebelum menyimpan
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal reset file pointer"})
	}

	// 4. Tentukan path penyimpanan (buat unik jika perlu)
	//    PENTING: Jangan pernah gunakan fileHeader.Filename secara langsung sebagai path
	//    karena bisa mengandung karakter berbahaya (e.g., "../"). Sanitasi selalu!
	safeFilename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filepath.Base(fileHeader.Filename)) // Contoh nama unik
	savePath := filepath.Join("./uploads", safeFilename) // Simpan di direktori ./uploads

	// Pastikan direktori uploads ada
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		log.Printf("Error membuat direktori uploads: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyiapkan penyimpanan file"})
	}

	// 5. Simpan file ke disk menggunakan c.SaveFile atau io.Copy
	// c.SaveFile lebih mudah:
	// err = c.SaveFile(fileHeader, savePath)

	// Atau menggunakan io.Copy (memberi lebih banyak kontrol, gunakan file yang sudah dibuka)
	dst, err := os.Create(savePath)
	if err != nil {
		log.Printf("Error membuat file tujuan %s: %v", savePath, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan file"})
	}
	defer dst.Close()

	_, err = io.Copy(dst, file) // Salin dari file yang di-upload ke file tujuan
	if err != nil {
		log.Printf("Error menyimpan file %s: %v", savePath, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan file"})
	}

	log.Printf("File '%s' (deskripsi: '%s', size: %d, mime: %s) berhasil diupload ke %s",
		fileHeader.Filename, description, fileHeader.Size, mimeType, savePath)

	// Beri response sukses
	return c.JSON(fiber.Map{
		"message":       "File berhasil diupload!",
		"original_name": fileHeader.Filename,
		"saved_path":    savePath,
		"size":          fileHeader.Size,
		"mime_type":     mimeType,
		"description":   description,
	})
}

// Jangan lupa import "fmt", "path/filepath", "os", "time", "log", "net/http", "io"
```
**Poin Penting File Upload:**
*   Gunakan `c.FormFile()` untuk mendapatkan `*multipart.FileHeader`.
*   Gunakan `c.FormValue()` untuk mendapatkan field teks lain dalam form yang sama.
*   **Selalu validasi** ukuran file dan tipe MIME. Jangan percaya ekstensi file dari client. Gunakan `http.DetectContentType`.
*   **Sanitasi nama file** sebelum menyimpannya ke disk untuk mencegah *path traversal attack*. `filepath.Base()` membantu mengambil nama file saja. Buat nama file unik (misalnya dengan timestamp atau UUID) untuk menghindari konflik.
*   Gunakan `c.SaveFile()` untuk cara mudah menyimpan, atau `fileHeader.Open()` dan `io.Copy` untuk kontrol lebih.
*   Pastikan direktori tujuan ada (`os.MkdirAll`).
*   Atur batas ukuran body request di konfigurasi Fiber (`BodyLimit`) untuk mencegah serangan DoS dengan file besar.

### Penanganan Response 📤

Bagian ini merangkum cara mengirim berbagai jenis response kembali ke client.

#### Mengatur Status Code

Gunakan `c.SendStatus(code)` untuk mengirim hanya status, atau `c.Status(code).<SendMethod>(...)` untuk mengatur status sebelum mengirim body.

```go
app.Post("/items", func(c *fiber.Ctx) error {
	// ... logika membuat item ...
	newItem := Item{ID: 5, Name: "Item Baru"}
	return c.Status(fiber.StatusCreated).JSON(newItem) // 201 Created
})

app.Get("/items/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := findItemByID(id) // Fungsi imajiner
	if err != nil {
		// Jika item tidak ditemukan
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item tidak ditemukan"}) // 404 Not Found
	}
	return c.JSON(item) // 200 OK (default jika tidak diatur)
})

app.Delete("/items/:id", func(c *fiber.Ctx) error {
	// ... logika menghapus item ...
	return c.SendStatus(fiber.StatusNoContent) // 204 No Content (umum untuk DELETE sukses)
})
```

#### Mengatur Headers

Gunakan `c.Set(key, value)` atau `c.Append(key, value)`.

```go
app.Get("/data", func(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8) // text/plain; charset=utf-8
	c.Set("X-RateLimit-Limit", "100")
	c.Set("X-RateLimit-Remaining", "99")
	// Atur header Cache-Control
	c.Set(fiber.HeaderCacheControl, "public, max-age=3600") // Cache selama 1 jam
	return c.SendString("Ini data teks biasa dengan header kustom.")
})

app.Post("/login", func(c *fiber.Ctx) error {
	// ... validasi login ...
	c.Cookie(&fiber.Cookie{ // Mengatur cookie sesi
		Name:     "session",
		Value:    "rahasia-user-123",
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"message": "Login berhasil"})
})
```

#### Mengirim Berbagai Tipe Data

*   **String:** `c.SendString("...")` -> `text/plain`
*   **Byte Slice:** `c.Send([]byte{...})` -> `application/octet-stream` (atau terdeteksi)
*   **JSON:** `c.JSON(data)` -> `application/json`
*   **XML:** `c.XML(data)` -> `application/xml`
*   **HTML (dari Template):** `c.Render("template.html", data)` -> `text/html` (memerlukan setup template engine)
*   **File (inline):** `c.SendFile("./path/to/file")` -> Tipe MIME terdeteksi
*   **File (download):** `c.Download("./path/to/file", "nama_download.ext")` -> Tipe MIME terdeteksi + `Content-Disposition: attachment`
*   **Redirect:** `c.Redirect("/new/location", 302)`

```go
// Contoh mengirim response error JSON terstruktur
func GetResource(c *fiber.Ctx) error {
	if !userHasPermission(c) { // Fungsi imajiner
		// 403 Forbidden
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"code":    "ACCESS_DENIED",
			"message": "Anda tidak memiliki izin untuk mengakses sumber daya ini.",
		})
	}
	// ...
	return c.JSON(fiber.Map{"data": "..."})
}
```

### Penanganan Error 💣

Penanganan error yang baik sangat penting untuk aplikasi yang robust. Fiber menyediakan beberapa mekanisme untuk ini.

#### Mengembalikan Error dari Handler

Cara paling dasar adalah mengembalikan `error` dari handler atau middleware Anda.

```go
func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := database.FindItem(id) // Fungsi DB bisa mengembalikan error

	if err != nil {
		// Jika error adalah 'record not found'
		if errors.Is(err, gorm.ErrRecordNotFound) { // Contoh dengan GORM
			log.Printf("Item %s tidak ditemukan", id)
			// Kembalikan error Fiber standar untuk Not Found
			return fiber.ErrNotFound // Akan menghasilkan 404 Not Found
		}
		// Untuk error database lainnya
		log.Printf("Error database saat mencari item %s: %v", id, err)
		// Kembalikan error generik
		return fiber.ErrInternalServerError // Akan menghasilkan 500 Internal Server Error
		// atau error kustom
		// return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data item")
	}

	return c.JSON(item)
}
```
Fiber memiliki beberapa error pre-defined yang nyaman (seperti `fiber.ErrBadRequest`, `fiber.ErrNotFound`, `fiber.ErrUnauthorized`, dll.) yang langsung memetakan ke status code HTTP yang sesuai.

#### Error Kustom (`fiber.NewError`)

Jika error pre-defined tidak cukup, Anda bisa membuat error Fiber kustom dengan status code dan pesan spesifik menggunakan `fiber.NewError(statusCode, message)`.

```go
func ProcessPayment(c *fiber.Ctx) error {
	// ... logika validasi input ...
	if !inputValid {
		// Buat error 400 Bad Request dengan pesan spesifik
		return fiber.NewError(fiber.StatusBadRequest, "Data pembayaran tidak lengkap atau tidak valid.")
	}

	err := paymentGateway.Charge(...) // Fungsi imajiner
	if err != nil {
		// Tangani error spesifik dari payment gateway
		if errors.Is(err, paymentGateway.ErrInsufficientFunds) {
			// Buat error 402 Payment Required
			return fiber.NewError(fiber.StatusPaymentRequired, "Dana tidak mencukupi.")
		}
		// Error gateway lainnya
		log.Printf("Error payment gateway: %v", err)
		// Buat error 503 Service Unavailable
		return fiber.NewError(fiber.StatusServiceUnavailable, "Layanan pembayaran sedang tidak tersedia.")
	}

	return c.JSON(fiber.Map{"status": "Pembayaran berhasil"})
}
```

#### Custom Error Handler

Secara default, ketika handler mengembalikan `error` (atau `fiber.Error`), Fiber akan menangkapnya dan mengirim response HTTP yang sesuai (menggunakan status code dari `fiber.Error` atau default 500, dan pesan dari error).

Anda dapat **menyesuaikan sepenuhnya** bagaimana error ini diubah menjadi response HTTP dengan menyediakan `ErrorHandler` kustom di konfigurasi Fiber. Ini berguna untuk:

*   Memformat semua response error secara konsisten (misalnya, selalu dalam format JSON tertentu).
*   Menyembunyikan detail error internal di lingkungan produksi.
*   Melakukan logging error terpusat.
*   Mengirim error ke sistem monitoring (seperti Sentry, Datadog).

```go
package main

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Fungsi Error Handler Kustom
func MyCustomErrorHandler(c *fiber.Ctx, err error) error {
	// Default status code adalah 500
	code := fiber.StatusInternalServerError
	message := "Terjadi kesalahan internal pada server."

	// Cek apakah error adalah *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	// Log error internal secara detail (hanya di server)
	log.Printf("[ErrorHandler] Status: %d, Error: %v, Path: %s", code, err, c.Path())

	// Kirim error ke Sentry/Datadog dll di sini jika perlu

	// Jangan kirim detail error internal ke client di produksi
	isProduction := os.Getenv("APP_ENV") == "production"
	if isProduction && code == fiber.StatusInternalServerError {
		message = "Maaf, terjadi kesalahan tak terduga."
	}

	// Set Content-Type jika belum diatur
	// (Penting jika error terjadi sebelum Content-Type diatur oleh handler)
	if c.Get(fiber.HeaderContentType) == "" {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	}

	// Kirim response error JSON yang konsisten
	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"code":    code, // Atau kode error internal aplikasi Anda
		"message": message,
		// "details": err.Error(), // HINDARI ini di produksi
	})
}

func main() {
	app := fiber.New(fiber.Config{
		// Daftarkan error handler kustom kita
		ErrorHandler: MyCustomErrorHandler,
	})

	// Penting: Middleware Recover sebaiknya tetap digunakan untuk menangkap panic
	// sebelum ErrorHandler kustom dipanggil.
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// Handler ini sukses
		return c.SendString("OK")
	})

	app.Get("/notfound", func(c *fiber.Ctx) error {
		// Mengembalikan error Fiber standar
		return fiber.ErrNotFound // Akan ditangani oleh MyCustomErrorHandler -> 404 JSON
	})

	app.Get("/badrequest", func(c *fiber.Ctx) error {
		// Mengembalikan error kustom
		return fiber.NewError(fiber.StatusBadRequest, "Parameter 'q' dibutuhkan.") // -> 400 JSON
	})

	app.Get("/dberror", func(c *fiber.Ctx) error {
		// Mensimulasikan error non-Fiber
		simulatedError := errors.New("database connection failed")
		return simulatedError // Akan ditangani oleh MyCustomErrorHandler -> 500 JSON
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Sesuatu yang sangat salah terjadi!") // Akan ditangkap Recover, lalu ke ErrorHandler -> 500 JSON
	})


	log.Fatal(app.Listen(":3000"))
}
```

Dengan `ErrorHandler` kustom, semua error yang dikembalikan dari handler (termasuk panic yang ditangkap `recover`) akan melewati fungsi ini, memberi Anda kontrol penuh atas response error akhir.

#### Middleware Recover

Middleware `recover` (dari `github.com/gofiber/fiber/v2/middleware/recover`) sangat penting. Fungsinya adalah untuk menangkap *panic* yang mungkin terjadi di handler atau middleware Anda.

Tanpa `recover`, panic akan menyebabkan server crash. Dengan `recover`, panic akan ditangkap, diubah menjadi error (biasanya `500 Internal Server Error`), dan kemudian diproses oleh `ErrorHandler` (default atau kustom).

**Selalu gunakan `recover.New()` sebagai salah satu middleware pertama Anda secara global.**

```go
app := fiber.New(fiber.Config{ErrorHandler: MyCustomErrorHandler})

// Recover harus didaftarkan SEBELUM middleware/handler lain yang mungkin panic
app.Use(recover.New(recover.Config{
	EnableStackTrace: true, // Tampilkan stack trace di log (berguna saat development)
}))

// Middleware dan Rute lainnya...
app.Get("/panic-now", func(c *fiber.Ctx) error {
	myMap := map[string]string{}
	// Ini akan menyebabkan panic (nil map dereference)
	myMap["key"] = "value" // <-- PANIC!
	return c.SendString("Tidak akan sampai sini")
})
```

### Konfigurasi (`fiber.Config`)

Saat membuat instance aplikasi Fiber dengan `fiber.New()`, Anda dapat meneruskan struct `fiber.Config` untuk menyesuaikan berbagai aspek perilakunya.

```go
config := fiber.Config{
	// Prefork men-spawn beberapa proses Go yang mendengarkan di port yang sama.
	// Berguna untuk memanfaatkan semua core CPU tanpa perlu logic Goroutine manual.
	// Tidak kompatibel dengan beberapa stateful middleware (misalnya, session default).
	// Default: false
	Prefork: os.Getenv("APP_ENV") == "production", // Aktifkan hanya di produksi

	// Nama aplikasi, muncul di header 'Server' jika ServerHeader tidak diatur.
	AppName: "My Awesome App v1.1",

	// Mengganti header 'Server' default ('Fiber'). Kosongkan untuk menyembunyikannya.
	ServerHeader: "MyWebServer",
	// ServerHeader: "", // Sembunyikan header Server

	// Routing yang ketat. Jika true, '/foo' dan '/foo/' dianggap berbeda.
	// Default: false
	StrictRouting: false,

	// Routing case-sensitive. Jika true, '/Foo' dan '/foo' berbeda.
	// Default: false
	CaseSensitive: false,

	// Ukuran maksimum body request dalam byte. Mencegah DoS.
	// Default: 4 * 1024 * 1024 (4MB)
	BodyLimit: 10 * 1024 * 1024, // 10 MB

	// Konfigurasi untuk template engine (lihat bagian Template Engine).
	// Views: ...,

	// Custom error handler (lihat bagian Penanganan Error).
	ErrorHandler: MyCustomErrorHandler,

	// Jumlah maksimum header request yang diizinkan.
	// Default: 1024
	ReadBufferSize: 8192, // Tingkatkan jika perlu menangani header besar

	// Timeout untuk membaca seluruh request (termasuk body).
	// Default: Tidak ada batas (mengandalkan timeout OS)
	ReadTimeout: 5 * time.Second,

	// Timeout untuk menulis response.
	// Default: Tidak ada batas
	WriteTimeout: 10 * time.Second,

	// Timeout untuk koneksi idle (keep-alive).
	// Default: Tidak ada batas (mengandalkan timeout OS)
	IdleTimeout: 60 * time.Second,

	// Konfigurasi lainnya... (lihat dokumentasi fiber.Config)
	// DisableKeepalive: false,
	// ReduceMemoryUsage: false, // Bisa mengurangi memori tapi mungkin sedikit lebih lambat
	// GETOnly: false, // Hanya izinkan metode GET
	// EnablePrintRoutes: true, // Cetak rute saat startup
	// Network: "tcp", // Protokol jaringan (tcp, tcp4, tcp6)
}

app := fiber.New(config)
```

Pilih konfigurasi yang sesuai dengan kebutuhan aplikasi Anda, terutama `Prefork`, `BodyLimit`, `ErrorHandler`, dan `Timeouts` untuk aplikasi produksi.

#### Konfigurasi Prefork

Mode `Prefork` adalah fitur unik Fiber (memanfaatkan fitur SO_REUSEPORT di Linux/BSD). Ketika `Prefork: true`, Fiber akan:

1.  Membuat *child process* sebanyak jumlah core CPU yang tersedia (atau sesuai `runtime.GOMAXPROCS(0)`).
2.  Setiap child process menjalankan instance aplikasi Fiber yang identik.
3.  Semua child process mendengarkan pada *port yang sama*.
4.  Kernel sistem operasi akan mendistribusikan koneksi masuk ke salah satu child process yang tersedia (load balancing level kernel).

**Keuntungan:**

*   **Pemanfaatan Multi-core Otomatis:** Cara mudah untuk membuat aplikasi Anda menggunakan semua core CPU tanpa perlu mengelola Goroutine secara eksplisit untuk menangani request secara paralel.
*   **Potensi Throughput Lebih Tinggi:** Bisa meningkatkan jumlah request per detik yang dapat ditangani.
*   **Isolasi:** Jika satu child process crash karena panic (meskipun `recover` harusnya menangkapnya), child process lain tetap berjalan dan melayani request.

**Kekurangan/Pertimbangan:**

*   **Stateful Middleware:** Middleware yang menyimpan state di memori proses (seperti middleware `session` default Fiber yang berbasis memori) **tidak akan berfungsi dengan benar**. Setiap child process memiliki memorinya sendiri, sehingga sesi yang dibuat di satu proses tidak akan dikenali oleh proses lain. Anda perlu menggunakan penyimpanan state eksternal (seperti Redis, Memcached, database) untuk session, rate limiting, dll., jika menggunakan prefork.
*   **Debugging:** Sedikit lebih kompleks untuk di-debug karena Anda memiliki banyak proses.
*   **Overhead:** Memulai banyak proses memakan lebih banyak memori daripada satu proses dengan banyak Goroutine.
*   **Hanya Linux/BSD:** Fitur SO_REUSEPORT tidak tersedia secara luas atau bekerja dengan cara yang sama di Windows atau macOS.

**Kapan Menggunakan Prefork?**

*   Aplikasi *stateless* (tidak bergantung pada state memori antar request).
*   Aplikasi yang CPU-bound dan perlu memanfaatkan semua core.
*   Berjalan di lingkungan Linux/BSD di produksi.

**Kapan Menghindari Prefork?**

*   Menggunakan middleware stateful berbasis memori.
*   Perlu state global yang dibagi antar semua handler (perlu sinkronisasi antar proses jika pakai prefork, lebih mudah dengan Goroutine di satu proses).
*   Berjalan di Windows/macOS.
*   Aplikasi I/O-bound (menunggu database, network) mungkin tidak mendapat banyak manfaat dibandingkan Goroutine biasa.

Jika ragu, mulai tanpa `Prefork` (`false`) dan aktifkan nanti jika benchmarking menunjukkan manfaat yang signifikan dan Anda siap menangani implikasi state.

### Template Engine 📄

Fiber memungkinkan Anda merender HTML dinamis menggunakan berbagai template engine Go. Fiber sendiri tidak menyertakan engine, tetapi menyediakan *interface* (`fiber.Views`) yang dapat diimplementasikan oleh adapter engine.

#### Konsep Template Engine

Template engine memisahkan logika presentasi (HTML) dari logika bisnis (Go). Anda membuat file template (misalnya `.html`, `.tmpl`, `.pug`) yang berisi markup HTML dicampur dengan sintaks khusus engine untuk menampilkan data dinamis, melakukan perulangan, kondisional, dll.

#### Menggunakan Template Engine Bawaan (HTML)

Go memiliki package `html/template` bawaan yang aman untuk HTML (melakukan escaping otomatis untuk mencegah XSS). Fiber menyediakan adapter untuk ini.

1.  **Instal Adapter:**
	```bash
	go get -u github.com/gofiber/template/html/v2
	```

2.  **Konfigurasi di Fiber:**
	Buat direktori untuk menyimpan file template Anda (misalnya, `views`). Buat file template sederhana, misal `views/index.html`:
	```html
	<!-- views/index.html -->
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Title}}</title> <!-- Menampilkan data 'Title' -->
	</head>
	<body>
		<h1>{{.Header}}</h1>
		<p>Selamat datang di halaman contoh!</p>

		<h2>Daftar Item:</h2>
		{{if .Items}} <!-- Cek jika Items ada dan tidak kosong -->
			<ul>
				{{range .Items}} <!-- Looping melalui slice Items -->
					<li>{{.}}</li> <!-- Tampilkan setiap item -->
				{{end}}
			</ul>
		{{else}}
			<p>Tidak ada item untuk ditampilkan.</p>
		{{end}}
	</body>
	</html>
	```

	Konfigurasikan engine di `main.go`:
	```go
	package main

	import (
		"log"

		"github.com/gofiber/fiber/v2"
		"github.com/gofiber/template/html/v2" // Import adapter
	)

	func main() {
		// 1. Buat instance engine, arahkan ke direktori views
		//    Direkomendasikan menggunakan Reload: true saat development
		//    agar perubahan template langsung terlihat tanpa restart server.
		engine := html.New("./views", ".html") // Cari file .html di ./views
		engine.Reload(true) // Aktifkan reload saat development
		engine.Debug(true) // Aktifkan debug logging saat development

		// 2. Buat aplikasi Fiber dengan engine yang dikonfigurasi
		app := fiber.New(fiber.Config{
			Views: engine, // Beritahu Fiber untuk menggunakan engine ini
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
			return c.Render("index", data)
		})

		app.Get("/kosong", func(c *fiber.Ctx) error {
			// Contoh tanpa item
			data := fiber.Map{
				"Title":  "Halaman Kosong",
				"Header": "Tidak ada item",
				"Items":  nil, // Atau []string{}
			}
			return c.Render("index", data)
		})

		log.Fatal(app.Listen(":3000"))
	}
	```

3.  **Jalankan dan Uji:** Jalankan `go run main.go` dan buka `http://localhost:3000` di browser. Anda akan melihat HTML yang dirender dengan data dari handler. Akses `http://localhost:3000/kosong` untuk melihat kondisi `else`.

#### Menggunakan Template Engine Lain

Fiber mendukung banyak engine populer lainnya melalui adapter terpisah (biasanya di `github.com/gofiber/template/...`). Contoh: Pug, Amber, Handlebars, Jet.

Misalnya, menggunakan **Handlebars**:

1.  **Instal Adapter:**
	```bash
	go get -u github.com/gofiber/template/handlebars/v2
	```
2.  **Buat Template Handlebars:** (misal `views/profile.hbs`)
	```handlebars
	<!-- views/profile.hbs -->
	<!DOCTYPE html>
	<html>
	<head>
		<title>Profil {{user.Name}}</title>
	</head>
	<body>
		<h1>Profil Pengguna</h1>
		<p>ID: {{user.ID}}</p>
		<p>Nama: {{user.Name}}</p>
		<p>Email: {{user.Email}}</p>

		{{#if isAdmin}}
			<p><strong>Status: Administrator</strong></p>
		{{else}}
			<p>Status: Pengguna Biasa</p>
		{{/if}}
	</body>
	</html>
	```
3.  **Konfigurasi di Fiber:**
	```go
	// ... import lainnya ...
	import (
		"github.com/gofiber/template/handlebars/v2" // Import adapter handlebars
	)

	func main() {
		// Buat engine Handlebars
		engine := handlebars.New("./views", ".hbs") // Ekstensi .hbs
		engine.Reload(true)

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
			return c.Render("profile", fiber.Map{
				"user": userData,
				"isAdmin": isAdmin,
			})
		})

		log.Fatal(app.Listen(":3000"))
	}
	```

Prosesnya mirip untuk engine lain: instal adapter, buat template dengan sintaks engine tersebut, dan konfigurasikan engine di `fiber.Config`.

#### Layouts

Banyak template engine (termasuk `html/template` dengan fungsi `define`/`template` dan adapter Handlebars) mendukung konsep *layout*. Layout adalah template dasar (kerangka HTML) yang mendefinisikan struktur umum halaman (header, footer, sidebar), dan konten spesifik per halaman disisipkan ke dalamnya.

Menggunakan layout mengurangi duplikasi kode HTML di banyak template.

Contoh dengan `html/template`:

1.  **Buat Layout (`views/layouts/main.html`):**
	```html
	<!-- views/layouts/main.html -->
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{template "title" .}} - Website Saya</title> <!-- Panggil blok title -->
		<link rel="stylesheet" href="/static/css/style.css">
	</head>
	<body>
		<header>Ini Header Umum</header>
		<main>
			{{template "content" .}} <!-- Panggil blok content -->
		</main>
		<footer>Ini Footer Umum</footer>
	</body>
	</html>
	```

2.  **Buat Template Konten (`views/about.html`):**
	```html
	<!-- views/about.html -->
	{{define "title"}}Tentang Kami{{end}} <!-- Definisikan blok title -->

	{{define "content"}} <!-- Definisikan blok content -->
		<h2>Tentang Perusahaan Kami</h2>
		<p>Ini adalah halaman tentang.</p>
	{{end}}
	```

3.  **Konfigurasi Engine dan Render dengan Layout:**
	```go
	// ... import ...
	import "github.com/gofiber/template/html/v2"

	func main() {
		// Muat semua template .html dari direktori views
		engine := html.New("./views", ".html")
		engine.Reload(true)

		app := fiber.New(fiber.Config{ Views: engine })

		app.Get("/about", func(c *fiber.Ctx) error {
			// Saat merender 'about', juga teruskan nama file layout 'main'
			// Data akan tersedia di kedua template (layout dan content)
			return c.Render("about", fiber.Map{}, "layouts/main")
		})

		// Rute lain bisa menggunakan layout yang sama
		app.Get("/contact", func(c *fiber.Ctx) error {
			// Buat views/contact.html serupa dengan about.html
			return c.Render("contact", fiber.Map{"Email": "info@example.com"}, "layouts/main")
		})

		log.Fatal(app.Listen(":3000"))
	}
	```

Saat `c.Render("about", data, "layouts/main")` dipanggil:
*   Fiber (melalui engine) akan memuat `views/about.html` dan `views/layouts/main.html`.
*   Engine akan mengeksekusi `layouts/main.html` sebagai template dasar.
*   Ketika `{{template "content" .}}` ditemukan di layout, engine akan mencari blok `{{define "content"}}` di `about.html` dan merendernya di sana. Hal yang sama berlaku untuk `{{template "title" .}}`.
*   Data (`fiber.Map{}`) diteruskan ke kedua template.

Adapter template engine lain mungkin memiliki cara berbeda untuk menangani layout (misalnya, Handlebars menggunakan Partials dan Helpers). Konsultasikan dokumentasi adapter spesifik.

### Menyajikan File Statis 📁

Aplikasi web hampir selalu perlu menyajikan file statis seperti CSS, JavaScript, gambar, font, dll. Fiber menyediakan middleware `Static` yang efisien untuk ini.

```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Buat direktori 'public' di root proyek Anda
	// Taruh file CSS, JS, gambar di dalamnya
	// Contoh: ./public/css/style.css
	//         ./public/js/script.js
	//         ./public/images/logo.png

	// Daftarkan middleware Static
	// Argumen pertama adalah prefix URL
	// Argumen kedua adalah path direktori di sistem file
	app.Static("/static", "./public")
	// Sekarang:
	// Request ke /static/css/style.css akan menyajikan file ./public/css/style.css
	// Request ke /static/images/logo.png akan menyajikan file ./public/images/logo.png

	// Anda bisa mendaftarkan beberapa direktori statis
	app.Static("/assets", "./assets") // Sajikan file dari ./assets di bawah URL /assets

	// Menyajikan file dari root URL (misalnya, favicon.ico atau index.html)
	// Gunakan prefix "/"
	app.Static("/", "./root_files")
	// Request ke /favicon.ico akan menyajikan ./root_files/favicon.ico

	// Rute aplikasi Anda lainnya
	app.Get("/", func(c *fiber.Ctx) error {
		// Contoh HTML yang mereferensikan file statis
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>App Fiber</title>
			<link rel="stylesheet" href="/static/css/style.css">
			<link rel="icon" href="/favicon.ico">
		</head>
		<body>
			<h1>Selamat Datang!</h1>
			<img src="/static/images/logo.png" alt="Logo">
			<script src="/static/js/script.js"></script>
		</body>
		</html>
		`
		// Kirim sebagai HTML
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendString(html)
	})

	log.Fatal(app.Listen(":3000"))
}
```

**Konfigurasi `Static`:**

Middleware `Static` dapat dikonfigurasi lebih lanjut:

```go
app.Static("/downloadables", "./files_to_download", fiber.Static{
	// Kompresi file statis (gzip, brotli) jika client mendukung.
	// Default: false (di v2. Awalnya true di versi lama)
	Compress: true,

	// Izinkan request byte range (penting untuk streaming video/audio).
	// Default: false (di v2)
	ByteRange: true,

	// Izinkan penjelajahan direktori (menampilkan daftar file jika tidak ada index.html).
	// Default: false (Jangan aktifkan di produksi kecuali benar-benar dibutuhkan!)
	Browse: false,

	// Nama file index default yang dicari saat mengakses direktori.
	// Default: "index.html"
	Index: "default.html",

	// Durasi Cache-Control max-age (detik). 0 berarti tidak ada cache.
	// Default: 0
	MaxAge: 3600, // Cache selama 1 jam
})
```

**Penting:**
*   Tempatkan middleware `Static` **sebelum** definisi rute Anda jika ada kemungkinan konflik path (misalnya, jika Anda memiliki rute `/static/users` dan juga direktori statis `/static`). Fiber akan mencoba mencocokkan file statis terlebih dahulu.
*   Pastikan path direktori yang Anda berikan ke `app.Static` benar relatif terhadap lokasi di mana Anda menjalankan binary aplikasi Go Anda.

### Validasi Request ✅

Memvalidasi data yang masuk dari client (body, query params, route params) adalah langkah krusial untuk keamanan dan integritas data aplikasi Anda. Jangan pernah percaya input dari client!

#### Pentingnya Validasi

*   **Keamanan:** Mencegah serangan seperti SQL Injection, Cross-Site Scripting (XSS), dan exploit lainnya yang memanfaatkan data tidak valid.
*   **Integritas Data:** Memastikan data yang disimpan di database atau diproses oleh aplikasi Anda memiliki format dan nilai yang benar.
*   **Pengalaman Pengguna:** Memberikan feedback yang jelas kepada pengguna jika input mereka salah.
*   **Stabilitas Aplikasi:** Mencegah error atau panic akibat data yang tidak terduga.

#### Menggunakan Library Validator

Fiber tidak menyertakan library validasi bawaan, tetapi sangat mudah diintegrasikan dengan library populer seperti [`go-playground/validator`](https://github.com/go-playground/validator). Ini adalah library yang sangat kuat dan banyak digunakan di ekosistem Go.

1.  **Instal Validator:**
	```bash
	go get github.com/go-playground/validator/v10
	```

2.  **Tambahkan Tag Validasi ke Struct:**
	Gunakan tag `validate` pada field struct yang ingin Anda validasi. Library ini mendukung banyak aturan validasi bawaan (required, email, url, min, max, len, uuid, dll.) dan memungkinkan validasi kustom.

	```go
	type RegisterUserInput struct {
		Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Age      int    `json:"age" validate:"omitempty,gte=18,lte=120"` // gte=greater than or equal, lte=less than or equal
		Website  string `json:"website" validate:"omitempty,url"` // omitempty: validasi hanya jika field tidak kosong
		UserType string `json:"user_type" validate:"required,oneof=admin user guest"` // Harus salah satu dari nilai ini
	}

	type ProductFilter struct {
		Category string `query:"category" validate:"omitempty,alpha"` // Hanya huruf
		MaxPrice *int   `query:"max_price" validate:"omitempty,gt=0"` // Pointer agar bisa null/kosong, validasi jika ada
		Page     int    `query:"page" validate:"omitempty,min=1"`
	}
	```
	Lihat [dokumentasi `go-playground/validator`](https://pkg.go.dev/github.com/go-playground/validator/v10) untuk daftar lengkap tag yang tersedia.

#### Contoh Implementasi

Anda perlu membuat instance validator dan memanggil metode `Struct()` setelah mem-binding data request.

```go
package main

import (
	"log"
	"strings"

	"github.com/go-playground/validator/v10" // Import validator
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Global instance validator (atau buat per request jika perlu konfigurasi berbeda)
var validate = validator.New()

// Struct untuk response error validasi yang lebih informatif
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

// Fungsi helper untuk memformat error validasi
func formatValidationErrors(err error) []ValidationErrorResponse {
	var errors []ValidationErrorResponse

	// Cek apakah error adalah tipe ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			// Buat pesan error yang lebih user-friendly (contoh sederhana)
			var message string
			switch fieldErr.Tag() {
			case "required":
				message = "Field ini wajib diisi."
			case "email":
				message = "Format email tidak valid."
			case "min":
				 message = "Minimal harus " + fieldErr.Param() + " karakter/nilai."
			case "max":
				 message = "Maksimal harus " + fieldErr.Param() + " karakter/nilai."
			 case "alphanum":
				 message = "Hanya boleh berisi huruf dan angka."
			case "oneof":
				message = "Harus salah satu dari: " + strings.Replace(fieldErr.Param(), " ", ", ", -1)
			default:
				message = "Field tidak valid (" + fieldErr.Tag() + ")"
			}

			errors = append(errors, ValidationErrorResponse{
				Field:   fieldErr.Field(), // Nama field struct
				Tag:     fieldErr.Tag(),   // Tag validasi yang gagal
				Message: message,          // Pesan kustom
			})
		}
	} else {
		// Jika error bukan ValidationErrors (jarang terjadi jika inputnya err dari validate.Struct)
		log.Printf("Warning: Error validasi tidak terduga: %v", err)
		errors = append(errors, ValidationErrorResponse{Message: "Error validasi tidak dikenal"})
	}
	return errors
}

// Handler untuk registrasi pengguna
func RegisterUserHandler(c *fiber.Ctx) error {
	input := new(RegisterUserInput) // Gunakan new() agar dapat pointer

	// 1. Bind body request ke struct
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Gagal memproses body request.",
			"details": err.Error(),
		})
	}

	// 2. Lakukan validasi pada struct yang sudah di-bind
	err := validate.Struct(input)
	if err != nil {
		// Jika validasi gagal, format error dan kirim 400 Bad Request
		validationErrors := formatValidationErrors(err)
		log.Printf("Validasi gagal untuk registrasi: %v", validationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Data yang diberikan tidak valid.",
			"errors":  validationErrors, // Kirim detail error validasi
		})
	}

	// 3. Jika validasi berhasil, lanjutkan proses
	log.Printf("Registrasi valid diterima: Username=%s, Email=%s, Age=%d", input.Username, input.Email, input.Age)
	// ... logika menyimpan pengguna ke database ...

	// Kirim response sukses
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Pengguna berhasil diregistrasi.",
		"user": fiber.Map{ // Jangan kirim password kembali!
			"username": input.Username,
			"email":    input.Email,
			"age": input.Age,
			"website": input.Website,
			"user_type": input.UserType,
		},
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/register", RegisterUserHandler)

	log.Fatal(app.Listen(":3000"))
}
```

**Cara Menguji:**
*   **Request Valid:**
	```bash
	curl -X POST http://localhost:3000/register -H "Content-Type: application/json" -d '{"username": "johndoe", "email": "john.doe@example.com", "password": "password123", "age": 25, "user_type": "user"}'
	```
	*(Harusnya mengembalikan 201 Created)*
*   **Request Tidak Valid (beberapa error):**
	```bash
	curl -X POST http://localhost:3000/register -H "Content-Type: application/json" -d '{"username": "jo", "email": "john.doe@", "password": "pass", "age": 15, "user_type": "superadmin"}'
	```
	*(Harusnya mengembalikan 400 Bad Request dengan daftar error validasi di body JSON)*

**Tips Validasi:**

*   **Validasi Sedini Mungkin:** Lakukan validasi segera setelah menerima data.
*   **Jangan Hanya Bergantung pada Validasi Frontend:** Validasi di browser (JavaScript) bagus untuk UX, tetapi *harus* divalidasi ulang di backend karena client dapat mem-bypass validasi frontend.
*   **Berikan Pesan Error yang Jelas:** Bantu pengguna memperbaiki input mereka. Format response error validasi yang konsisten.
*   **Gunakan Pointer untuk Field Opsional:** Jika sebuah field (seperti `Age` atau `Website`) bersifat opsional *dan* memiliki aturan validasi (seperti `gte=18` atau `url`), gunakan pointer (`*int`, `*string`). Dengan begitu, jika client tidak mengirim field tersebut, nilainya akan `nil`, dan tag `omitempty` akan melewatkan validasi. Jika menggunakan tipe non-pointer ( `int`, `string`), nilai default Go (0 atau "") akan divalidasi, yang mungkin tidak diinginkan.
*   **Validasi Kustom:** `go-playground/validator` memungkinkan Anda mendaftarkan fungsi validasi kustom untuk logika yang lebih kompleks (misalnya, memeriksa apakah username sudah ada di database).

---

## 6. Topik Lanjutan 🧭

Setelah menguasai dasar-dasarnya, mari jelajahi beberapa fitur dan konsep yang lebih canggih di Fiber dan pengembangan web Go.

### WebSocket

WebSocket menyediakan saluran komunikasi dua arah (full-duplex) melalui satu koneksi TCP. Ini sangat berguna untuk aplikasi real-time seperti chat, notifikasi langsung, game online, dll.

Fiber menyediakan package middleware WebSocket yang mudah digunakan: `github.com/gofiber/contrib/websocket`.

1.  **Instal:**
	```bash
	go get github.com/gofiber/contrib/websocket
	```

2.  **Contoh Penggunaan:**
	```go
	package main

	import (
		"log"

		"github.com/gofiber/contrib/websocket" // Import websocket
		"github.com/gofiber/fiber/v2"
		"github.com/gofiber/fiber/v2/middleware/logger"
	)

	func main() {
		app := fiber.New()
		app.Use(logger.New())

		// Middleware untuk memastikan request adalah upgrade WebSocket
		app.Use("/ws", func(c *fiber.Ctx) error {
			// Periksa apakah header menunjukkan permintaan upgrade WebSocket
			if websocket.IsWebSocketUpgrade(c) {
				c.Locals("allowed", true)
				return c.Next()
			}
			// Jika bukan permintaan WebSocket, kirim 426 Upgrade Required
			return fiber.ErrUpgradeRequired
		})

		// Handler untuk koneksi WebSocket di path /ws/:id
		app.Get("/ws/:id", websocket.New(func(conn *websocket.Conn) {
			// conn adalah *websocket.Conn yang membungkus koneksi
			// Dapatkan parameter dari URL asli (sebelum upgrade)
			id := conn.Params("id")
			log.Printf("WebSocket terhubung untuk ID: %s dari %s", id, conn.RemoteAddr())

			// Variabel untuk tipe pesan, pesan, dan error
			var (
				mt  int
				msg []byte
				err error
			)

			// Loop tak terbatas untuk membaca pesan dari client
			for {
				// Baca pesan dari client
				// conn.ReadMessage() adalah blocking call
				if mt, msg, err = conn.ReadMessage(); err != nil {
					// Jika ada error (koneksi ditutup, dll.), log dan keluar dari loop
					log.Println("Error membaca pesan:", err)
					break // Keluar dari loop, menutup koneksi di sisi server
				}

				log.Printf("Pesan diterima dari ID %s: %s (Tipe: %d)", id, msg, mt)

				// Kirim pesan kembali ke client (echo)
				// Anda bisa mengirim pesan teks (websocket.TextMessage) atau biner (websocket.BinaryMessage)
				if err = conn.WriteMessage(mt, msg); err != nil {
					log.Println("Error menulis pesan:", err)
					break // Keluar jika gagal menulis
				}
			}
			// Kode setelah loop akan dieksekusi saat koneksi ditutup (baik oleh client atau server)
			log.Printf("WebSocket terputus untuk ID: %s", id)
			// Lakukan pembersihan jika perlu (misalnya, hapus user dari daftar online)
		}))

		log.Fatal(app.Listen(":3000"))
	}
	```

**Penjelasan:**
*   `websocket.IsWebSocketUpgrade(c)`: Memeriksa header `Connection: Upgrade` dan `Upgrade: websocket`.
*   `websocket.New(handler)`: Middleware utama yang menangani handshake upgrade dan memanggil fungsi `handler` Anda setelah koneksi berhasil dibuat. `handler` menerima `*websocket.Conn`.
*   `*websocket.Conn`: Objek untuk berinteraksi dengan koneksi WebSocket. Metode utamanya adalah:
	*   `ReadMessage()`: Membaca pesan masuk (blocking).
	*   `WriteMessage(messageType int, data []byte)`: Mengirim pesan.
	*   `WriteJSON(v interface{})`, `ReadJSON(v interface{})`: Mengirim/menerima data dalam format JSON.
	*   `Close()`: Menutup koneksi.
	*   `LocalAddr()`, `RemoteAddr()`: Mendapatkan alamat lokal/remote.
	*   `Params(key string)`, `Query(key string)`, `Cookies(key string)`, `Locals(key string)`: Mengakses data dari konteks HTTP *sebelum* upgrade terjadi.
*   Loop `for {}`: Pola umum untuk terus menerus membaca pesan dari client sampai koneksi ditutup atau terjadi error.

**Mengelola Banyak Koneksi:**
Untuk aplikasi chat atau notifikasi, Anda perlu cara untuk melacak semua koneksi aktif dan mengirim pesan ke koneksi tertentu atau ke semua koneksi (broadcast). Ini biasanya melibatkan penggunaan map (dengan mutex untuk konkurensi) atau channel Go.

```go
// Contoh (sangat sederhana) manajemen koneksi untuk broadcast
var clients = make(map[*websocket.Conn]bool) // Map koneksi aktif
var register = make(chan *websocket.Conn)     // Channel untuk mendaftarkan koneksi baru
var broadcast = make(chan []byte)             // Channel untuk pesan broadcast
var unregister = make(chan *websocket.Conn)   // Channel untuk menghapus koneksi

func runHub() {
	for {
		select {
		case conn := <-register:
			clients[conn] = true
			log.Println("Koneksi terdaftar:", conn.RemoteAddr())
		case message := <-broadcast:
			log.Println("Broadcasting pesan:", string(message))
			// Kirim pesan ke semua client yang terdaftar
			for conn := range clients {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Println("Error broadcast ke", conn.RemoteAddr(), err)
					// Koneksi error, hapus dari daftar
					// (Perlu mekanisme unregister yang lebih baik dalam produksi)
					conn.Close()
					delete(clients, conn)
				}
			}
		case conn := <-unregister:
			// Hapus koneksi dari map jika masih ada
			if _, ok := clients[conn]; ok {
				log.Println("Koneksi dihapus:", conn.RemoteAddr())
				delete(clients, conn)
				conn.Close() // Pastikan ditutup
			}
		}
	}
}

// Di main.go
func main() {
	// ... setup app fiber ...

	// Jalankan hub di goroutine terpisah
	go runHub()

	app.Get("/ws", websocket.New(func(conn *websocket.Conn) {
		// 1. Daftarkan koneksi baru
		register <- conn
		log.Printf("Koneksi WebSocket baru dari: %s", conn.RemoteAddr())

		// 2. Pastikan koneksi dihapus saat fungsi handler selesai (koneksi ditutup)
		defer func() {
			unregister <- conn
		}()

		// 3. Loop membaca pesan dari client ini
		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					 log.Printf("Error baca websocket: %v", err)
				} else {
					log.Printf("Koneksi websocket ditutup normal: %s", conn.RemoteAddr())
				}
				break // Keluar loop jika error/tutup
			}
		   if mt == websocket.TextMessage {
				// Kirim pesan yang diterima ke channel broadcast
				log.Printf("Pesan dari %s: %s", conn.RemoteAddr(), string(msg))
				broadcast <- msg
			} else {
				 log.Printf("Menerima tipe pesan non-teks: %d", mt)
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}

// Jangan lupa import "log", "github.com/gofiber/contrib/websocket", "github.com/gofiber/fiber/v2"
```
Ini adalah contoh dasar. Aplikasi produksi memerlukan penanganan error yang lebih baik, struktur data yang lebih efisien (terutama untuk banyak koneksi), dan mungkin pemisahan logic hub ke package tersendiri.

### Server-Sent Events (SSE)

SSE adalah standar web yang memungkinkan server mengirim pembaruan (event) ke client melalui koneksi HTTP yang tetap terbuka. Berbeda dengan WebSocket, SSE bersifat *satu arah* (server-ke-client). Ini lebih sederhana dari WebSocket dan cocok untuk kasus seperti feed berita langsung, pembaruan status, notifikasi, dll., di mana komunikasi dua arah tidak diperlukan.

Fiber tidak memiliki middleware SSE bawaan khusus seperti WebSocket, tetapi SSE mudah diimplementasikan menggunakan handler biasa karena dasarnya adalah response HTTP dengan `Content-Type: text/event-stream` dan format pesan tertentu.

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp" // Import fasthttp untuk Stream
)

func sseHandler(c *fiber.Ctx) error {
	// 1. Set header yang diperlukan untuk SSE
	c.Set(fiber.HeaderContentType, "text/event-stream")
	c.Set(fiber.HeaderCacheControl, "no-cache")
	c.Set(fiber.HeaderConnection, "keep-alive")
	c.Set(fiber.HeaderTransferEncoding, "chunked") // Atau pastikan response tidak di-buffer

	// 2. Gunakan c.Context().SetBodyStreamWriter untuk streaming response
	// Ini memungkinkan kita menulis ke response secara bertahap.
	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		log.Println("SSE Client terhubung:", c.IP())
		eventID := 0

		// Loop untuk mengirim event secara berkala
		for {
			select {
			// Cek apakah koneksi client masih aktif (opsional tapi bagus)
			case <-c.Context().Done(): // Context bawaan Fiber/Fasthttp
				log.Println("SSE Client terputus (context done):", c.IP())
				return // Hentikan loop jika client disconnect

			// Tunggu interval waktu tertentu
			case <-time.After(2 * time.Second):
				eventID++
				// Format pesan SSE:
				// id: <unique_id>\n
				// event: <event_name>\n (opsional)
				// data: <your_data>\n\n (data bisa multiline jika diawali 'data: ')

				// Kirim event 'message' dengan data waktu saat ini
				fmt.Fprintf(w, "id: %d\n", eventID)
				fmt.Fprintf(w, "event: server-time\n")
				fmt.Fprintf(w, "data: {\"time\": \"%s\"}\n\n", time.Now().Format(time.RFC3339))

				// Kirim event lain (contoh)
				if eventID % 5 == 0 {
					fmt.Fprintf(w, "id: %d-ping\n", eventID)
					fmt.Fprintf(w, "event: ping\n") // Event tanpa data
					fmt.Fprintf(w, "data: \n\n")
				}

				// Flush buffer untuk memastikan data dikirim ke client
				if err := w.Flush(); err != nil {
					// Error flushing biasanya berarti client disconnect
					log.Printf("SSE Error flushing / client disconnect: %v", err)
					// Hentikan goroutine jika tidak bisa flush
					// (c.Context().Done() mungkin belum ter-trigger)
					return
				}
				log.Printf("SSE event %d dikirim ke %s", eventID, c.IP())
			}
		}
	})

	// Penting: Return nil di sini karena response ditulis oleh stream writer
	// Fiber/Fasthttp akan menangani penyelesaian response setelah stream writer selesai.
	return nil
}

func main() {
	app := fiber.New()

	app.Get("/events", sseHandler)

	// Halaman HTML sederhana untuk menguji SSE dari browser
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(`
			<h1>SSE Test</h1>
			<ul id="events"></ul>
			<script>
				const eventsList = document.getElementById('events');
				const evtSource = new EventSource("/events"); // Hubungkan ke endpoint SSE

				// Handler untuk event bernama 'server-time'
				evtSource.addEventListener("server-time", function(event) {
					console.log("Received server-time event:", event.data);
					const data = JSON.parse(event.data);
					const newItem = document.createElement("li");
					newItem.textContent = "Server Time: " + data.time + " (ID: " + event.lastEventId + ")";
					eventsList.appendChild(newItem);
				});

				 // Handler untuk event bernama 'ping'
				evtSource.addEventListener("ping", function(event) {
					console.log("Received ping event (ID: " + event.lastEventId + ")");
					const newItem = document.createElement("li");
					newItem.textContent = "PING! (ID: " + event.lastEventId + ")";
					eventsList.appendChild(newItem);
				});

				// Handler untuk event 'message' default (jika 'event:' tidak diset)
				evtSource.onmessage = function(event) {
					console.log("Received generic message:", event.data);
					 const newItem = document.createElement("li");
					newItem.textContent = "Generic Message: " + event.data;
					eventsList.appendChild(newItem);
				};

				// Handler untuk error koneksi
				evtSource.onerror = function(err) {
					console.error("EventSource failed:", err);
					const newItem = document.createElement("li");
					newItem.textContent = "Connection error!";
					newItem.style.color = "red";
					eventsList.appendChild(newItem);
					// Browser biasanya akan mencoba reconnect otomatis
					// evtSource.close(); // Tutup manual jika perlu
				};

				console.log("Connecting to SSE stream...");
			</script>
		`)
	})

	log.Fatal(app.Listen(":3000"))
}
```

**Penjelasan:**
*   **Headers:** `Content-Type: text/event-stream`, `Cache-Control: no-cache`, `Connection: keep-alive` sangat penting.
*   **`c.Context().SetBodyStreamWriter(func(w *bufio.Writer))`**: Ini adalah kunci streaming di Fiber/Fasthttp. Anda mendapatkan `bufio.Writer` yang terhubung langsung ke response stream.
*   **Format Pesan:** Setiap pesan diakhiri dengan dua newline (`\n\n`). Pesan dapat memiliki field `id`, `event`, dan `data`.
*   **`w.Flush()`**: Penting untuk dipanggil setelah menulis setiap event agar data dikirim ke client dan tidak tertahan di buffer.
*   **Penanganan Disconnect:** Memeriksa `c.Context().Done()` atau error saat `w.Flush()` adalah cara mendeteksi jika client menutup koneksi.
*   **Client-Side (JavaScript):** Menggunakan objek `EventSource` bawaan browser untuk terhubung dan mendengarkan event. `addEventListener` digunakan untuk event bernama, `onmessage` untuk event default, `onerror` untuk error.

SSE adalah alternatif yang bagus untuk WebSocket jika Anda hanya perlu push data dari server ke client.

### Integrasi Database

Fiber adalah framework web, bukan ORM (Object-Relational Mapper) atau library database. Integrasi database dilakukan menggunakan library Go standar atau pihak ketiga, sama seperti pada aplikasi Go lainnya.

Langkah-langkah umum:

1.  **Pilih Library Database:**
	*   **`database/sql` (Bawaan Go):** Interface standar untuk database SQL. Anda perlu driver spesifik untuk database Anda (misalnya, `github.com/lib/pq` untuk PostgreSQL, `github.com/go-sql-driver/mysql` untuk MySQL, `github.com/mattn/go-sqlite3` untuk SQLite). Memberikan kontrol penuh tapi memerlukan penulisan query SQL manual.
	*   **`sqlx` (`github.com/jmoiron/sqlx`):** Ekstensi untuk `database/sql` yang memudahkan scan hasil query ke struct dan bekerja dengan slice/map. Masih perlu menulis SQL.
	*   **GORM (`gorm.io/gorm`):** ORM paling populer untuk Go. Menyediakan API high-level untuk query, migrasi, relasi, hook, dll. Mengurangi penulisan SQL manual tapi memiliki learning curve sendiri.
	*   **Driver NoSQL:** Jika menggunakan MongoDB, Cassandra, dll., gunakan driver resmi atau komunitas untuk database tersebut (misalnya `go.mongodb.org/mongo-driver`).

2.  **Inisialisasi Koneksi (Connection Pool):**
	*   Jangan membuka koneksi baru untuk setiap request. Gunakan *connection pool* yang dikelola oleh library database.
	*   Inisialisasi pool saat aplikasi dimulai (misalnya, di fungsi `main` atau fungsi init terpisah).
	*   Simpan objek pool (misalnya `*sql.DB`, `*sqlx.DB`, `*gorm.DB`) agar dapat diakses oleh handler.

3.  **Akses Database di Handler:**
	*   Cara paling sederhana adalah membuat objek pool database menjadi variabel global (tidak ideal untuk testing dan skala besar).
	*   Cara yang lebih baik:
		*   **Dependency Injection:** Masukkan objek pool ke dalam struct handler Anda.
		*   **Middleware + Locals:** Buat middleware yang menambahkan objek pool (atau transaksi) ke `c.Locals()` untuk digunakan oleh handler.
		*   **Receiver Method:** Definisikan handler sebagai method pada struct yang memiliki akses ke pool database.

**Contoh dengan `database/sql` dan PostgreSQL (Driver: `lib/pq`):**

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq" // Import driver (blank import)
)

// Struct untuk data produk
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

// --- Setup Database ---

// Global variable (cara sederhana, pertimbangkan DI untuk proyek besar)
var db *sql.DB

// Fungsi untuk inisialisasi koneksi DB
func initDatabase() (*sql.DB, error) {
	// Ambil DSN (Data Source Name) dari environment variable
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback DSN untuk development jika env var tidak diset
		dsn = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
		log.Println("DATABASE_URL tidak diset, menggunakan DSN default:", dsn)
	}

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi db: %w", err)
	}

	// Konfigurasi connection pool
	db.SetMaxOpenConns(25) // Jumlah maksimum koneksi terbuka
	db.SetMaxIdleConns(25) // Jumlah maksimum koneksi idle
	db.SetConnMaxLifetime(5 * time.Minute) // Waktu maksimum koneksi bisa digunakan ulang

	// Coba ping ke database untuk memastikan koneksi berhasil
	if err = db.Ping(); err != nil {
		db.Close() // Tutup jika ping gagal
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}

	log.Println("Koneksi database berhasil dibuat!")
	return db, nil
}

// --- Handler ---

// Handler untuk mendapatkan semua produk
func getProductsHandler(c *fiber.Ctx) error {
	rows, err := db.QueryContext(c.Context(), "SELECT id, name, price, created_at FROM products ORDER BY created_at DESC")
	if err != nil {
		log.Printf("Error query produk: %v", err)
		return fiber.ErrInternalServerError // Kirim 500
	}
	defer rows.Close() // Penting: Selalu tutup rows

	products := []Product{} // Slice untuk menampung hasil
	for rows.Next() {
		var p Product
		// Scan hasil query ke field struct Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
			log.Printf("Error scan row produk: %v", err)
			return fiber.ErrInternalServerError
		}
		products = append(products, p)
	}

	// Cek error yang mungkin terjadi selama iterasi rows
	if err = rows.Err(); err != nil {
		log.Printf("Error iterasi rows produk: %v", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(products) // Kirim hasil sebagai JSON
}

// Handler untuk membuat produk baru
func createProductHandler(c *fiber.Ctx) error {
	// 1. Bind & Validate input
	input := new(struct {
		Name  string  `json:"name" validate:"required"`
		Price float64 `json:"price" validate:"required,gt=0"`
	})

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body tidak valid"})
	}

	// (Tambahkan validasi dengan validator di sini jika perlu)

	// 2. Eksekusi query INSERT
	var newID int
	var createdAt time.Time
	err := db.QueryRowContext(c.Context(), // Gunakan context request untuk pembatalan
		"INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, created_at",
		input.Name, input.Price,
	).Scan(&newID, &createdAt) // Scan ID dan created_at yang dikembalikan

	if err != nil {
		log.Printf("Error insert produk: %v", err)
		// Handle error spesifik (misalnya, constraint violation) jika perlu
		return fiber.ErrInternalServerError
	}

	// 3. Buat response
	newProduct := Product{
		ID:        newID,
		Name:      input.Name,
		Price:     input.Price,
		CreatedAt: createdAt,
	}

	return c.Status(fiber.StatusCreated).JSON(newProduct)
}

func main() {
	// Inisialisasi database saat startup
	var err error
	db, err = initDatabase()
	if err != nil {
		log.Fatalf("Gagal menginisialisasi database: %v", err)
	}
	// Pastikan koneksi ditutup saat aplikasi berhenti (misalnya saat shutdown)
	// defer db.Close() // Penutupan lebih baik ditangani di graceful shutdown

	app := fiber.New()
	app.Use(logger.New())

	// Rute produk
	app.Get("/products", getProductsHandler)
	app.Post("/products", createProductHandler)
	// (Tambahkan handler GET /products/:id, PUT /products/:id, DELETE /products/:id)

	// Implementasi Graceful Shutdown (lihat bagian Graceful Shutdown)
	// ...

	log.Fatal(app.Listen(":3000"))
}
```

**Pola Dependency Injection (Lebih Baik):**

```go
// product_handler.go
package handlers

import (
	"database/sql"
	"log"
	// ... import lainnya
)

// Struct Handler dengan dependency DB
type ProductHandler struct {
	DB *sql.DB
}

// Buat constructor untuk handler
func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{DB: db}
}

// Ubah handler menjadi method pada struct ProductHandler
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	rows, err := h.DB.QueryContext(c.Context(), "SELECT ... FROM products ...")
	// ... sisa logika sama, gunakan h.DB ...
	return c.JSON(...)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	input := new(...)
	if err := c.BodyParser(input); err != nil { ... }
	// ... validasi ...

	var newID int
	err := h.DB.QueryRowContext(c.Context(), "INSERT ...", input.Name, input.Price).Scan(&newID)
	// ... sisa logika sama, gunakan h.DB ...
	return c.Status(fiber.StatusCreated).JSON(...)
}


// main.go
package main

import (
	"database/sql"
	"log"
	// ... import fiber, driver pq, handlers package ...
	"myproject/handlers" // Sesuaikan path import
)

func initDatabase() (*sql.DB, error) { ... } // Sama seperti sebelumnya

func main() {
	db, err := initDatabase()
	if err != nil { log.Fatalf(...) }
	// defer db.Close()

	app := fiber.New()
	app.Use(logger.New())

	// Buat instance handler produk dengan menyuntikkan DB
	productHandler := handlers.NewProductHandler(db)

	// Daftarkan rute menggunakan method dari instance handler
	app.Get("/products", productHandler.GetProducts)
	app.Post("/products", productHandler.CreateProduct)
	// ... rute produk lainnya ...

	log.Fatal(app.Listen(":3000"))
}
```
Pola DI ini membuat handler lebih mudah di-test (Anda bisa mock `*sql.DB`) dan lebih terstruktur.

### Autentikasi & Otorisasi (JWT, Sesi)

Mengamankan endpoint API atau halaman web adalah kebutuhan umum. Dua pendekatan populer adalah:

1.  **Sesi (Session-based Authentication):**
	*   **Cara Kerja:**
		1.  Pengguna login dengan username/password.
		2.  Server memverifikasi kredensial.
		3.  Server membuat ID sesi unik, menyimpannya (di memori, Redis, DB) bersama data pengguna (misalnya ID pengguna).
		4.  Server mengirim ID sesi ke client, biasanya disimpan dalam *cookie*.
		5.  Untuk request berikutnya, client mengirim cookie sesi.
		6.  Server mengambil ID sesi dari cookie, mencarinya di penyimpanan sesi, dan mendapatkan data pengguna terkait. Jika sesi valid, akses diberikan.
		7.  Saat logout, sesi di server dihapus.
	*   **Pros:** Konsep relatif sederhana, stateful (mudah menyimpan data sesi), logout di server mudah (hapus sesi).
	*   **Cons:** Stateful (memerlukan penyimpanan sisi server, bisa jadi bottleneck), masalah skalabilitas (jika pakai memori & prefork/multi-server), kurang cocok untuk API stateless murni atau mobile app (cookie tidak selalu ideal).
	*   **Implementasi di Fiber:** Gunakan middleware `session` (`github.com/gofiber/fiber/v2/middleware/session`). Middleware ini mendukung berbagai *storage* (memory, Redis, PostgreSQL, MySQL, dll.).

2.  **Token (JWT - JSON Web Token):**
	*   **Cara Kerja:**
		1.  Pengguna login.
		2.  Server memverifikasi kredensial.
		3.  Server membuat *token* (JWT) yang berisi *payload* (data pengguna seperti ID, role, dll.) dan *signature* (menggunakan secret key yang hanya diketahui server).
		4.  Server mengirim token ke client. Client menyimpannya (localStorage, sessionStorage, header Authorization).
		5.  Untuk request berikutnya, client mengirim token (biasanya di header `Authorization: Bearer <token>`).
		6.  Server menerima token, *memverifikasi signature* menggunakan secret key. Jika signature valid, server *mempercayai* payload di dalam token (tidak perlu cek ke DB/session store).
		7.  Server menggunakan data dari payload (misalnya user ID) untuk otorisasi.
	*   **Pros:** Stateless (server tidak perlu menyimpan state sesi), cocok untuk API & microservices, skalabel, cocok untuk mobile app, banyak library tersedia.
	*   **Cons:** Token bisa lebih besar dari ID sesi, logout lebih kompleks (token tidak bisa dibatalkan di server secara default, perlu blacklist/short expiry), data payload terlihat (meski tidak bisa diubah tanpa secret key), perlu penanganan refresh token untuk sesi panjang.
	*   **Implementasi di Fiber:** Gunakan middleware JWT (`github.com/gofiber/contrib/jwt`) atau library JWT Go populer (`github.com/golang-jwt/jwt/v5`) dan buat middleware kustom.

**Contoh Middleware JWT Kustom (menggunakan `golang-jwt/jwt/v5`):**

```go
package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Kunci rahasia untuk menandatangani token (simpan di env var!)
var jwtSecret = []byte("kunci-rahasia-super-aman-jangan-hardcode") // GANTI INI!

// Struct untuk claims (payload) JWT kustom
type MyCustomClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims // Embed standard claims (Issuer, Subject, Audience, ExpiresAt, NotBefore, IssuedAt, JWTID)
}

// Fungsi untuk generate token JWT
func GenerateJWT(userID int, role string) (string, error) {
	// Set custom claims
	claims := MyCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			// Set expiration time (misal 1 jam)
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app", // Nama aplikasi Anda
			Subject:   "user-auth",
		},
	}

	// Buat token baru dengan claims dan metode signing HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Middleware untuk memproteksi rute
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Header Authorization tidak ada",
			})
		}

		// Cek format header "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Format header Authorization salah (harus 'Bearer <token>')",
			})
		}
		tokenString := parts[1]

		// Parse dan validasi token
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Pastikan metode signing adalah HMAC (HS256) sesuai yang kita gunakan
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Metode signing token tidak valid")
			}
			// Kembalikan secret key kita
			return jwtSecret, nil
		})

		if err != nil {
			log.Printf("Error parsing/validasi token: %v", err)
			// Cek jenis error (misalnya, token expired)
			if errors.Is(err, jwt.ErrTokenExpired) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token telah kedaluwarsa"})
			}
			// Error lain (signature tidak valid, format salah, dll)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
		}

		// Jika token valid, ambil claims
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			// Simpan informasi pengguna dari claims ke Locals untuk digunakan handler
			c.Locals("userID", claims.UserID)
			c.Locals("userRole", claims.Role)
			c.Locals("jwtClaims", claims) // Simpan semua claims jika perlu
			log.Printf("Middleware JWT: Akses diberikan untuk User ID %d (Role: %s)", claims.UserID, claims.Role)
			return c.Next() // Lanjutkan ke handler
		}

		// Jika claims tidak bisa di-cast atau token tidak valid (seharusnya sudah ditangani err sebelumnya)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid atau claims rusak"})
	}
}

// (Opsional) Middleware untuk otorisasi berdasarkan role
func AuthorizeRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil role dari Locals (diasumsikan middleware Protected sudah berjalan)
		role, ok := c.Locals("userRole").(string)
		if !ok {
			log.Println("Middleware Authorize: Role pengguna tidak ditemukan di Locals")
			 return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Akses ditolak (role tidak diketahui)"})
		}

		// Cek apakah role pengguna ada di daftar role yang diizinkan
		isAllowed := false
		for _, allowed := range allowedRoles {
			if role == allowed {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			log.Printf("Middleware Authorize: Akses ditolak untuk role '%s'. Role yang diizinkan: %v", role, allowedRoles)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Anda tidak memiliki izin yang cukup"}) // 403 Forbidden
		}

		log.Printf("Middleware Authorize: Akses diizinkan untuk role '%s'", role)
		return c.Next() // Lanjutkan jika diizinkan
	}
}


// Di main.go atau auth_routes.go
/*
import (
	"myproject/middleware" // Import package middleware Anda
	// ... import lain
)

func main() {
	app := fiber.New()
	// ... middleware lain (logger, recover) ...

	// --- Rute Autentikasi ---
	auth := app.Group("/auth")
	auth.Post("/login", func(c *fiber.Ctx) error {
		// 1. Ambil username/password dari body
		// 2. Verifikasi ke database
		// 3. Jika valid, generate JWT
		userID := 123 // Contoh ID dari DB
		userRole := "admin" // Contoh role dari DB
		token, err := middleware.GenerateJWT(userID, userRole)
		if err != nil {
			 return fiber.ErrInternalServerError
		}
		// 4. Kirim token ke client
		return c.JSON(fiber.Map{"token": token})
	})

	// --- Rute Terproteksi ---
	api := app.Group("/api")
	api.Use(middleware.Protected()) // Terapkan middleware JWT di sini

	api.Get("/me", func(c *fiber.Ctx) error {
		// Ambil data dari Locals yang disimpan middleware
		userID := c.Locals("userID").(int)
		userRole := c.Locals("userRole").(string)
		return c.JSON(fiber.Map{"user_id": userID, "role": userRole})
	})

	// --- Rute Terproteksi dengan Otorisasi Role ---
	adminApi := api.Group("/admin")
	// Hanya user dengan role 'admin' yang bisa akses endpoint di grup ini
	adminApi.Use(middleware.AuthorizeRole("admin"))

	adminApi.Get("/users", func(c *fiber.Ctx) error {
		// Hanya admin yang bisa sampai sini
		return c.JSON(fiber.Map{"message": "Daftar semua pengguna (khusus admin)"})
	})

	// Endpoint lain yang bisa diakses role 'admin' atau 'editor'
	api.Post("/articles", middleware.AuthorizeRole("admin", "editor"), func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(int)
		return c.JSON(fiber.Map{"message": "Membuat artikel baru (oleh user " + fmt.Sprint(userID) + ")"})
	})


	log.Fatal(app.Listen(":3000"))
}
*/

// Jangan lupa import "errors", "fmt"
```

**Penting untuk JWT:**
*   **Secret Key:** Jaga kerahasiaan secret key Anda! Simpan di environment variable atau sistem manajemen secret, jangan di-hardcode.
*   **Expiration:** Selalu set waktu kedaluwarsa (`ExpiresAt`) pada token untuk membatasi waktu validitasnya.
*   **HTTPS:** Selalu gunakan HTTPS untuk mengirim token agar tidak mudah dicuri.
*   **Payload:** Jangan simpan data sensitif di payload JWT karena mudah dibaca (meski tidak bisa diubah). Cukup simpan ID pengguna, role, atau informasi non-sensitif lainnya.
*   **Refresh Token:** Untuk sesi yang lebih panjang, implementasikan mekanisme *refresh token*. Token utama (access token) memiliki masa berlaku singkat (menit/jam), sedangkan refresh token memiliki masa berlaku lebih panjang (hari/minggu) dan disimpan lebih aman. Client menggunakan refresh token untuk mendapatkan access token baru tanpa perlu login ulang.
*   **Blacklist (Opsional):** Jika Anda perlu cara untuk membatalkan token sebelum kedaluwarsa (misalnya saat pengguna ganti password atau logout paksa), Anda perlu implementasi blacklist sisi server (misalnya di Redis) untuk menyimpan ID token yang sudah tidak valid. Middleware JWT perlu memeriksa blacklist ini.

Pilih strategi autentikasi (sesi atau token) yang paling sesuai dengan kebutuhan aplikasi Anda.

### Pengujian (Testing)

Menulis tes otomatis sangat penting untuk memastikan aplikasi Anda berfungsi seperti yang diharapkan dan untuk mencegah regresi saat Anda melakukan perubahan. Go memiliki dukungan pengujian bawaan yang kuat, dan ini bekerja dengan baik untuk menguji aplikasi Fiber.

Pendekatan umumnya adalah menggunakan package `net/http/httptest` untuk membuat request HTTP tiruan dan meneruskannya langsung ke handler Fiber Anda tanpa perlu menjalankan server HTTP sungguhan.

**Struktur Tes:**

*   File tes harus diakhiri dengan `_test.go` (misalnya, `main_test.go`, `handlers/product_handler_test.go`).
*   Fungsi tes harus diawali dengan `Test` dan menerima `t *testing.T` (misalnya, `TestGetProductsHandler`).
*   Gunakan package `testing` Go dan assertion library (opsional tapi membantu) seperti `testify/assert` atau `testify/require`.

**Contoh Tes untuk Handler Sederhana:**

Misalkan kita punya `main.go` seperti ini:
```go
// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/echo", func(c *fiber.Ctx) error {
		body := c.Body() // Ambil body mentah
		// Kirim kembali body yang sama dengan content type asli
		return c.Status(fiber.StatusOK).Send(body)
	})

	app.Get("/users/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.JSON(fiber.Map{"user": name})
	})
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
```

Sekarang buat file `main_test.go`:

```go
// main_test.go
package main // Harus package yang sama dengan kode yang diuji

import (
	"io"
	"net/http"          // Import net/http untuk konstanta status code dll.
	"net/http/httptest" // Import httptest
	"strings"
	"testing" // Import testing

	"github.com/gofiber/fiber/v2" // Import fiber
	"github.com/stretchr/testify/assert" // Import testify/assert (opsional)
	"github.com/stretchr/testify/require" // Import testify/require (opsional)
)

// Fungsi helper untuk membuat request dan mendapatkan response
func performRequest(app *fiber.App, method, path string, body io.Reader, headers map[string]string) *http.Response {
	// Buat request HTTP tiruan
	req := httptest.NewRequest(method, path, body)

	// Tambahkan header jika ada
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Jalankan request melalui handler Fiber
	// app.Test() adalah metode Fiber khusus untuk testing
	resp, err := app.Test(req, -1) // -1 untuk tidak ada timeout
	if err != nil {
		// Fail test jika app.Test() gagal (jarang terjadi)
		panic(err) // atau t.Fatalf("app.Test failed: %v", err) jika t tersedia
	}

	return resp // Kembalikan response HTTP standar
}


// Tes untuk rute GET /
func TestGetRoot(t *testing.T) {
	// 1. Setup Aplikasi Fiber (hanya untuk tes ini)
	app := fiber.New()
	setupRoutes(app) // Daftarkan rute

	// 2. Buat Request
	// req := httptest.NewRequest("GET", "/", nil)
	// resp, _ := app.Test(req)
	resp := performRequest(app, "GET", "/", nil, nil)
	defer resp.Body.Close() // Penting: Selalu tutup body response

	// 3. Assertions (menggunakan testify/assert)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code harus 200 OK")

	// Baca body response
	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Harus bisa membaca body response") // require: fail test jika error
	assert.Equal(t, "Hello, World!", string(bodyBytes), "Body response harus 'Hello, World!'")
}

// Tes untuk rute POST /echo
func TestPostEcho(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	// Data yang akan dikirim di body
	requestBody := `{"message": "ping"}`
	headers := map[string]string{"Content-Type": "application/json"} // Set header

	resp := performRequest(app, "POST", "/echo", strings.NewReader(requestBody), headers)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code harus 200 OK")

	// Baca body response
	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	// Bandingkan body response dengan body request
	assert.JSONEq(t, requestBody, string(bodyBytes), "Body response harus sama dengan body request (JSON)")
	// Cek Content-Type response (meskipun handler tidak set eksplisit, Fiber/Fasthttp mungkin mendeteksinya)
	// assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")
}

// Tes untuk rute GET /users/:name
func TestGetUserByName(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	userName := "alice"
	path := "/users/" + userName

	resp := performRequest(app, "GET", path, nil, nil)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code harus 200 OK")
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json", "Content-Type harus application/json")

	// Baca dan unmarshal body JSON response
	var result map[string]string
	err := json.NewDecoder(resp.Body).Decode(&result) // Import "encoding/json"
	require.NoError(t, err, "Harus bisa decode JSON response")

	// Cek isi JSON
	assert.Equal(t, userName, result["user"], "Nilai 'user' di JSON harus sama dengan parameter path")
}

// Tes untuk rute yang tidak ada (404 Not Found)
func TestNotFound(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	resp := performRequest(app, "GET", "/path/tidak/ada", nil, nil)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Status code harus 404 Not Found")
}

// Jangan lupa import "encoding/json"
```

**Menjalankan Tes:**

Buka terminal di direktori proyek Anda dan jalankan:

```bash
go test ./... -v
```

*   `go test`: Perintah dasar untuk menjalankan tes.
*   `./...`: Menjalankan tes di direktori saat ini dan semua subdirektori.
*   `-v`: Mode verbose, menampilkan nama tes yang dijalankan dan hasilnya (PASS/FAIL).

**Testing dengan Database:**

Menguji handler yang berinteraksi dengan database memerlukan strategi tambahan:

1.  **Database Tes Terpisah:** Cara terbaik adalah menggunakan database tes yang terpisah (misalnya, database Docker sementara atau database khusus tes). Tes Anda akan terhubung ke database ini, melakukan setup data yang diperlukan, menjalankan handler, dan kemudian membersihkan data (atau mengembalikan transaksi). Ini memastikan tes berjalan dalam kondisi terisolasi dan tidak mengganggu database development/produksi.
2.  **Mocking Database:** Menggunakan library mocking (seperti `sqlmock` untuk `database/sql` atau fitur mock GORM) untuk meniru perilaku database tanpa benar-benar terhubung ke database. Ini lebih cepat tetapi mungkin tidak menangkap semua perilaku database nyata. Cocok untuk unit test murni pada logika handler.

**Contoh (Konsep) Testing Handler dengan DB:**

```go
// product_handler_test.go
package handlers_test // Gunakan package _test untuk black-box testing

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"myproject/handlers" // Import package handler asli
	"myproject/models" // Import model (misal, struct Product)
	_ "github.com/lib/pq" // Import driver
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Fungsi helper untuk setup DB tes (contoh sederhana, idealnya pakai Docker/migrasi)
func setupTestDB(t *testing.T) *sql.DB {
	dsn := "postgres://user:password@localhost:5433/testdb?sslmode=disable" // DB Tes di port berbeda
	db, err := sql.Open("postgres", dsn)
	require.NoError(t, err, "Harus bisa konek ke DB tes")
	// Bersihkan tabel sebelum tes (atau gunakan transaksi)
	_, err = db.Exec("DELETE FROM products")
	require.NoError(t, err)
	return db
}

func TestCreateProductHandler(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Setup Fiber App dan Handler dengan DB tes
	app := fiber.New()
	productHandler := handlers.NewProductHandler(db) // Gunakan DI
	app.Post("/products", productHandler.CreateProduct)

	// Data input
	inputData := map[string]interface{}{
		"name": "Test Product",
		"price": 99.99,
	}
	body, _ := json.Marshal(inputData)

	// Buat request
	req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Jalankan request
	resp, err := app.Test(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Harus 201 Created")

	// Cek body response
	var createdProduct models.Product
	err = json.NewDecoder(resp.Body).Decode(&createdProduct)
	require.NoError(t, err)
	assert.Equal(t, inputData["name"], createdProduct.Name)
	assert.Equal(t, inputData["price"], createdProduct.Price)
	assert.NotZero(t, createdProduct.ID, "ID produk harus dihasilkan")
	assert.NotZero(t, createdProduct.CreatedAt, "CreatedAt harus diisi")

	// (Opsional) Verifikasi data di DB tes
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM products WHERE id = $1", createdProduct.ID).Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "Produk harus tersimpan di DB")
}

// ... tes lainnya (GetProducts, GetProductByID, etc.) ...
```

**Tips Testing:**
*   **Fokus pada Input dan Output:** Tes handler harus fokus pada apakah input request yang berbeda menghasilkan output response (status code, headers, body) yang diharapkan.
*   **Isolasi:** Usahakan menguji setiap handler secara terpisah. Jika handler memanggil service lain, pertimbangkan untuk mock service tersebut dalam unit test handler.
*   **Gunakan Test Suites:** Untuk tes yang lebih terstruktur, gunakan fitur test suite Go atau `testify/suite` untuk setup/teardown yang lebih baik.
*   **Cakupan Kode (Code Coverage):** Gunakan `go test ./... -cover` untuk melihat persentase kode Anda yang dicakup oleh tes. Targetkan cakupan yang tinggi, tetapi fokus pada pengujian jalur kritis dan logika kompleks.

### Deployment

Setelah aplikasi Anda siap, langkah selanjutnya adalah men-deploynya ke server produksi. Ada banyak cara untuk men-deploy aplikasi Go/Fiber:

1.  **Compile & Copy Binary:**
	*   **Cara Kerja:** Compile aplikasi Go Anda menjadi *single static binary* di mesin development Anda (atau CI/CD server), lalu salin binary tersebut ke server produksi dan jalankan.
	*   **Compile:**
		```bash
		# Untuk Linux (paling umum untuk server)
		GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o myapp main.go

		# -ldflags="-s -w": Menghapus simbol debug dan DWARF, menghasilkan binary lebih kecil.
		# -o myapp: Nama output binary.
		```
	*   **Jalankan di Server:** Salin `myapp` ke server, lalu jalankan: `./myapp`
	*   **Pros:** Sangat sederhana, dependensi minimal di server (hanya binary), cepat dimulai.
	*   **Cons:** Perlu setup server manual (OS, firewall, reverse proxy), perlu cara untuk menjalankan binary sebagai service (agar otomatis restart jika crash atau server reboot).

2.  **Menggunakan Supervisor (Systemd, SupervisorD):**
	*   **Cara Kerja:** Gunakan manajer proses seperti `systemd` (umum di Linux modern) atau `supervisord` untuk mengelola binary aplikasi Anda sebagai *service*. Manajer ini akan menangani start, stop, restart otomatis, logging, dll.
	*   **Contoh Unit File Systemd (`/etc/systemd/system/myapp.service`):**
		```ini
		[Unit]
		Description=Aplikasi Fiber Saya
		After=network.target # Jalankan setelah jaringan siap

		[Service]
		User=namauseraplikasi # Jalankan sebagai user non-root
		Group=namauseraplikasi
		WorkingDirectory=/path/ke/direktori/aplikasi # Direktori kerja aplikasi
		EnvironmentFile=/path/ke/direktori/aplikasi/.env # Muat env vars dari file (opsional)
		ExecStart=/path/ke/binary/myapp # Path ke binary Anda
		Restart=on-failure # Restart jika gagal
		RestartSec=5s # Tunggu 5 detik sebelum restart
		StandardOutput=journal # Arahkan stdout ke journald
		StandardError=journal # Arahkan stderr ke journald
		LimitNOFILE=65536 # Tingkatkan batas file descriptor (penting untuk banyak koneksi)

		[Install]
		WantedBy=multi-user.target # Aktifkan saat boot multi-user
		```
	*   **Perintah Systemd:**
		*   `sudo systemctl daemon-reload`: Muat ulang konfigurasi setelah membuat/mengubah file service.
		*   `sudo systemctl enable myapp`: Aktifkan service agar start saat boot.
		*   `sudo systemctl start myapp`: Jalankan service sekarang.
		*   `sudo systemctl status myapp`: Lihat status service.
		*   `sudo systemctl stop myapp`: Hentikan service.
		*   `sudo journalctl -u myapp -f`: Lihat log service (ikuti log baru).
	*   **Pros:** Manajemen proses yang andal, restart otomatis, logging terpusat.
	*   **Cons:** Masih perlu setup server manual.

3.  **Docker:**
	*   **Cara Kerja:** Kemas aplikasi Anda beserta semua dependensinya (termasuk Go runtime jika perlu, meskipun binary statis lebih umum) ke dalam *Docker image*. Jalankan image ini sebagai *container* di server mana pun yang memiliki Docker terinstal.
	*   **Contoh `Dockerfile` (Multi-stage build untuk binary kecil):**
		```dockerfile
		# --- Tahap 1: Build ---
		FROM golang:1.21-alpine AS builder

		WORKDIR /app

		# Copy go mod dan sum, lalu unduh dependensi (memanfaatkan cache Docker)
		COPY go.mod go.sum ./
		RUN go mod download

		# Copy sisa source code
		COPY . .

		# Build binary statis
		# CGO_ENABLED=0 penting untuk static build di Alpine
		# -ldflags="-s -w" untuk ukuran lebih kecil
		RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /myapp main.go

		# --- Tahap 2: Run ---
		# Gunakan image dasar yang sangat kecil (scratch atau alpine)
		FROM alpine:latest
		# FROM scratch # Image paling kecil, tapi tanpa shell/tools

		WORKDIR /app

		# (Opsional) Install sertifikat CA jika aplikasi perlu koneksi HTTPS keluar
		RUN apk --no-cache add ca-certificates

		# Copy binary yang sudah di-build dari tahap builder
		COPY --from=builder /myapp /myapp

		# (Opsional) Copy file statis atau template jika tidak di-embed
		# COPY --from=builder /app/public ./public
		# COPY --from=builder /app/views ./views

		# Expose port yang didengarkan aplikasi Fiber
		EXPOSE 3000

		# Set environment variable default (bisa di-override saat run)
		ENV APP_ENV=production
		ENV LISTEN_ADDR=:3000
		# ENV DATABASE_URL=...

		# Perintah untuk menjalankan aplikasi saat container start
		# Entrypoint memastikan binary adalah perintah utama
		ENTRYPOINT ["/myapp"]
		# Cmd bisa ditambahkan untuk argumen default ke entrypoint
		# CMD ["--some-flag"]
		```
	*   **Build & Run:**
		*   `docker build -t myapp-image .`
		*   `docker run -d -p 8080:3000 --name myapp-container -e DATABASE_URL="xxx" myapp-image`
			*   `-d`: Run detached (background).
			*   `-p 8080:3000`: Map port 8080 host ke port 3000 container.
			*   `--name`: Beri nama container.
			*   `-e`: Set environment variable.
	*   **Pros:** Lingkungan konsisten (dev, staging, prod), dependensi terisolasi, mudah diskalakan, portabel, banyak tooling (Docker Compose, Kubernetes).
	*   **Cons:** Perlu belajar Docker, image bisa jadi besar jika tidak dioptimalkan, overhead container (kecil).

4.  **Kubernetes (K8s):**
	*   **Cara Kerja:** Platform orkestrasi container. Anda mendefinisikan state aplikasi Anda (deployment, service, ingress, configmap, secret) dalam file YAML, dan Kubernetes akan menangani deployment, scaling, load balancing, self-healing, dll., di cluster server.
	*   **Pros:** Sangat skalabel, high availability, manajemen kompleksitas untuk aplikasi besar/microservices.
	*   **Cons:** Sangat kompleks, learning curve curam, memerlukan infrastruktur cluster.

5.  **Platform as a Service (PaaS):**
	*   **Cara Kerja:** Platform seperti Heroku, Google App Engine, AWS Elastic Beanstalk, Render, Fly.io. Anda cukup push kode Anda (atau Docker image), dan platform menangani infrastruktur, deployment, scaling, logging.
	*   **Pros:** Sangat mudah digunakan, fokus pada kode bukan infrastruktur, scaling otomatis (tergantung platform).
	*   **Cons:** Kurang fleksibel dibanding solusi self-hosted/IaaS, bisa lebih mahal pada skala besar, vendor lock-in.

6.  **Serverless (Functions as a Service - FaaS):**
	*   **Cara Kerja:** Platform seperti AWS Lambda, Google Cloud Functions, Azure Functions. Anda deploy kode fungsi handler individual. Platform otomatis menjalankan fungsi saat ada request dan men-skalakannya secara otomatis (bahkan hingga nol saat tidak ada traffic).
	*   **Pros:** Bayar per penggunaan, scaling otomatis ekstrem, tidak perlu kelola server.
	*   **Cons:** Cold starts (latensi pada request pertama setelah idle), batasan runtime/durasi, state sulit dikelola, kurang cocok untuk aplikasi stateful atau koneksi long-running (WebSocket/SSE perlu solusi lain). Cocok untuk API stateless atau task kecil.

**Rekomendasi:**

*   **Proyek Kecil/Personal:** Compile & Copy + Systemd/SupervisorD, atau PaaS (Render, Fly.io).
*   **Aplikasi Menengah/Startup:** Docker + Docker Compose (untuk multi-container), atau PaaS.
*   **Aplikasi Besar/Microservices:** Docker + Kubernetes, atau PaaS yang skalabel.
*   **API Sangat Spesifik/Event-Driven:** Serverless (FaaS).

**Reverse Proxy (Nginx, Caddy, Traefik):**

Hampir selalu **direkomendasikan** untuk menjalankan aplikasi Fiber Anda di belakang *reverse proxy* di produksi.

*   **Tugas Reverse Proxy:**
	*   **Terminasi SSL/TLS:** Menangani enkripsi HTTPS, aplikasi Anda bisa berjalan di HTTP biasa di belakang proxy.
	*   **Load Balancing:** Mendistribusikan traffic jika Anda menjalankan beberapa instance aplikasi.
	*   **Caching:** Menyimpan response statis atau dinamis.
	*   **Kompresi:** Mengompresi response (meskipun Fiber bisa melakukannya).
	*   **Rate Limiting/Security:** Menambahkan lapisan keamanan tambahan.
	*   **Menyajikan File Statis:** Bisa lebih efisien daripada Fiber untuk file statis bervolume tinggi.
	*   **Routing Berbasis Host/Path:** Mengarahkan request ke aplikasi yang berbeda di server yang sama.

**Contoh Konfigurasi Nginx Sederhana:**

```nginx
# /etc/nginx/sites-available/myapp.conf

server {
	listen 80; # Dengar di port 80 (HTTP)
	server_name yourdomain.com www.yourdomain.com; # Domain Anda

	# Redirect HTTP ke HTTPS (jika menggunakan SSL)
	# location / {
	#     return 301 https://$host$request_uri;
	# }

	# Jika menggunakan SSL (direkomendasikan), atur di server block lain
	# listen 443 ssl http2;
	# server_name yourdomain.com www.yourdomain.com;
	# ssl_certificate /path/to/fullchain.pem;
	# ssl_certificate_key /path/to/privkey.pem;
	# include /etc/letsencrypt/options-ssl-nginx.conf; # Pengaturan SSL tambahan
	# ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

	location / {
		proxy_pass http://localhost:3000; # Teruskan request ke aplikasi Fiber di port 3000
		proxy_set_header Host $host; # Teruskan header Host asli
		proxy_set_header X-Real-IP $remote_addr; # Teruskan IP asli client
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $scheme; # Teruskan protokol (http/https)

		# Pengaturan tambahan untuk WebSocket jika perlu
		# proxy_http_version 1.1;
		# proxy_set_header Upgrade $http_upgrade;
		# proxy_set_header Connection "upgrade";
		# proxy_read_timeout 86400; # Timeout panjang untuk koneksi WebSocket
	}

	# (Opsional) Biarkan Nginx menyajikan file statis secara langsung
	# location /static/ {
	#     alias /path/ke/direktori/aplikasi/public/;
	#     expires 7d; # Cache di browser selama 7 hari
	#     access_log off; # Matikan log akses untuk file statis
	# }
}

# Jangan lupa buat symlink:
# sudo ln -s /etc/nginx/sites-available/myapp.conf /etc/nginx/sites-enabled/
# sudo systemctl restart nginx
```

### Kinerja & Optimasi

Fiber sudah sangat cepat secara default, tetapi ada beberapa hal yang bisa Anda lakukan untuk mengoptimalkan kinerja lebih lanjut:

*   **Prefork:** Seperti dibahas sebelumnya, `Prefork: true` bisa membantu memanfaatkan semua core CPU untuk aplikasi stateless di Linux/BSD. Lakukan benchmarking untuk kasus Anda.
*   **Hindari Alokasi Berlebih:** Go cepat, tetapi garbage collector (GC) bisa menjadi bottleneck. Hindari alokasi memori yang tidak perlu di dalam hot path (handler yang sering dipanggil). Gunakan `sync.Pool` untuk objek yang sering dibuat ulang. Fasthttp dan Fiber sudah banyak menggunakan pool internal.
*   **Gunakan `fiber.Map` vs `map[string]interface{}`:** Untuk response JSON sederhana, `fiber.Map` sedikit lebih efisien karena merupakan alias dari `map[string]interface{}` tapi mungkin ada optimasi internal terkait pool.
*   **Binding Cepat:** `BodyParser` cukup cepat, tetapi jika Anda tahu pasti format inputnya (misal selalu JSON), menggunakan `c.App().JSONDecoder(body, &out)` (setelah `c.Body()`) mungkin sedikit lebih cepat karena tidak perlu deteksi tipe. Ukur perbedaannya.
*   **Middleware Selektif:** Hanya gunakan middleware yang benar-benar Anda butuhkan. Terapkan middleware pada lingkup yang paling sempit (rute/grup) jika tidak diperlukan secara global.
*   **Query Database Efisien:** Bottleneck paling umum seringkali ada di database.
	*   Gunakan indeks (index) yang tepat pada kolom yang sering di-query.
	*   Hindari query N+1 (ambil data terkait dalam satu atau dua query, bukan banyak query kecil).
	*   Gunakan `SELECT` hanya kolom yang Anda butuhkan, bukan `SELECT *`.
	*   Optimalkan query yang lambat (gunakan `EXPLAIN ANALYZE` di DB Anda).
	*   Gunakan connection pool yang dikonfigurasi dengan baik.
*   **Caching:** Implementasikan caching untuk data yang jarang berubah atau mahal untuk dihasilkan.
	*   **Cache HTTP:** Gunakan header `Cache-Control`, `ETag`, `Last-Modified` agar browser atau CDN bisa cache response. Middleware `etag` Fiber bisa membantu.
	*   **Cache Sisi Server:** Simpan hasil query DB atau komputasi mahal di memori (misalnya pakai `ristretto`, `go-cache`) atau di store eksternal (Redis, Memcached). Ada middleware cache untuk Fiber (bawaan atau pihak ketiga) yang bisa otomatis cache response berdasarkan URL/header.
*   **Kompresi:** Aktifkan kompresi Gzip/Brotli (misalnya dengan middleware `compress` atau di reverse proxy) untuk mengurangi ukuran response dan mempercepat transfer.
*   **Profil Aplikasi (Profiling):** Gunakan tool profiling Go (`pprof`) untuk mengidentifikasi bottleneck CPU dan memori secara akurat. Fiber memudahkan integrasi `pprof` melalui middleware `pprof`. Analisis hasilnya dengan `go tool pprof`.
*   **Build Flags:** Gunakan `-ldflags="-s -w"` saat build untuk produksi (binary lebih kecil).
*   **Versi Go Terbaru:** Selalu gunakan versi Go stabil terbaru, karena seringkali membawa peningkatan kinerja.

**Penting:** Jangan melakukan optimasi prematur. Tulis kode yang bersih dan benar terlebih dahulu, lalu gunakan profiling untuk menemukan bottleneck *nyata* sebelum mengoptimalkan.

### Struktur Proyek

Mengatur kode proyek dengan baik menjadi penting seiring bertambahnya ukuran aplikasi. Tidak ada satu cara "benar", tetapi berikut beberapa pola umum:

**1. Struktur Flat (Sederhana):**

Cocok untuk proyek sangat kecil. Semua file di root.

```
proyek-fiber/
├── go.mod
├── go.sum
├── main.go         # Setup Fiber, DB, Middleware global, Rute
├── handlers.go     # Semua fungsi handler
├── models.go       # Definisi struct data (User, Product)
└── main_test.go
```

**2. Berdasarkan Fitur/Domain:**

Kelompokkan kode berdasarkan fitur atau domain bisnis.

```
proyek-fiber/
├── go.mod
├── go.sum
├── main.go             # Setup Fiber, DB, Middleware global, Panggil setup rute
├── config/             # Fungsi/struct untuk load konfigurasi
│   └── config.go
├── internal/           # Package internal (tidak bisa diimpor dari luar proyek)
│   ├── database/       # Kode koneksi & inisialisasi DB
│   │   └── database.go
│   ├── middleware/     # Middleware kustom
│   │   └── auth.go
│   ├── models/         # Struct data/domain
│   │   ├── user.go
│   │   └── product.go
│   ├── auth/           # Fitur Autentikasi
│   │   ├── handler.go  # Handler terkait auth (login, register)
│   │   ├── service.go  # Logika bisnis auth (opsional)
│   │   └── routes.go   # Setup rute /auth/*
│   ├── product/        # Fitur Produk
│   │   ├── handler.go  # Handler CRUD produk
│   │   ├── repository.go # Logika akses DB produk (interface + implementasi)
│   │   ├── service.go  # Logika bisnis produk (opsional)
│   │   └── routes.go   # Setup rute /products/*
│   └── ...             # Fitur/domain lainnya
├── cmd/                # Titik masuk aplikasi (main package)
│   └── api/
│       └── main.go     # Pindahkan main.go ke sini (opsional tapi umum)
├── web/                # File frontend (jika ada)
│   ├── templates/
│   └── static/
└── Dockerfile
```
*   **`internal/`**: Package di sini hanya bisa diimpor oleh kode lain di dalam `proyek-fiber`. Ini bagus untuk menyembunyikan detail implementasi.
*   **Pemisahan Tanggung Jawab:**
	*   `handlers` (atau `xxx/handler.go`): Menerima request Fiber, mem-parse input, memanggil service/repository, memformat response. Seharusnya *tidak* berisi logika bisnis atau query DB langsung.
	*   `service` (opsional): Berisi logika bisnis inti yang tidak terikat pada HTTP atau DB. Bisa digunakan ulang.
	*   `repository` (atau `store`, `dao`): Berisi logika akses data (query DB). Biasanya berupa interface dan implementasi konkret (misalnya, `PostgresProductRepository`). Ini memudahkan testing dan penggantian storage.
	*   `routes` (atau `router`): Mendefinisikan rute Fiber dan menghubungkannya ke handler yang sesuai.
*   **Dependency Injection:** Sangat cocok dengan struktur ini. `main.go` akan menginisialisasi DB, membuat instance repository, lalu service, lalu handler (menyuntikkan dependensi ke bawah), dan terakhir mendaftarkan rute.

**3. Struktur Layered (Berlapis):**

Mirip dengan struktur domain, tetapi pengelompokan utama berdasarkan lapisan arsitektur.

```
proyek-fiber/
├── go.mod
├── go.sum
├── main.go
├── config/
├── pkg/                # Library/utilitas yang bisa digunakan ulang (bisa diimpor dari luar)
│   └── validator/
├── internal/
│   ├── delivery/       # Lapisan presentasi (misal, HTTP)
│   │   ├── http/
│   │   │   ├── middleware/
│   │   │   ├── handler/    # Semua handler HTTP
│   │   │   │   ├── product_handler.go
│   │   │   │   └── user_handler.go
│   │   │   └── router.go   # Setup semua rute HTTP
│   │   └── grpc/       # (Jika ada delivery gRPC)
│   ├── usecase/        # Lapisan logika bisnis (service)
│   │   ├── product_usecase.go
│   │   └── user_usecase.go
│   ├── repository/     # Lapisan akses data
│   │   ├── postgres/   # Implementasi repo untuk Postgres
│   │   │   ├── product_repo.go
│   │   │   └── user_repo.go
│   │   └── redis/      # Implementasi repo untuk Redis (cache, dll)
│   ├── domain/         # Definisi entitas/model inti & interface repo/usecase
│   │   ├── product.go
│   │   └── user.go
├── cmd/api/main.go
└── ...
```

Pilih struktur yang paling masuk akal untuk ukuran dan kompleksitas proyek Anda. Mulai sederhana dan refactor seiring pertumbuhan. Konsistensi adalah kunci.

### Graceful Shutdown

Saat Anda perlu menghentikan atau me-restart server aplikasi (misalnya saat deployment versi baru atau maintenance), penting untuk melakukannya secara *graceful* (anggun). Artinya:

1.  Server berhenti menerima koneksi *baru*.
2.  Server menunggu koneksi yang *sedang aktif* untuk selesai diproses (dalam batas waktu tertentu).
3.  Setelah semua koneksi selesai atau timeout tercapai, server benar-benar berhenti.

Ini mencegah pemutusan koneksi mendadak yang dapat menyebabkan error pada client atau kehilangan data.

Fiber (melalui Fasthttp) menyediakan metode `app.Shutdown()` atau `app.ShutdownWithTimeout()` untuk ini. Anda perlu menangkap sinyal OS (seperti `SIGINT` dari Ctrl+C atau `SIGTERM` dari `kill`/systemd) untuk memicu shutdown.

```go
package main

import (
	"context" // Diperlukan untuk ShutdownWithContext (opsional tapi baik)
	"log"
	"os"
	"os/signal" // Untuk menangkap sinyal OS
	"syscall"   // Untuk konstanta sinyal (SIGINT, SIGTERM)
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// ... import DB, dll ...
)

// func initDatabase() (*sql.DB, error) { ... }
// var db *sql.DB

func main() {
	// --- Setup Aplikasi ---
	app := fiber.New(fiber.Config{
		// ReadTimeout/WriteTimeout/IdleTimeout membantu dalam graceful shutdown
		IdleTimeout: 5 * time.Second,
	})
	app.Use(logger.New())

	// Setup database (contoh)
	// db, err := initDatabase()
	// if err != nil { log.Fatalf("DB init failed: %v", err) }
	// defer func() {
	//     log.Println("Menutup koneksi database...")
	//     if err := db.Close(); err != nil {
	//         log.Printf("Error menutup DB: %v", err)
	//     }
	// }() // Defer penutupan DB

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("Memproses request ke / ...")
		time.Sleep(3 * time.Second) // Simulasi kerja lama
		log.Println("Selesai memproses request ke /")
		return c.SendString("Request diproses!")
	})

	// --- Start Server di Goroutine ---
	// Jalankan app.Listen di goroutine agar tidak memblokir main goroutine
	// sehingga kita bisa menunggu sinyal shutdown.
	go func() {
		listenAddr := ":3000" // Ambil dari config/env var
		log.Printf("Server mulai mendengarkan di %s", listenAddr)
		if err := app.Listen(listenAddr); err != nil {
			// Error selain http.ErrServerClosed adalah fatal
			if !errors.Is(err, http.ErrServerClosed) { // Perlu import "errors" dan "net/http"
				 log.Fatalf("Gagal menjalankan server: %v", err)
			}
			 log.Println("Server berhenti mendengarkan.")
		}
	}()

	// --- Menunggu Sinyal Shutdown ---
	quit := make(chan os.Signal, 1) // Channel untuk menerima sinyal
	// Tangkap sinyal SIGINT (Ctrl+C) dan SIGTERM (kill default, systemd stop)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Blokir sampai sinyal diterima
	receivedSignal := <-quit
	log.Printf("Menerima sinyal shutdown: %s. Memulai graceful shutdown...", receivedSignal)

	// --- Proses Graceful Shutdown ---
	// Beri batas waktu untuk shutdown (misalnya 10 detik)
	// shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// Panggil app.Shutdown() atau app.ShutdownWithTimeout()
	// err := app.ShutdownWithContext(shutdownCtx) // Fiber v2.45.0+

	// Atau tanpa context timeout (Fiber akan menunggu koneksi idle selesai, dibatasi IdleTimeout)
	err := app.Shutdown()

	if err != nil {
		log.Fatalf("Error saat shutdown server: %v", err)
	}

	log.Println("Graceful shutdown selesai. Server keluar.")
	// Di sini, koneksi DB yang di-defer penutupannya juga akan dijalankan.
}
```

**Penjelasan:**
1.  **Start Server di Goroutine:** `app.Listen()` bersifat blocking. Kita menjalankannya di goroutine agar `main` bisa lanjut menunggu sinyal.
2.  **Setup Signal Channel:** Buat channel `quit` dan gunakan `signal.Notify` untuk mengarahkan sinyal `SIGINT` dan `SIGTERM` ke channel ini.
3.  **Menunggu Sinyal:** `<-quit` akan memblokir eksekusi sampai salah satu sinyal diterima.
4.  **Panggil `app.Shutdown()`:** Setelah sinyal diterima, panggil `app.Shutdown()`. Fiber akan:
	*   Berhenti menerima koneksi baru.
	*   Menunggu koneksi aktif selesai (dibatasi oleh `IdleTimeout` di `fiber.Config`).
	*   Menutup listener.
5.  **Handle Error Shutdown:** Periksa error dari `Shutdown()`.
6.  **Cleanup Lain:** Jika ada resource lain yang perlu ditutup (seperti koneksi DB), lakukan setelah `Shutdown()` selesai atau gunakan `defer` di `main`.

Dengan implementasi ini, saat Anda menghentikan aplikasi (Ctrl+C atau `systemctl stop`), aplikasi akan mencoba menyelesaikan request yang sedang berjalan sebelum benar-benar keluar.

---

## 7. Contoh Aplikasi (CRUD Sederhana) 📝

Mari kita buat contoh API RESTful sederhana untuk mengelola data "Buku" menggunakan Fiber dan menyimpan data di memori (untuk kesederhanaan, dalam aplikasi nyata gunakan database).

```go
package main

import (
	"fmt"
	"log"
	"sync" // Untuk Mutex agar aman dari race condition
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/go-playground/validator/v10"
)

// --- Model ---
type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// --- Input Validation Struct ---
type CreateBookInput struct {
	Title  string `json:"title" validate:"required,min=3"`
	Author string `json:"author" validate:"required,min=3"`
}

type UpdateBookInput struct {
	Title  *string `json:"title" validate:"omitempty,min=3"` // Pointer agar opsional
	Author *string `json:"author" validate:"omitempty,min=3"` // Pointer agar opsional
}

// --- In-Memory Storage ---
var (
	books      = make(map[int]Book) // Map untuk menyimpan buku (ID -> Book)
	nextBookID = 1                   // Counter ID buku
	bookMutex  = &sync.RWMutex{}     // Mutex untuk melindungi akses ke map books
	validate   = validator.New()     // Instance validator
)

// --- Error Response Helper ---
func validationErrorResponse(err error) fiber.Map {
	var errors []fiber.Map
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			 errors = append(errors, fiber.Map{
				 "field": fieldErr.Field(),
				 "tag": fieldErr.Tag(),
				 "message": fmt.Sprintf("Validation failed on '%s' with tag '%s'", fieldErr.Field(), fieldErr.Tag()),
			 })
		}
	}
	 return fiber.Map{"status": "fail", "errors": errors}
}

// --- Handlers ---

// GET /books - Mendapatkan semua buku
func getBooks(c *fiber.Ctx) error {
	bookMutex.RLock() // Lock untuk membaca
	defer bookMutex.RUnlock()

	// Buat slice dari value map
	bookList := make([]Book, 0, len(books))
	for _, book := range books {
		bookList = append(bookList, book)
	}

	return c.JSON(fiber.Map{"status": "success", "data": bookList})
}

// GET /books/:id - Mendapatkan buku berdasarkan ID
func getBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "ID buku tidak valid"})
	}

	bookMutex.RLock()
	defer bookMutex.RUnlock()

	book, exists := books[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("Buku dengan ID %d tidak ditemukan", id)})
	}

	return c.JSON(fiber.Map{"status": "success", "data": book})
}

// POST /books - Membuat buku baru
func createBook(c *fiber.Ctx) error {
	input := new(CreateBookInput)

	// Parse & Validate Body
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Request body tidak valid"})
	}
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrorResponse(err))
	}

	bookMutex.Lock() // Lock untuk menulis
	defer bookMutex.Unlock()

	// Buat buku baru
	newBook := Book{
		ID:        nextBookID,
		Title:     input.Title,
		Author:    input.Author,
		CreatedAt: time.Now(),
	}

	// Simpan ke map
	books[nextBookID] = newBook
	nextBookID++ // Increment ID untuk buku berikutnya

	log.Printf("Buku baru dibuat: %+v", newBook)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": newBook})
}

// PUT /books/:id - Memperbarui buku
func updateBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "ID buku tidak valid"})
	}

	input := new(UpdateBookInput)
	// Parse & Validate Body
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Request body tidak valid"})
	}
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrorResponse(err))
	}

	bookMutex.Lock()
	defer bookMutex.Unlock()

	// Cek apakah buku ada
	book, exists := books[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("Buku dengan ID %d tidak ditemukan", id)})
	}

	// Update field jika ada di input (karena pakai pointer di UpdateBookInput)
	if input.Title != nil {
		book.Title = *input.Title
	}
	if input.Author != nil {
		book.Author = *input.Author
	}

	// Simpan buku yang sudah diupdate kembali ke map
	books[id] = book

	log.Printf("Buku diupdate: %+v", book)
	return c.JSON(fiber.Map{"status": "success", "data": book})
}

// DELETE /books/:id - Menghapus buku
func deleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "ID buku tidak valid"})
	}

	bookMutex.Lock()
	defer bookMutex.Unlock()

	// Cek apakah buku ada sebelum dihapus
	_, exists := books[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("Buku dengan ID %d tidak ditemukan", id)})
	}

	// Hapus buku dari map
	delete(books, id)

	log.Printf("Buku dengan ID %d dihapus", id)
	// Kirim 204 No Content untuk DELETE sukses
	return c.SendStatus(fiber.StatusNoContent)
}

// --- Setup & Main ---
func setupRoutes(app *fiber.App) {
	// Grup untuk API buku
	bookApi := app.Group("/books")

	bookApi.Get("/", getBooks)       // GET /books
	bookApi.Post("/", createBook)    // POST /books
	bookApi.Get("/:id", getBook)     // GET /books/:id
	bookApi.Put("/:id", updateBook)  // PUT /books/:id
	bookApi.Delete("/:id", deleteBook) // DELETE /books/:id
}

func main() {
	app := fiber.New()

	// Middleware
	app.Use(recover.New()) // Tangkap panic
	app.Use(logger.New())  // Log request

	// Setup rute
	setupRoutes(app)

	// Handle 404 untuk rute yang tidak cocok
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Endpoint '%s' dengan metode '%s' tidak ditemukan.", c.Path(), c.Method()),
		})
	})


	log.Println("Server berjalan di port 3000...")
	log.Fatal(app.Listen(":3000"))
}
```

**Cara Menjalankan dan Menguji:**

1.  Simpan kode di atas sebagai `main.go`.
2.  Jalankan: `go run main.go`.
3.  Gunakan `curl` atau Postman/Insomnia untuk menguji endpoint:
	*   **Create:** `curl -X POST http://localhost:3000/books -H "Content-Type: application/json" -d '{"title": "Go Programming Blueprints", "author": "Mat Ryer"}'`
	*   **Create (Invalid):** `curl -X POST http://localhost:3000/books -H "Content-Type: application/json" -d '{"title": "Go"}'` (Author required)
	*   **Get All:** `curl http://localhost:3000/books`
	*   **Get One:** `curl http://localhost:3000/books/1` (Ganti 1 dengan ID yang ada)
	*   **Update:** `curl -X PUT http://localhost:3000/books/1 -H "Content-Type: application/json" -d '{"author": "M. Ryer"}'` (Hanya update author)
	*   **Delete:** `curl -X DELETE http://localhost:3000/books/1`
	*   **Get Deleted:** `curl http://localhost:3000/books/1` (Harusnya 404 Not Found)
	*   **Not Found Route:** `curl http://localhost:3000/nonexistent` (Harusnya 404 JSON dari middleware)

Contoh ini mendemonstrasikan dasar-dasar CRUD, routing, parsing body, validasi input, penggunaan map dengan mutex untuk penyimpanan sederhana, dan penanganan error dasar.

---

## 8. Dokumentasi API 📚

Dokumentasi API yang baik sangat penting agar pengguna (frontend developer, tim lain, pengguna publik) dapat memahami cara menggunakan API Anda.

*   **Dokumentasi Kode Go (GoDoc):**
	*   Selalu tambahkan komentar pada fungsi publik, struct, interface, dan package Anda mengikuti format GoDoc.
	*   Gunakan `go doc ./...` untuk melihat dokumentasi secara lokal.
	*   Publikasikan ke [pkg.go.dev](https://pkg.go.dev/) agar dapat diakses publik.
	*   Ini penting untuk pengembang lain yang menggunakan library/kode Anda.

*   **Dokumentasi API RESTful (Spesifikasi):**
	*   Dokumentasikan setiap endpoint API:
		*   Path (e.g., `/users/:id`)
		*   Metode HTTP (GET, POST, PUT, DELETE)
		*   Deskripsi singkat tujuan endpoint.
		*   Parameter (path, query, header) - nama, tipe, wajib/opsional, deskripsi.
		*   Request Body (jika ada) - format (JSON, XML), contoh, deskripsi field, aturan validasi.
		*   Response Sukses - status code, format body, contoh, deskripsi field.
		*   Response Error - status code, format body, kemungkinan kode error, deskripsi.
		*   Autentikasi/Otorisasi yang diperlukan.
	*   **Alat Bantu:**
		*   **OpenAPI Specification (Swagger):** Standar industri untuk mendeskripsikan API RESTful dalam format JSON atau YAML.
			*   Banyak alat bantu untuk membuat UI interaktif dari spesifikasi OpenAPI (Swagger UI, Redoc).
			*   Bisa ditulis manual atau di-generate dari kode/komentar.
			*   Fiber memiliki integrasi (tidak resmi tapi populer) seperti `github.com/arsmn/fiber-swagger/v3` yang dapat meng-generate spesifikasi dari komentar atau struct dan menyajikan Swagger UI.
			*   Contoh komentar untuk `fiber-swagger`:
				```go
				// @Summary      Get a book by ID
				// @Description  Get details for a specific book
				// @Tags         Books
				// @Accept       json
				// @Produce      json
				// @Param        id   path      int  true  "Book ID"
				// @Success      200  {object}  BookResponseDoc // Definisikan struct doc terpisah jika perlu
				// @Failure      400  {object}  ErrorResponseDoc
				// @Failure      404  {object}  ErrorResponseDoc
				// @Router       /books/{id} [get]
				func getBook(c *fiber.Ctx) error { ... }
				```
		*   **API Blueprint:** Format berbasis Markdown lain untuk dokumentasi API.
		*   **Postman Collections:** Bisa diekspor sebagai dokumentasi dasar.

Pilih format dan alat bantu yang sesuai, tetapi **pastikan dokumentasi API Anda selalu diperbarui** seiring perubahan kode.

---

## 9. Praktik Terbaik (Best Practices) ✨

Beberapa tips dan praktik terbaik saat mengembangkan aplikasi dengan Go Fiber:

1.  **Struktur Proyek yang Jelas:** Gunakan struktur direktori yang logis (berdasarkan fitur atau lapisan) seiring pertumbuhan proyek.
2.  **Konfigurasi Terpusat:** Muat konfigurasi (port, DSN database, secret key) dari environment variables atau file konfigurasi, jangan hardcode. Gunakan library seperti `viper` atau struct dengan `envconfig`.
3.  **Penanganan Error Konsisten:** Gunakan `ErrorHandler` kustom untuk memformat semua response error secara seragam dan menyembunyikan detail internal di produksi. Log error secara detail di server.
4.  **Validasi Input Secara Menyeluruh:** Selalu validasi *semua* input dari client (body, query, params, headers) menggunakan library seperti `go-playground/validator`.
5.  **Gunakan Middleware dengan Bijak:** Manfaatkan middleware untuk tugas lintas sektoral (logging, auth, recovery, CORS). Terapkan pada lingkup yang tepat (global, grup, rute).
6.  **Dependency Injection (DI):** Gunakan DI untuk mengelola dependensi (seperti koneksi DB) agar kode lebih testable dan terstruktur. Hindari variabel global jika memungkinkan.
7.  **Tulis Tes:** Tulis unit test untuk logika bisnis/service/repository dan integration test untuk handler HTTP Anda. Targetkan cakupan kode yang baik.
8.  **Keamanan:**
	*   Gunakan HTTPS di produksi (terminasi SSL di reverse proxy).
	*   Lindungi dari serangan umum (SQLi, XSS, CSRF - Fiber tidak secara otomatis melindungi dari CSRF untuk aplikasi berbasis sesi/HTML form, perlu middleware CSRF).
	*   Jangan ekspos informasi sensitif di log atau response error.
	*   Jaga kerahasiaan secret key (JWT, API keys).
	*   Atur `BodyLimit` untuk mencegah DoS.
	*   Gunakan Autentikasi & Otorisasi yang kuat.
9.  **Logging yang Efektif:** Gunakan logger terstruktur (misalnya `zerolog`, `zap`) selain logger bawaan Fiber. Log informasi yang relevan (request ID, user ID jika ada) untuk memudahkan debugging. Atur level log yang berbeda untuk development dan produksi.
10. **Graceful Shutdown:** Implementasikan graceful shutdown untuk menangani penghentian server dengan benar.
11. **Kinerja:** Profil aplikasi Anda untuk menemukan bottleneck sebelum melakukan optimasi. Manfaatkan caching jika sesuai.
12. **Dokumentasi:** Tulis GoDoc untuk kode Anda dan dokumentasikan API Anda dengan jelas.
13. **Manajemen Dependensi:** Gunakan Go Modules (`go.mod`, `go.sum`) untuk mengelola dependensi. Perbarui dependensi secara berkala (perhatikan breaking changes).
14. **Gunakan Konteks (`context.Context`):** Teruskan `c.Context()` (yang merupakan `context.Context`) ke pemanggilan fungsi yang berjalan lama atau I/O-bound (seperti query DB, request HTTP keluar) untuk mendukung pembatalan (cancellation) dan timeout.

---

## 10. Berkontribusi 🤝

Fiber adalah proyek open-source. Kontribusi selalu diterima! Lihat panduan kontribusi resmi di repositori Fiber: [CONTRIBUTING.md](https://github.com/gofiber/fiber/blob/master/.github/CONTRIBUTING.md).

Cara berkontribusi:

*   Melaporkan bug.
*   Mengajukan permintaan fitur.
*   Menulis atau memperbaiki dokumentasi.
*   Mengirimkan Pull Request dengan perbaikan bug atau fitur baru (pastikan mengikuti gaya kode dan menyertakan tes).
*   Membantu menjawab pertanyaan di Diskusi GitHub atau komunitas lainnya.

---

## 11. Lisensi 📜

Go Fiber dirilis di bawah **Lisensi MIT**. Lihat file [LICENSE](https://github.com/gofiber/fiber/blob/master/LICENSE) untuk detail lengkap.

---

## 12. Ucapan Terima Kasih 🙏

*   Tim inti Go Fiber dan semua kontributornya.
*   Pencipta Fasthttp, Valyala.
*   Komunitas Express.js atas inspirasinya.
*   Komunitas Go yang luar biasa.

---

Semoga panduan lengkap ini bermanfaat dalam perjalanan Anda mempelajari dan menggunakan Go Fiber v2! Selamat mencoba! 🎉
