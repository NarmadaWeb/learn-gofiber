package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

type SearchQuery struct {
	Query string `query:"q"`
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
}

type CreateUserInput struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
}


func main() {
	app := fiber.New()

	app.Get("/search", func(c *fiber.Ctx) error {
		var sq SearchQuery
		if err := c.QueryParser(&sq); err != nil {
			log.Printf("Query parsing error: %v", err)
			return c.Status(fiber.StatusBadRequest).SendString("Query tidak valid: " + err.Error())
		}
		log.Printf("Search query parsed: %+v", sq)
		return c.JSON(fiber.Map{
			"message": "Query parsed successfully",
			"data":    sq,
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var input CreateUserInput
		if err := c.BodyParser(&input); err != nil {
			log.Printf("Body parsing error: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body tidak valid: " + err.Error()})
		}
		log.Printf("User input parsed: %+v", input)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User created",
			"user":    input,
		})
	})

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
