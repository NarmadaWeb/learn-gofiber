package main

import (
	"fmt"
	"log"
	"my-fiber-guide-examples/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

var usersDB = map[string]User{
	"admin": {ID: 1, Username: "admin", Password: "password123", Role: "admin"},
	"editor":{ID: 2, Username: "editor", Password: "password456", Role: "editor"},
	"user":  {ID: 3, Username: "user", Password: "password789", Role: "user"},
}

func loginHandler(c *fiber.Ctx) error {
	input := new(LoginInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Input tidak valid"})
	}

	user, found := usersDB[input.Username]
	if !found {
		log.Printf("Login attempt failed: User '%s' not found", input.Username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Username atau password salah"})
	}

	if user.Password != input.Password {
		log.Printf("Login attempt failed: Invalid password for user '%s'", input.Username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Username atau password salah"})
	}

	token, err := middleware.GenerateJWT(user.ID, user.Role)
	if err != nil {
		log.Printf("Error generating JWT for user %d: %v", user.ID, err)
		return fiber.ErrInternalServerError
	}

	log.Printf("Login successful for user '%s' (ID: %d, Role: %s)", user.Username, user.ID, user.Role)
	return c.JSON(fiber.Map{"status": "success", "token": token})
}

func getMeHandler(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	userRole := c.Locals("userRole").(string)

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user_id": userID,
			"role":    userRole,
		},
	})
}

func adminGetUsersHandler(c *fiber.Ctx) error {
	log.Println("Admin handler '/api/admin/users' accessed")
	userList := []fiber.Map{}
	for _, u := range usersDB {
		userList = append(userList, fiber.Map{"id": u.ID, "username": u.Username, "role": u.Role})
	}
	return c.JSON(fiber.Map{"status": "success", "data": userList})
}

func createArticleHandler(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	userRole := c.Locals("userRole").(string)
	log.Printf("Article creation handler accessed by User ID %d (Role: %s)", userID, userRole)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Artikel baru sedang dibuat oleh User %d (Role: %s)", userID, userRole),
	})
}


func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	auth := app.Group("/auth")
	auth.Post("/login", loginHandler)

	api := app.Group("/api")
	api.Use(middleware.Protected())

	api.Get("/me", getMeHandler)

	api.Get("/public-data", func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(int)
		return c.JSON(fiber.Map{"message": "Ini data publik untuk user terautentikasi", "accessed_by": userID})
	})

	adminApi := api.Group("/admin")
	adminApi.Use(middleware.AuthorizeRole("admin"))
	adminApi.Get("/users", adminGetUsersHandler)
	adminApi.Delete("/users/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": fmt.Sprintf("Menghapus user ID %s (action by admin)", c.Params("id"))})
	})

	api.Post("/articles", middleware.AuthorizeRole("admin", "editor"), createArticleHandler)


	log.Println("Starting JWT Auth Example Server on port 3000...")
	log.Println("Set JWT_SECRET_KEY environment variable for production.")
	log.Println("Endpoints:")
	log.Println("  POST /auth/login (public)")
	log.Println("  GET /api/me (protected)")
	log.Println("  GET /api/admin/users (protected, admin only)")
	log.Println("  POST /api/articles (protected, admin or editor)")

	log.Fatal(app.Listen(":3000"))
}
