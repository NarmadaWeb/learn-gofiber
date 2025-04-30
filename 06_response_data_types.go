package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"encoding/xml"
	"os"
	"time"
)

func userHasPermission(c *fiber.Ctx) bool {
	apiKey := c.Get("X-API-Key")
	return apiKey == "valid-key"
}

type XmlResponse struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Message string   `xml:"message"`
	Code    int      `xml:"code"`
}

type TemplateData struct {
	Title string
	Time  time.Time
}


func main() {
	app := fiber.New(fiber.Config{})

	dummyFilePath := "./data_file.txt"
	_ = os.WriteFile(dummyFilePath, []byte("Ini adalah konten dari file data."), 0644)
	defer os.Remove(dummyFilePath)

	app.Get("/resource", func(c *fiber.Ctx) error {
		log.Println("Checking permission for /resource")
		if !userHasPermission(c) {
			log.Println("Permission denied for /resource")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"code":    "ACCESS_DENIED",
				"message": "Anda tidak memiliki izin untuk mengakses sumber daya ini.",
			})
		}
		log.Println("Permission granted for /resource")
		return c.JSON(fiber.Map{"data": "Ini data rahasia Anda!", "status": "success"})
	})

	app.Get("/string", func(c *fiber.Ctx) error {
		return c.SendString("Ini adalah response string biasa.")
	})

	app.Get("/bytes", func(c *fiber.Ctx) error {
		byteData := []byte{72, 101, 108, 108, 111}
		return c.Send(byteData)
	})

	app.Get("/xml", func(c *fiber.Ctx) error {
		xmlData := XmlResponse{
			Status:  "success",
			Message: "Ini response XML",
			Code:    200,
		}
		log.Println("Sending XML response")
		xmlBytes, err := xml.MarshalIndent(xmlData, "", "  ")
		if err != nil {
			return fiber.ErrInternalServerError
		}
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationXMLCharsetUTF8)
		return c.Send(xmlBytes)
	})

	app.Get("/html-template", func(c *fiber.Ctx) error {
		log.Println("Rendering HTML template (memerlukan Views engine)")
		data := TemplateData{
			Title: "Halaman Info",
			Time:  time.Now(),
		}
		return c.Status(fiber.StatusNotImplemented).SendString("Template engine tidak di-setup untuk contoh ini")
	})

	app.Get("/file-inline", func(c *fiber.Ctx) error {
		log.Println("Sending file inline:", dummyFilePath)
		return c.SendFile(dummyFilePath)
	})

	app.Get("/file-download", func(c *fiber.Ctx) error {
		log.Println("Sending file for download:", dummyFilePath)
		return c.Download(dummyFilePath, "laporan_penting.txt")
	})

	app.Get("/goto-google", func(c *fiber.Ctx) error {
		log.Println("Redirecting to Google")
		return c.Redirect("https://google.com", fiber.StatusTemporaryRedirect)
	})

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
