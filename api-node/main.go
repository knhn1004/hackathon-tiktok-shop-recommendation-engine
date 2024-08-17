package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/config"
)

func main() {
	// Load configuration
	err := config.Load()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	// Create a new Fiber app
	app := fiber.New()

	// Define a route for the root path
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString(config.OpenAIKey)
	})

	// Start the server
	log.Fatal(app.Listen(config.ServerAddr))
}
