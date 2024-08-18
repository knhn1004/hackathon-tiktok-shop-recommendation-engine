package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.JWTMiddleware())
	SetupArticleRoutes(app)
	SetupShopRoutes(app)
	SetupProductRoutes(app)
	SetupInteractionRoutes(app)
	SetupRecommendationRoutes(app)
	// Add other route setup functions here as needed
}
