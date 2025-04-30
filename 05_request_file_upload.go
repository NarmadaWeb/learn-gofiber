package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadFileHandler(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file_upload")
	if err != nil {
		log.Printf("Error mendapatkan file: %v", err)
		if err.Error() == "there is no uploaded file associated with the given key" || err.Error() == "multipart: no such file" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Field 'file_upload' tidak ditemukan atau kosong"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses upload file: " + err.Error()})
	}

	description := c.FormValue("description", "Tidak ada deskripsi")

	log.Printf("Menerima file: %s, Size: %d, Description: %s", fileHeader.Filename, fileHeader.Size, description)

	maxSize := int64(5 * 1024 * 1024)
	if fileHeader.Size > maxSize {
		log.Printf("Ukuran file %d melebihi batas %d", fileHeader.Size, maxSize)
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{"error": "Ukuran file melebihi batas 5MB"})
	}

	allowedMIMETypes := map[string]bool{
		"image/jpeg":      true,
		"image/png":       true,
		"application/pdf": true,
		"text/plain":      true,
	}
	file, err := fileHeader.Open()
	if err != nil {
		log.Printf("Error membuka file header: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuka file"})
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		log.Printf("Error membaca buffer file: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membaca file"})
	}
	mimeType := http.DetectContentType(buffer)
	log.Printf("Detected MIME type: %s", mimeType)

	if !allowedMIMETypes[mimeType] {
		log.Printf("Tipe MIME '%s' tidak diizinkan", mimeType)
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"error":         "Tipe file tidak didukung",
			"detected_mime": mimeType,
			"allowed_mimes": func() []string {
				keys := make([]string, 0, len(allowedMIMETypes))
				for k := range allowedMIMETypes {
					keys = append(keys, k)
				}
				return keys
			}(),
		})
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Printf("Error reset file pointer: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal reset file pointer"})
	}

	safeFilenamePart := filepath.Base(fileHeader.Filename)
	uniqueFilename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), safeFilenamePart)
	uploadDir := "./uploads"
	savePath := filepath.Join(uploadDir, uniqueFilename)

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Printf("Error membuat direktori uploads: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyiapkan penyimpanan file"})
	}

	dst, err := os.Create(savePath)
	if err != nil {
		log.Printf("Error membuat file tujuan %s: %v", savePath, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuat file tujuan"})
	}
	defer dst.Close()

	bytesCopied, err := io.Copy(dst, file)
	if err != nil {
		log.Printf("Error menyalin file ke %s: %v", savePath, err)
		os.Remove(savePath)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyalin file"})
	}

	log.Printf("File '%s' (deskripsi: '%s', size: %d, mime: %s) berhasil diupload ke %s (%d bytes copied)",
		fileHeader.Filename, description, fileHeader.Size, mimeType, savePath, bytesCopied)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":       "File berhasil diupload!",
		"original_name": fileHeader.Filename,
		"saved_path":    savePath,
		"saved_filename": uniqueFilename,
		"size":          fileHeader.Size,
		"mime_type":     mimeType,
		"description":   description,
	})
}

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Post("/upload", UploadFileHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(`
			<h2>Upload File</h2>
			<form action="/upload" method="post" enctype="multipart/form-data">
				<label for="file_upload">Pilih file:</label><br>
				<input type="file" id="file_upload" name="file_upload"><br><br>
				<label for="description">Deskripsi:</label><br>
				<input type="text" id="description" name="description"><br><br>
				<input type="submit" value="Upload">
			</form>
		`)
	})

	log.Println("Starting server on port 3000...")
	log.Println("Upload directory: ./uploads")
	log.Fatal(app.Listen(":3000"))
}
