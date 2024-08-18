package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/handlers"
)

func SetupInteractionRoutes(app *fiber.App) {
	interactionGroup := app.Group("/api/interactions")

	interactionGroup.Post("/articles", handlers.CreateArticleInteraction)
	interactionGroup.Post("/products", handlers.CreateProductInteraction)
}