package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

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
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/login.html")
	})

	// API
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		if c.Params("*") == "test" {
			return c.SendString("Test case api endpoint")
		}
		return c.SendString(msg) // => ✋ register
	})

	app.Post("/api/login", func(c *fiber.Ctx) error {
		return c.Send(c.Body())
	})

	app.Listen(":3000")
}
