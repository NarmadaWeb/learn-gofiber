package main_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	mainApp "my-fiber-guide-examples"
	"my-fiber-guide-examples/internal/handlers"
	"my-fiber-guide-examples/internal/models"

	_ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	log.Println("Setting up database for tests...")
	testDSN := os.Getenv("TEST_DATABASE_URL")
	if testDSN == "" {
		testDSN = "postgres://user:password@localhost:5433/testdb?sslmode=disable"
		log.Println("WARNING: TEST_DATABASE_URL env var not set, using default:", testDSN)
	}

	var err error
	testDB, err = sql.Open("postgres", testDSN)
	if err != nil {
		log.Fatalf("FATAL: Failed to open test database connection: %v", err)
	}

	testDB.SetMaxOpenConns(10)
	testDB.SetMaxIdleConns(5)
	testDB.SetConnMaxLifetime(1 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = testDB.PingContext(ctx)
	if err != nil {
		testDB.Close()
		log.Fatalf("FATAL: Failed to ping test database: %v", err)
	}

	log.Println("Test database connected. Running tests...")
	exitCode := m.Run()

	log.Println("Closing test database connection...")
	if err := testDB.Close(); err != nil {
		log.Printf("ERROR: Failed to close test database connection: %v", err)
	}
	log.Println("Test database connection closed.")

	os.Exit(exitCode)
}

func clearProductsTable(t *testing.T) {
	t.Helper()
	_, err := testDB.ExecContext(context.Background(), "TRUNCATE TABLE products RESTART IDENTITY CASCADE")
	require.NoError(t, err, "Failed to clear products table")
	log.Println("Products table truncated for test:", t.Name())
}

func setupTestApp() *fiber.App {
	app := fiber.New()
	productHandler := handlers.NewProductHandler(testDB)
	app.Post("/products", productHandler.CreateProduct)
	app.Get("/products", productHandler.GetProducts)
	app.Get("/products/:id", productHandler.GetProductByID)
	return app
}

func performRequestTestDB(t *testing.T, app *fiber.App, method, path string, body io.Reader, headers map[string]string) *http.Response {
	req := httptest.NewRequest(method, path, body)
	if headers != nil {
		for k, v := range headers { req.Header.Add(k, v) }
	}
	resp, err := app.Test(req, -1)
	require.NoError(t, err, "app.Test failed during DB test")
	return resp
}

func TestCreateProductHandler_Integration(t *testing.T) {
	clearProductsTable(t)
	app := setupTestApp()

	inputData := map[string]interface{}{
		"name": "Test Product DB",
		"price": 123.45,
	}
	bodyBytes, _ := json.Marshal(inputData)
	bodyReader := bytes.NewReader(bodyBytes)
	headers := map[string]string{"Content-Type": fiber.MIMEApplicationJSON}

	resp := performRequestTestDB(t, app, "POST", "/products", bodyReader, headers)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Should return 201 Created")

	var createdProduct models.Product
	err := json.NewDecoder(resp.Body).Decode(&createdProduct)
	require.NoError(t, err, "Should be able to decode response JSON")
	assert.Equal(t, inputData["name"], createdProduct.Name)
	assert.Equal(t, inputData["price"], createdProduct.Price)
	assert.NotZero(t, createdProduct.ID, "Product ID should be generated by DB")
	assert.NotZero(t, createdProduct.CreatedAt, "CreatedAt should be set by DB")

	var dbCount int
	var dbName string
	err = testDB.QueryRowContext(context.Background(),
		"SELECT name, COUNT(*) FROM products WHERE id = $1 GROUP BY name", createdProduct.ID).Scan(&dbName, &dbCount)
	require.NoError(t, err, "Failed to query created product from DB")
	assert.Equal(t, 1, dbCount, "Product count in DB should be 1")
	assert.Equal(t, createdProduct.Name, dbName, "Product name in DB should match created product")
}

func TestGetProductsHandler_Integration(t *testing.T) {
	clearProductsTable(t)
	app := setupTestApp()

	_, err := testDB.ExecContext(context.Background(), "INSERT INTO products (name, price) VALUES ($1, $2), ($3, $4)",
		"Product A", 10.0, "Product B", 20.0)
	require.NoError(t, err, "Failed to seed data")

	resp := performRequestTestDB(t, app, "GET", "/products", nil, nil)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Should return 200 OK")

	var products []models.Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	require.NoError(t, err, "Should decode product list JSON")
	assert.Len(t, products, 2, "Should return 2 products")
	assert.Equal(t, "Product B", products[0].Name, "First product should be Product B (ordered by created_at desc implicitly)")
	assert.Equal(t, "Product A", products[1].Name, "Second product should be Product A")

}

func TestGetProductByIDHandler_Integration(t *testing.T) {
	clearProductsTable(t)
	app := setupTestApp()

	var insertedID int
	err := testDB.QueryRowContext(context.Background(),
		"INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id", "Specific Product", 55.5).Scan(&insertedID)
	require.NoError(t, err, "Failed to seed specific product")
	require.NotZero(t, insertedID, "Inserted ID should not be zero")

	pathFound := fmt.Sprintf("/products/%d", insertedID)
	respFound := performRequestTestDB(t, app, "GET", pathFound, nil, nil)
	defer respFound.Body.Close()

	assert.Equal(t, http.StatusOK, respFound.StatusCode, "[Found] Should return 200 OK")
	var product models.Product
	err = json.NewDecoder(respFound.Body).Decode(&product)
	require.NoError(t, err, "[Found] Should decode product JSON")
	assert.Equal(t, insertedID, product.ID, "[Found] ID should match")
	assert.Equal(t, "Specific Product", product.Name, "[Found] Name should match")

	pathNotFound := "/products/99999"
	respNotFound := performRequestTestDB(t, app, "GET", pathNotFound, nil, nil)
	defer respNotFound.Body.Close()

	assert.Equal(t, http.StatusNotFound, respNotFound.StatusCode, "[NotFound] Should return 404 Not Found")
}

var _ = mainApp.Placeholder
