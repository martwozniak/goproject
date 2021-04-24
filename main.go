package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Homepage
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	// Static resources like CSS, JS, IMG
	app.Static("/", "./public")

	// Serve HTML file
	app.Get("/default", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/default.html")
	})

	app.Listen(":3000")
}
