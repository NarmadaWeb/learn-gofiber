package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/go-playground/validator/v10"
)

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateBookInput struct {
	Title  string `json:"title" validate:"required,min=3"`
	Author string `json:"author" validate:"required,min=3"`
}

type UpdateBookInput struct {
	Title  *string `json:"title" validate:"omitempty,min=3"`
	Author *string `json:"author" validate:"omitempty,min=3"`
}

var (
	books      = make(map[int]Book)
	nextBookID = 1
	bookMutex  = &sync.RWMutex{}
	validate   = validator.New()
)

var ErrBookNotFound = errors.New("buku tidak ditemukan")

type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}
func validationErrorResponse(err error) fiber.Map {
	var errors []ValidationErrorDetail
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			errors = append(errors, ValidationErrorDetail{
				Field: fieldErr.Field(),
				Tag: fieldErr.Tag(),
				Message: fmt.Sprintf("Field '%s' gagal validasi pada aturan '%s'", fieldErr.Field(), fieldErr.Tag()),
			})
		}
	} else {
		errors = append(errors, ValidationErrorDetail{Message: "Error validasi tidak dikenal: " + err.Error()})
	}
	return fiber.Map{"status": "fail", "message": "Data input tidak valid", "errors": errors}
}

func errorResponse(message string) fiber.Map {
	return fiber.Map{"status": "error", "message": message}
}

func failResponse(message string) fiber.Map {
	return fiber.Map{"status": "fail", "message": message}
}

func successResponse(data interface{}) fiber.Map {
	return fiber.Map{"status": "success", "data": data}
}

func getBooks(c *fiber.Ctx) error {
	bookMutex.RLock()
	defer bookMutex.RUnlock()

	bookList := make([]Book, 0, len(books))
	for _, book := range books {
		bookList = append(bookList, book)
	}

	sort.Slice(bookList, func(i, j int) bool {
		return bookList[i].ID < bookList[j].ID
	})

	return c.JSON(successResponse(bookList))
}

func getBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse("ID buku harus berupa angka"))
	}

	bookMutex.RLock()
	book, exists := books[id]
	bookMutex.RUnlock()

	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(failResponse(fmt.Sprintf("Buku dengan ID %d tidak ditemukan", id)))
	}

	return c.JSON(successResponse(book))
}

func createBook(c *fiber.Ctx) error {
	input := new(CreateBookInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse("Request body tidak valid: " + err.Error()))
	}

	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrorResponse(err))
	}

	bookMutex.Lock()
	defer bookMutex.Unlock()

	newBook := Book{
		ID:        nextBookID,
		Title:     input.Title,
		Author:    input.Author,
		CreatedAt: time.Now().UTC(),
	}

	books[nextBookID] = newBook
	currentID := nextBookID
	nextBookID++

	log.Printf("Buku baru dibuat: ID=%d, Title=%s, Author=%s", currentID, newBook.Title, newBook.Author)
	return c.Status(fiber.StatusCreated).JSON(successResponse(newBook))
}

func updateBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse("ID buku harus berupa angka"))
	}

	input := new(UpdateBookInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse("Request body tidak valid: " + err.Error()))
	}

	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrorResponse(err))
	}

	bookMutex.Lock()
	defer bookMutex.Unlock()

	book, exists := books[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(failResponse(fmt.Sprintf("Buku dengan ID %d tidak ditemukan", id)))
	}

	updated := false
	if input.Title != nil && *input.Title != "" && *input.Title != book.Title {
		book.Title = *input.Title
		updated = true
		log.Printf("Updating book %d: Title set to '%s'", id, book.Title)
	}
	if input.Author != nil && *input.Author != "" && *input.Author != book.Author {
		book.Author = *input.Author
		updated = true
		log.Printf("Updating book %d: Author set to '%s'", id, book.Author)
	}

	if !updated {
		log.Printf("Tidak ada perubahan data untuk buku ID %d", id)
		return c.JSON(successResponse(book))
	}

	books[id] = book

	log.Printf("Buku ID %d diupdate: %+v", id, book)
	return c.JSON(successResponse(book))
}

func deleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse("ID buku harus berupa angka"))
	}

	bookMutex.Lock()
	defer bookMutex.Unlock()

	_, exists := books[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(failResponse(fmt.Sprintf("Buku dengan ID %d tidak ditemukan", id)))
	}

	delete(books, id)

	log.Printf("Buku dengan ID %d dihapus", id)
	return c.SendStatus(fiber.StatusNoContent)
}

func setupRoutes(app *fiber.App) {
	bookApi := app.Group("/books")

	bookApi.Get("/", getBooks)
	bookApi.Post("/", createBook)
	bookApi.Get("/:id<int>", getBook)
	bookApi.Put("/:id<int>", updateBook)
	bookApi.Delete("/:id<int>", deleteBook)

	bookApi.Get("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse("ID buku harus berupa angka"))
	})
	bookApi.Put("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse("ID buku harus berupa angka"))
	})
	bookApi.Delete("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse("ID buku harus berupa angka"))
	})
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "Terjadi kesalahan server internal"

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			} else {
				log.Printf("[DefaultErrorHandler] Non-Fiber error: %v", err)
			}
			log.Printf("[DefaultErrorHandler] Path: %s, Error: %v", c.Path(), err)
			return c.Status(code).JSON(errorResponse(message))
		},
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${ip}\n",
	}))

	books[nextBookID] = Book{ID: nextBookID, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", CreatedAt: time.Now().UTC().Add(-time.Hour)}
	nextBookID++
	books[nextBookID] = Book{ID: nextBookID, Title: "Concurrency in Go", Author: "Katherine Cox-Buday", CreatedAt: time.Now().UTC()}
	nextBookID++

	setupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(failResponse(
			fmt.Sprintf("Endpoint '%s' dengan metode '%s' tidak ditemukan.", c.OriginalURL(), c.Method()),
		))
	})


	log.Println("CRUD Example Server berjalan di port 3000...")
	log.Fatal(app.Listen(":3000"))
}
