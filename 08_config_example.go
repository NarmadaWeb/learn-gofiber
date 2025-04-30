package main

import (
	"log"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
	"errors"
)

func myCustomErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	log.Printf("Custom error handler caught: %v (Code: %d)", err, code)
	return c.Status(code).JSON(fiber.Map{"error": err.Error()})
}

func main() {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "My Awesome App v1.1 (Default)"
	}
	isProduction := os.Getenv("APP_ENV") == "production"


	config := fiber.Config{
		Prefork: isProduction,
		AppName: appName,
		ServerHeader: "MyWebServer",
		StrictRouting: false,
		CaseSensitive: false,
		BodyLimit: 10 * 1024 * 1024,
		ErrorHandler: myCustomErrorHandler,
		ReadBufferSize: 8192,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout: 60 * time.Second,
		EnablePrintRoutes: true,
	}

	app := fiber.New(config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from configured app!")
	})
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.ErrBadGateway
	})

	log.Printf("App Name: %s", app.Config().AppName)
	log.Printf("Prefork Enabled: %t", app.Config().Prefork)
	log.Printf("Body Limit: %d bytes", app.Config().BodyLimit)
	log.Printf("Read Timeout: %s", app.Config().ReadTimeout)
	log.Printf("Write Timeout: %s", app.Config().WriteTimeout)
	log.Printf("Idle Timeout: %s", app.Config().IdleTimeout)

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
