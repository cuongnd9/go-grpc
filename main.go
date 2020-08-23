package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	// Define constants
	const port = 3000

	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) {
		// nolint: gomnd
		c.SendStatus(404)
	})

	// Start server
	log.Fatal(app.Listen(port))
}

// Handler
func hello(c *fiber.Ctx) {
	c.Send("Hello, World 👋!")
}
