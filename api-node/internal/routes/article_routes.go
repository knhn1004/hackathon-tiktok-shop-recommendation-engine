package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/handlers"
)

func SetupArticleRoutes(app *fiber.App) {
	articleGroup := app.Group("/api/articles")

	articleGroup.Get("/", handlers.GetArticles)
	articleGroup.Get("/:id/comments", handlers.GetArticleComments)
	articleGroup.Post("/:id/like", handlers.LikeArticle)
	articleGroup.Delete("/:id/like", handlers.UnlikeArticle)
	articleGroup.Post("/:id/comments", handlers.CreateComment)
}