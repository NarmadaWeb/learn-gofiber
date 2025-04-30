package main

import (
	"errors"
	"log"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var ErrRecordNotFound = errors.New("record not found")

var database = struct {
	FindItem func(id string) (Item, error)
}{
	FindItem: func(id string) (Item, error) {
		log.Printf("Database: Mencari item ID %s", id)
		if id == "found" {
			return Item{ID: "found", Name: "Item Ditemukan"}, nil
		}
		if id == "notfound" {
			return Item{}, ErrRecordNotFound
		}
		return Item{}, errors.New("koneksi database gagal")
	},
}


func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := database.FindItem(id)

	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			log.Printf("Item %s tidak ditemukan: %v", id, err)
			return fiber.ErrNotFound
		}
		log.Printf("Error database saat mencari item %s: %v", id, err)
		return fiber.NewError(fiber.StatusServiceUnavailable, "Gagal mengambil data item dari database")
	}

	return c.JSON(item)
}

func main() {
	app := fiber.New()

	app.Get("/items/:id", GetItem)

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
