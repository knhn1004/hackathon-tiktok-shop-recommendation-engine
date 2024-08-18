// recommendation_routes.go
package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/handlers"
)

func SetupRecommendationRoutes(app *fiber.App) {
	app.Get("/api/articles/:articleId/recommendations", handlers.GetProductRecommendations)
}