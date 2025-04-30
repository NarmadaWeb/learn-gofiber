package main

import (
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/data", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		c.Set("X-RateLimit-Limit", "100")
		c.Set("X-RateLimit-Remaining", "99")
		c.Set(fiber.HeaderCacheControl, "public, max-age=3600")

		c.Append(fiber.HeaderLink, "<http://example.com/meta>; rel=\"meta\"")
		c.Append(fiber.HeaderLink, "<http://example.com/next>; rel=\"next\"")

		log.Println("Sending response for /data with custom headers")
		return c.SendString("Ini data teks biasa dengan header kustom.")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		log.Println("Processing login, setting cookie...")
		c.Cookie(&fiber.Cookie{
			Name:     "session",
			Value:    "rahasia-user-123-" + time.Now().String(),
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "Lax",
			Path:     "/",
		})

		c.Append(fiber.HeaderSetCookie, "theme=dark; Path=/; Max-Age=86400")

		return c.JSON(fiber.Map{"message": "Login berhasil, cookie di-set"})
	})

	app.Get("/check-login", func(c *fiber.Ctx) error {
		sessionCookie := c.Cookies("session")
		themeCookie := c.Cookies("theme")
		log.Printf("Checking login: session='%s', theme='%s'", sessionCookie, themeCookie)
		if sessionCookie != "" {
			return c.JSON(fiber.Map{"loggedIn": true, "session": sessionCookie, "theme": themeCookie})
		}
		return c.JSON(fiber.Map{"loggedIn": false})
	})


	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
