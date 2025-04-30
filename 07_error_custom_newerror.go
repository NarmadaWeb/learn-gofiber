package main

import (
	"errors"
	"log"
	"github.com/gofiber/fiber/v2"
)

var ErrInsufficientFunds = errors.New("dana tidak mencukupi")

var paymentGateway = struct {
	Charge func(amount float64) error
}{
	Charge: func(amount float64) error {
		log.Printf("Payment Gateway: Mencoba charge %.2f", amount)
		if amount <= 0 {
			return errors.New("jumlah harus positif")
		}
		if amount > 100 {
			return ErrInsufficientFunds
		}
		if amount == 50 {
			return errors.New("koneksi ke bank gagal")
		}
		log.Printf("Payment Gateway: Charge %.2f berhasil", amount)
		return nil
	},
}

var inputValid = true

func ProcessPayment(c *fiber.Ctx) error {
	amount, err := c.QueryFloat("amount", 0)
	if err != nil || amount == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Parameter 'amount' (float) dibutuhkan.")
	}

	log.Printf("Memproses pembayaran untuk jumlah: %.2f", amount)

	if !inputValid {
		log.Println("Input tidak valid terdeteksi")
		return fiber.NewError(fiber.StatusBadRequest, "Data pembayaran tidak lengkap atau tidak valid.")
	}

	err = paymentGateway.Charge(amount)
	if err != nil {
		log.Printf("Error dari payment gateway: %v", err)
		if errors.Is(err, ErrInsufficientFunds) {
			return fiber.NewError(fiber.StatusPaymentRequired, "Dana tidak mencukupi.")
		}
		log.Printf("Error payment gateway tidak dikenal: %v", err)
		return fiber.NewError(fiber.StatusServiceUnavailable, "Layanan pembayaran sedang tidak tersedia.")
	}

	return c.JSON(fiber.Map{"status": "Pembayaran berhasil", "amount": amount})
}

func main() {
	app := fiber.New()

	app.Post("/pay", ProcessPayment)

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
