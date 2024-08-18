package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/config"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/routes"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
)

func main() {
	// Load configuration
	err := config.Load()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	// Initialize the database
	err = db.InitDB(
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	// Create a new Fiber app
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	app.Use(logger.New())

	// Setup routes
	routes.SetupRoutes(app)


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
