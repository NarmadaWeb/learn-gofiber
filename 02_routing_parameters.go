package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/users/:userId/books/:bookId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		bookId := c.Params("bookId")

		return c.SendString("User ID: " + userId + ", Book ID: " + bookId)
	})

	app.Get("/product/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("ID produk harus berupa angka")
		}
		log.Printf("Processing product with ID: %d", id)
		return c.JSON(fiber.Map{"product_id": id, "status": "found"})
	})

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
