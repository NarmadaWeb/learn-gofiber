package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"my-fiber-guide-examples/internal/models"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	DB *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{DB: db}
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	rows, err := h.DB.QueryContext(c.Context(), "SELECT id, name, price, created_at FROM products ORDER BY created_at DESC")
	if err != nil {
		log.Printf("Error query produk: %v", err)
		return fiber.ErrInternalServerError
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
			log.Printf("Error scan row produk: %v", err)
			return fiber.ErrInternalServerError
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterasi rows produk: %v", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(products)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	input := new(struct {
		Name  string  `json:"name" validate:"required"`
		Price float64 `json:"price" validate:"required,gt=0"`
	})

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body tidak valid"})
	}

	var newID int
	var createdAt time.Time
	err := h.DB.QueryRowContext(c.Context(),
		"INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, created_at",
		input.Name, input.Price,
	).Scan(&newID, &createdAt)

	if err != nil {
		log.Printf("Error insert produk: %v", err)
		return fiber.ErrInternalServerError
	}

	newProduct := models.Product{
		ID:        newID,
		Name:      input.Name,
		Price:     input.Price,
		CreatedAt: createdAt,
	}
	log.Printf("Produk baru dibuat (DI): %+v", newProduct)
	return c.Status(fiber.StatusCreated).JSON(newProduct)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID produk tidak valid"})
	}

	var p models.Product
	query := "SELECT id, name, price, created_at FROM products WHERE id = $1"
	err = h.DB.QueryRowContext(c.Context(), query, id).Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Produk ID %d tidak ditemukan (DI)", id)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("Produk dengan ID %d tidak ditemukan", id)})
		}
		log.Printf("Error query produk ID %d (DI): %v", id, err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(p)
}
