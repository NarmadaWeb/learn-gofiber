package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"

	"my-fiber-guide-examples/internal/handlers"
)

func initDatabase() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
		log.Println("WARNING: DATABASE_URL environment variable not set. Using default DSN:", dsn)
	}

	var err error
	db, err := sql.Open("postgres", dsn)
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
	if err != nil { log.Printf("Warning: Gagal membuat tabel products: %v", err)
	} else { log.Println("Tabel 'products' siap.") }


	return db, nil
}

func main() {
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("FATAL: Gagal menginisialisasi database: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New())

	productHandler := handlers.NewProductHandler(db)

	api := app.Group("/products")
	api.Get("/", productHandler.GetProducts)
	api.Post("/", productHandler.CreateProduct)
	api.Get("/:id", productHandler.GetProductByID)

	log.Println("Starting server (DI Example) on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
