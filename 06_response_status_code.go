package main

import (
	"log"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var ErrItemNotFound = errors.New("item not found in storage")

func findItemByID(id string) (Item, error) {
	if id == "123" {
		return Item{ID: "123", Name: "Found Item"}, nil
	}
	if id == "error" {
		return Item{}, errors.New("simulated database error")
	}
	return Item{}, ErrItemNotFound
}

func deleteItemByID(id string) error {
	log.Printf("Deleting item with ID: %s", id)
	if id == "notfound" {
		return ErrItemNotFound
	}
	return nil
}


func main() {
	app := fiber.New()

	app.Post("/items", func(c *fiber.Ctx) error {
		log.Println("Creating a new item...")
		newItem := Item{ID: "5", Name: "Item Baru"}
		return c.Status(fiber.StatusCreated).JSON(newItem)
	})

	app.Get("/items/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		log.Printf("Finding item with ID: %s", id)
		item, err := findItemByID(id)

		if err != nil {
			log.Printf("Error finding item %s: %v", id, err)
			if errors.Is(err, ErrItemNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item tidak ditemukan"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengambil item"})
		}

		return c.JSON(item)
	})

	app.Delete("/items/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := deleteItemByID(id)
		if err != nil {
			log.Printf("Error deleting item %s: %v", id, err)
			if errors.Is(err, ErrItemNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item tidak ditemukan untuk dihapus"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus item"})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
