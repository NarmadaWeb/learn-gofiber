package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"time"
	"fmt"
)

type BookResponseDoc struct {
	Status string      `json:"status" example:"success"`
	Data   BookDataDoc `json:"data"`
}
type BookDataDoc struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"The Go Programming Language"`
	Author    string    `json:"author" example:"Alan A. A. Donovan"`
	CreatedAt time.Time `json:"created_at" example:"2023-10-27T10:00:00Z"`
}

type ErrorResponseDoc struct {
	Status  string `json:"status" example:"fail"`
	Message string `json:"message" example:"Buku tidak ditemukan"`
}

func getBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponseDoc{Status: "fail", Message: "ID buku harus berupa angka"})
	}
	log.Printf("Handler getBook dipanggil untuk ID: %d (Swagger example)", id)

	if id == 1 {
		book := BookDataDoc {
			ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", CreatedAt: time.Now(),
		}
		return c.JSON(BookResponseDoc{Status:"success", Data: book})
	}

	return c.Status(fiber.StatusNotFound).JSON(ErrorResponseDoc{Status: "fail", Message: fmt.Sprintf("Buku dengan ID %d tidak ditemukan", id)})
}

func main() {
	app := fiber.New()

	app.Get("/books/:id", getBook)

	log.Println("Starting server with Swagger annotation example on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
