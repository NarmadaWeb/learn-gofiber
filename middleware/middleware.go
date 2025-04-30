package middleware

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getEnv("JWT_SECRET_KEY", "kunci-rahasia-super-aman-jangan-hardcode"))

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Printf("Warning: Environment variable %s not set, using default value.", key)
	return fallback
}

type MyCustomClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, role string) (string, error) {
	claims := MyCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app",
			Subject:   fmt.Sprint(userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}

	return signedToken, nil
}

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Println("Protected Middleware: Header Authorization tidak ada")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": "fail",
				"message": "Header Authorization dibutuhkan",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			log.Printf("Protected Middleware: Format header Authorization salah: %s", authHeader)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": "fail",
				"message": "Format header Authorization salah (harus 'Bearer <token>')",
			})
		}
		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("Protected Middleware: Metode signing token tidak valid: %v", token.Header["alg"])
				return nil, fmt.Errorf("metode signing token tidak valid: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			log.Printf("Protected Middleware: Error parsing/validasi token: %v", err)
			if errors.Is(err, jwt.ErrTokenExpired) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status": "fail",
					"message": "Token telah kedaluwarsa",
					"code": "TOKEN_EXPIRED",
				})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": "fail",
				"message": "Token tidak valid",
				"code": "TOKEN_INVALID",
				"details": err.Error(),
			})
		}

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			c.Locals("userID", claims.UserID)
			c.Locals("userRole", claims.Role)
			c.Locals("jwtClaims", claims)
			log.Printf("Protected Middleware: Akses diberikan untuk User ID %d (Role: %s)", claims.UserID, claims.Role)
			return c.Next()
		}

		log.Println("Protected Middleware: Token tidak valid atau claims rusak setelah parsing")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "fail",
			"message": "Token tidak valid atau claims rusak",
			"code": "TOKEN_CLAIMS_INVALID",
		})
	}
}

func AuthorizeRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleVal := c.Locals("userRole")
		if roleVal == nil {
			log.Println("Authorize Middleware: Role pengguna (userRole) tidak ditemukan di Locals")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status": "fail",
				"message": "Akses ditolak (role tidak diketahui)",
				"code": "AUTHZ_ROLE_MISSING",
				})
		}

		role, ok := roleVal.(string)
		if !ok {
			log.Printf("Authorize Middleware: Role pengguna (userRole) di Locals bukan string (%T)", roleVal)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
				"message": "Tipe data role tidak sesuai",
				"code": "INTERNAL_ROLE_TYPE_ERROR",
			})
		}

		allowedMap := make(map[string]struct{}, len(allowedRoles))
		for _, allowed := range allowedRoles {
			allowedMap[allowed] = struct{}{}
		}

		if _, found := allowedMap[role]; !found {
			log.Printf("Authorize Middleware: Akses ditolak untuk role '%s'. Role yang diizinkan: %v", role, allowedRoles)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status": "fail",
				"message": "Anda tidak memiliki izin yang cukup untuk mengakses sumber daya ini",
				"code": "AUTHZ_INSUFFICIENT_PERMISSION",
			})
		}

		log.Printf("Authorize Middleware: Akses diizinkan untuk role '%s'", role)
		return c.Next()
	}
}
