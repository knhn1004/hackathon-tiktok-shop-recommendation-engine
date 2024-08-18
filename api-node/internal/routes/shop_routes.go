package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/handlers"
)


func SetupShopRoutes(app *fiber.App) {
	shopGroup := app.Group("/api/shops")

	shopGroup.Get("/:shopId/products", handlers.GetProductsByShopID)
	shopGroup.Post("/:shopId/order", handlers.OrderCart)
}

func SetupProductRoutes(app *fiber.App) {
	productGroup := app.Group("/api/products")

	productGroup.Get("/:productId", handlers.GetProductDetail)
}