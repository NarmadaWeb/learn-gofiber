package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

type SearchParams struct {
	Query    string `query:"q"`
	Page     int    `query:"page" default:"1"`
	Limit    int    `query:"limit" default:"10"`
	Sort     string `query:"sort"`
	ShowMeta bool   `query:"show_meta"`
}

func SearchHandler(c *fiber.Ctx) error {
	searchTerm := c.Query("q")
	pageStr := c.Query("page", "1")
	limit, err := c.QueryInt("limit", 10)
	if err != nil {
		log.Printf("Error parsing limit (manual): %v - value used: %d", err, limit)
	}

	log.Printf("Manual Parsing - Mencari '%s', Halaman: %s, Limit: %d", searchTerm, pageStr, limit)

	var params SearchParams
	if err := c.QueryParser(&params); err != nil {
		log.Printf("Struct parsing error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Query parameter tidak valid: " + err.Error()})
	}
	log.Printf("Struct Binding - Mencari '%s', Halaman: %d, Limit: %d, Sort: '%s', Meta: %t",
		params.Query, params.Page, params.Limit, params.Sort, params.ShowMeta)

	return c.JSON(fiber.Map{"status": "success", "params_used": params})
}

func main() {
	app := fiber.New()

	app.Get("/search", SearchHandler)

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
