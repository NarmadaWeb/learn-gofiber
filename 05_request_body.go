package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

type CreatePostInput struct {
	Title   string   `json:"title" form:"title" xml:"title" validate:"required"`
	Content string   `json:"content" form:"content" xml:"content" validate:"required"`
	Tags    []string `json:"tags" form:"tags" xml:"tags"`
}

func CreatePostHandler(c *fiber.Ctx) error {
	var input CreatePostInput

	if err := c.BodyParser(&input); err != nil {
		log.Printf("Error parsing body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Gagal memproses body request",
			"details": err.Error(),
		})
	}

	log.Printf("Membuat post baru: Title='%s', Content='%s', Tags=%v", input.Title, input.Content, input.Tags)

	newPost := fiber.Map{
		"id": 1,
		"title": input.Title,
		"content": input.Content,
		"tags": input.Tags,
		"status": "created",
	}
	return c.Status(fiber.StatusCreated).JSON(newPost)
}

func verifySignature(body []byte, signature string) bool {
	log.Printf("Verifying signature '%s' for body: %s", signature, string(body))
	return signature == "secret-signature"
}

func WebhookHandler(c *fiber.Ctx) error {
	rawBody := c.Body()

	signature := c.Get("X-Webhook-Signature")
	if !verifySignature(rawBody, signature) {
		log.Println("Webhook signature tidak valid")
		return c.Status(fiber.StatusUnauthorized).SendString("Signature tidak valid")
	}
	log.Println("Webhook signature valid")

	var payload map[string]interface{}
	if err := fiber.Unmarshal(rawBody, &payload); err != nil {
		log.Printf("Error unmarshaling webhook payload: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Gagal parse payload JSON setelah membaca body mentah")
	}

	log.Printf("Webhook diterima dan diproses: %+v", payload)

	return c.SendStatus(fiber.StatusOK)
}


func main() {
	app := fiber.New()

	app.Post("/posts", CreatePostHandler)
	app.Post("/webhook", WebhookHandler)

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
