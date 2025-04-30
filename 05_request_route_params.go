package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

type UserRouteParams struct {
	UserID int `params:"userId"`
}

type OrderRouteParams struct {
	OrderID int `params:"orderId"`
}

type UserOrderRouteParams struct {
	UserID  int `params:"userId"`
	OrderID int `params:"orderId"`
}

func GetUserOrder(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
	orderIdStr := c.Params("orderId")
	log.Printf("Manual - User ID: %s, Order ID: %s", userIdStr, orderIdStr)

	userIdInt, errUser := c.ParamsInt("userId")
	orderIdInt, errOrder := c.ParamsInt("orderId")
	if errUser != nil || errOrder != nil {
		log.Printf("Manual Int Parse Error - User: %v, Order: %v", errUser, errOrder)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID atau Order ID tidak valid (harus angka)"})
	}
	log.Printf("Manual Int - User ID: %d, Order ID: %d", userIdInt, orderIdInt)

	var params UserOrderRouteParams
	if err := c.ParamsParser(&params); err != nil {
		log.Printf("Struct Parsing Error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Gagal parse route parameters: " + err.Error()})
	}
	log.Printf("Struct Binding - User ID: %d, Order ID: %d", params.UserID, params.OrderID)

	return c.JSON(fiber.Map{"user_id": params.UserID, "order_id": params.OrderID, "status": "found"})
}

func main() {
	app := fiber.New()

	app.Get("/users/:userId/orders/:orderId", GetUserOrder)

	log.Println("Starting server on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
