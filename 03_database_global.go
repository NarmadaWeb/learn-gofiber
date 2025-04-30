package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB

func initDatabase() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
		log.Println("WARNING: DATABASE_URL environment variable not set. Using default DSN:", dsn)
	}

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi db: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}

	log.Println("Koneksi database berhasil dibuat!")

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price NUMERIC(10, 2) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.ExecContext(context.Background(), createTableSQL)
	if err != nil {
		log.Printf("Warning: Gagal membuat tabel products: %v", err)
	} else {
		log.Println("Tabel 'products' siap.")
	}


	return db, nil
}

func getProductsHandler(c *fiber.Ctx) error {
	rows, err := db.QueryContext(c.Context(), "SELECT id, name, price, created_at FROM products ORDER BY created_at DESC")
	if err != nil {
		log.Printf("Error query produk: %v", err)
		return fiber.ErrInternalServerError
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
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

func createProductHandler(c *fiber.Ctx) error {
	input := new(struct {
		Name  string  `json:"name" validate:"required"`
		Price float64 `json:"price" validate:"required,gt=0"`
	})

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body tidak valid"})
	}

	var newID int
	var createdAt time.Time
	err := db.QueryRowContext(c.Context(),
		"INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, created_at",
		input.Name, input.Price,
	).Scan(&newID, &createdAt)

	if err != nil {
		log.Printf("Error insert produk: %v", err)
		return fiber.ErrInternalServerError
	}

	newProduct := Product{
		ID:        newID,
		Name:      input.Name,
		Price:     input.Price,
		CreatedAt: createdAt,
	}
	log.Printf("Produk baru dibuat: %+v", newProduct)
	return c.Status(fiber.StatusCreated).JSON(newProduct)
}

func getProductByIDHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID produk tidak valid"})
	}

	var p Product
	query := "SELECT id, name, price, created_at FROM products WHERE id = $1"
	err = db.QueryRowContext(c.Context(), query, id).Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Produk ID %d tidak ditemukan", id)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("Produk dengan ID %d tidak ditemukan", id)})
		}
		log.Printf("Error query produk ID %d: %v", id, err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(p)
}


func main() {
	var err error
	db, err = initDatabase()
	if err != nil {
		log.Fatalf("FATAL: Gagal menginisialisasi database: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/products")
	api.Get("/", getProductsHandler)
	api.Post("/", createProductHandler)
	api.Get("/:id", getProductByIDHandler)

	log.Println("Starting server on port 3000...")
	log.Println("Pastikan PostgreSQL berjalan dan DATABASE_URL di-set, atau sesuaikan DSN default.")
	log.Fatal(app.Listen(":3000"))
}
