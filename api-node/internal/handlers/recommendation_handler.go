package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/recommendation"
)

func GetProductRecommendations(c fiber.Ctx) error {
    articleID, err := strconv.ParseUint(c.Params("articleId"), 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid article ID"})
    }

    userID := c.Locals("userId").(string)

    client, err := recommendation.NewClient("recommendation:50051")
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to connect to recommendation service"})
    }
    defer client.Close()

    productIDs, err := client.GetRecommendations(c.Context(), userID, articleID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get recommendations"})
    }

    // Load products from database
    var recommendedProducts []models.Product
    result := db.DB.Where("id IN ?", productIDs).Find(&recommendedProducts)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch recommended products"})
    }

    return c.JSON(recommendedProducts)
}