package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

type ProductFilter struct {
	Category string `query:"category"`
	MaxPrice int    `query:"max_price"`
	SortBy   string `query:"sort"`
}

type UpdateProductInput struct {
	Name        string  `json:"name" form:"name"`
	Description *string `json:"description" form:"description"`
	Price       float64 `json:"price" form:"price" validate:"required,gt=0"`
	IsActive    bool    `json:"is_active" form:"is_active"`
}

type ProductRouteParams struct {
	ProductID int `params:"id"`
}

func SearchProducts(c *fiber.Ctx) error {
	var filter ProductFilter
	if err := c.QueryParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Filter tidak valid: " + err.Error()})
	}
	log.Printf("Searching products with filter: %+v", filter)
	return c.JSON(fiber.Map{"message": "Hasil pencarian", "filters": filter})
}

func UpdateProduct(c *fiber.Ctx) error {
	var params ProductRouteParams
	if err := c.ParamsParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID Produk tidak valid: " + err.Error()})
	}

	var input UpdateProductInput
	if err := c.BodyParser(&input); err != nil {
		if err == fiber.ErrUnprocessableEntity || err.Error() == "Unprocessable Entity" {
			 return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body request kosong atau format salah"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Data input tidak valid: " + err.Error()})
	}

	log.Printf("Memperbarui produk ID %d dengan data: %+v", params.ProductID, input)
	return c.JSON(fiber.Map{"message": "Produk berhasil diperbarui", "id": params.ProductID, "updated_data": input})
}

func main() {
	app := fiber.New()

	app.Get("/products/search", SearchProducts)
	app.Put("/products/:id", UpdateProduct)

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
