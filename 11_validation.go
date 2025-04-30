package main

import (
	"log"
	"strings"
	"fmt" // For message formatting

	"github.com/go-playground/validator/v10" // Import validator
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Global instance validator (atau buat per request jika perlu konfigurasi berbeda)
var validate = validator.New()

// --- Struct Input ---
type RegisterUserInput struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Age      *int   `json:"age" validate:"omitempty,gte=18,lte=120"` // Pointer + omitempty for optional validation
	Website  string `json:"website" validate:"omitempty,url"`        // omitempty: validasi hanya jika field tidak kosong
	UserType string `json:"user_type" validate:"required,oneof=admin user guest"` // Harus salah satu dari nilai ini
	Terms    bool   `json:"terms" validate:"required,eq=true"` // Harus true
}


// --- Error Handling ---
// Struct untuk response error validasi yang lebih informatif
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   any    `json:"value,omitempty"` // Nilai yang gagal divalidasi (hati-hati expose data sensitif)
	Message string `json:"message"`
}

// Fungsi helper untuk memformat error validasi
func formatValidationErrors(err error) []ValidationErrorResponse {
	var errors []ValidationErrorResponse

	// Cek apakah error adalah tipe ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			// Buat pesan error yang lebih user-friendly (contoh sederhana)
			var message string
			field := fieldErr.Field()
			tag := fieldErr.Tag()
			param := fieldErr.Param()
			// val := fieldErr.Value() // Nilai yang gagal

			switch tag {
			case "required":
				message = fmt.Sprintf("Field '%s' wajib diisi.", field)
			case "email":
				message = fmt.Sprintf("Field '%s' harus berupa format email yang valid.", field)
			case "url":
				message = fmt.Sprintf("Field '%s' harus berupa URL yang valid.", field)
			case "min":
				message = fmt.Sprintf("Field '%s' minimal harus %s karakter/nilai.", field, param)
			case "max":
				message = fmt.Sprintf("Field '%s' maksimal harus %s karakter/nilai.", field, param)
			case "gte":
				message = fmt.Sprintf("Field '%s' minimal harus bernilai %s.", field, param)
			case "lte":
				message = fmt.Sprintf("Field '%s' maksimal harus bernilai %s.", field, param)
			case "alphanum":
				message = fmt.Sprintf("Field '%s' hanya boleh berisi huruf dan angka.", field)
			case "oneof":
				message = fmt.Sprintf("Field '%s' harus salah satu dari: %s", field, strings.Replace(param, " ", ", ", -1))
			case "eq":
				message = fmt.Sprintf("Field '%s' harus bernilai '%s'.", field, param)
			default:
				message = fmt.Sprintf("Field '%s' tidak valid (aturan: %s).", field, tag)
			}

			errors = append(errors, ValidationErrorResponse{
				Field:   field,     // Nama field struct
				Tag:     tag,       // Tag validasi yang gagal
				// Value:   val, // Hati-hati menampilkan value
				Message: message,   // Pesan kustom
			})
		}
	} else {
		// Jika error bukan ValidationErrors (jarang terjadi jika inputnya err dari validate.Struct)
		log.Printf("Warning: Error validasi tidak terduga: %v (%T)", err, err)
		errors = append(errors, ValidationErrorResponse{Message: "Error validasi tidak dikenal: " + err.Error()})
	}
	return errors
}

// --- Handler ---
// Handler untuk registrasi pengguna
func RegisterUserHandler(c *fiber.Ctx) error {
	input := new(RegisterUserInput) // Gunakan new() agar dapat pointer

	// 1. Bind body request ke struct
	if err := c.BodyParser(input); err != nil {
		log.Printf("Binding error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Gagal memproses body request.",
			"details": err.Error(),
		})
	}
	log.Printf("Input received: %+v", input)
	// Perhatikan nilai Age jika tidak dikirim: nil
	if input.Age != nil {
		log.Printf("Input Age: %d", *input.Age)
	} else {
		log.Println("Input Age: nil")
	}


	// 2. Lakukan validasi pada struct yang sudah di-bind
	err := validate.Struct(input)
	if err != nil {
		// Jika validasi gagal, format error dan kirim 400 Bad Request
		validationErrors := formatValidationErrors(err)
		log.Printf("Validasi gagal untuk registrasi: %v", validationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Data yang diberikan tidak valid.",
			"errors":  validationErrors, // Kirim detail error validasi
		})
	}

	// 3. Jika validasi berhasil, lanjutkan proses
	log.Printf("Registrasi valid diterima: Username=%s, Email=%s", input.Username, input.Email)
	// ... logika menyimpan pengguna ke database ...

	// Kirim response sukses
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Pengguna berhasil diregistrasi.",
		"user": fiber.Map{ // Jangan kirim password kembali!
			"username": input.Username,
			"email":    input.Email,
			"age":      input.Age, // Bisa jadi nil
			"website": input.Website,
			"user_type": input.UserType,
		},
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/register", RegisterUserHandler)

	log.Println("Starting server on port 3000...")
	log.Println("POST to /register with JSON body to test validation.")
	log.Fatal(app.Listen(":3000"))
}
