package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/handlers"
)

func SetupUserRoutes(app *fiber.App) {
	userGroup := app.Group("/api/users")

	userGroup.Post("/", handlers.CreateUserProfile)
	// Add other user-related routes here
}
