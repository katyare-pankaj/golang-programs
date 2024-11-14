package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Example handler that may panic
	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("This handler will panic!")
	})

	app.Listen(":3000")
	fmt.Println("Server running on http://localhost:3000")
}
