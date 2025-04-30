package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	ID   int
	Role string
}

func UserAuthMiddleware(c *fiber.Ctx) error {
	log.Println("UserAuthMiddleware: Mencoba mendapatkan data pengguna...")
	userID := c.QueryInt("user_id", 0)
	userRole := c.Query("role", "guest")

	if userID == 0 {
		log.Println("UserAuthMiddleware: User ID tidak ditemukan atau 0.")
	}

	user := User{ID: userID, Role: userRole}

	c.Locals("currentUser", user)
	c.Locals("requestID", "xyz-789-"+time.Now().String())

	log.Printf("UserAuthMiddleware: Pengguna ID %d (Role: %s) disimpan di Locals", user.ID, user.Role)
	return c.Next()
}

func GetUserProfile(c *fiber.Ctx) error {
	reqIDVal := c.Locals("requestID")
	reqID, ok := reqIDVal.(string)
	if !ok {
		log.Println("GetUserProfile: requestID bukan string atau nil")
		reqID = "unknown"
	}

	userVal := c.Locals("currentUser")
	user, ok := userVal.(User)

	log.Printf("GetUserProfile: Request ID = %s", reqID)

	if !ok {
		log.Println("GetUserProfile: Data pengguna (User) tidak ditemukan atau tipe salah di Locals!")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error internal: data pengguna hilang atau rusak",
			"request_id": reqID,
		})
	}

	log.Printf("GetUserProfile: Mengambil profil untuk pengguna ID %d (Role: %s)", user.ID, user.Role)
	return c.JSON(fiber.Map{
		"message":   "Profil Pengguna",
		"user_id":   user.ID,
		"user_role": user.Role,
		"request_id": reqID,
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/profile", UserAuthMiddleware, GetUserProfile)

	app.Get("/profile-guest", UserAuthMiddleware, GetUserProfile)


	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
