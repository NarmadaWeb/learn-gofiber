package main

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func MyCustomErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Terjadi kesalahan internal pada server."

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	} else {
		log.Printf("[ErrorHandler - NonFiberError] Error asli: %v", err)
	}

	log.Printf("[ErrorHandler] Status: %d, Error: %v, Path: %s, Raw Error Type: %T", code, err, c.Path(), err)

	isProduction := os.Getenv("APP_ENV") == "production"
	if isProduction && code >= 500 {
		message = "Maaf, terjadi kesalahan tak terduga di server."
	}

	if c.Get(fiber.HeaderContentType) == "" {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	}

	errorResponse := fiber.Map{
		"status":  "error",
		"code":    code,
		"message": message,
	}
	if !isProduction || code < 500 {
		errorResponse["details"] = err.Error()
	}

	return c.Status(code).JSON(errorResponse)
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: MyCustomErrorHandler,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/notfound", func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	app.Get("/badrequest", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusBadRequest, "Parameter 'q' dibutuhkan.")
	})

	app.Get("/dberror", func(c *fiber.Ctx) error {
		simulatedError := errors.New("database connection failed")
		return simulatedError
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Sesuatu yang sangat salah terjadi!")
	})


	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
