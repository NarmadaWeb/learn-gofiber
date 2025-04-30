package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error { return c.SendStatus(200) }).Name("get.users")
	app.Post("/users", func(c *fiber.Ctx) error { return c.SendStatus(201) })
	api := app.Group("/api")
	api.Get("/status", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) }).Name("api.status")


	app.Get("/debug/routes", func(c *fiber.Ctx) error {
		includeInternal := c.QueryBool("internal", false)

		routes := app.GetRoutes(includeInternal)

		var routeList []fiber.Map
		for _, route := range routes {
			routeList = append(routeList, fiber.Map{
				"method": route.Method,
				"name":   route.Name,
				"path":   route.Path,
				"params": route.Params,
			})
		}
		return c.JSON(routeList)
	})

	log.Println("Starting server on port 3000...")
	log.Println("Access /debug/routes to see registered routes.")
	log.Println("Access /debug/routes?internal=true to include internal routes.")
	log.Fatal(app.Listen(":3000"))
}
