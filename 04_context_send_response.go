package main

import (
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
	"os"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := fiber.New()

	app.Post("/users", func(c *fiber.Ctx) error {
		newUser := User{ID: 1, Name: "Fiber User"}
		return c.Status(fiber.StatusCreated).JSON(newUser)
	})

	dummyFilePath := "./dummy_report.txt"
	_ = os.WriteFile(dummyFilePath, []byte("Ini isi file laporan dummy."), 0644)
	defer os.Remove(dummyFilePath)

	app.Get("/report", func(c *fiber.Ctx) error {
		log.Println("Sending file:", dummyFilePath)
		return c.SendFile(dummyFilePath)
	})

	app.Get("/download-report", func(c *fiber.Ctx) error {
		log.Println("Offering file for download:", dummyFilePath)
		return c.Download(dummyFilePath, "Laporan Keren Sekali.txt")
	})

	app.Get("/old-path", func(c *fiber.Ctx) error {
		log.Println("Redirecting from /old-path to /new-path")
		return c.Redirect("/new-path", fiber.StatusMovedPermanently)
	})

	app.Get("/new-path", func(c *fiber.Ctx) error {
		return c.SendString("Anda telah sampai di New Path!")
	})

	app.Get("/set-cookie", func(c *fiber.Ctx) error {
		cookie := fiber.Cookie{
			Name:     "session_id",
			Value:    "random-session-string-12345",
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "Lax",
		}
		c.Cookie(&cookie)
		log.Println("Cookie 'session_id' telah di-set.")
		return c.SendString("Cookie 'session_id' telah di-set!")
	})

	app.Get("/get-cookie", func(c *fiber.Ctx) error {
		sessionID := c.Cookies("session_id", "default-value")
		log.Printf("Nilai cookie 'session_id': %s", sessionID)
		return c.SendString("Nilai Cookie 'session_id': " + sessionID)
	})

	app.Get("/clear-cookie", func(c *fiber.Ctx) error {
		c.ClearCookie("session_id")
		log.Println("Cookie 'session_id' telah dihapus.")
		return c.SendString("Cookie 'session_id' telah dihapus.")
	})


	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
