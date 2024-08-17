package main

import (
	"log"
	"strings"

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

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Ensure the server binds to all available network interfaces
	parts := strings.Split(config.ServerAddr, ":")
	port := parts[len(parts)-1]
	addr := "0.0.0.0:" + port

	// Start the server
	log.Printf("Server starting on %s", addr)
	log.Fatal(app.Listen(addr))
}
