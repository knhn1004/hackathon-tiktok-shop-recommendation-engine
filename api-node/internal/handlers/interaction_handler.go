package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
)

func CreateArticleInteraction(c fiber.Ctx) error {
	var input struct {
		ArticleID       uint   `json:"articleId"`
		InteractionType string `json:"interactionType"`
		Duration        int    `json:"duration,omitempty"`
	}
	if err := c.Bind().JSON(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
	}

	userID := c.Locals("userId").(string)

	var userProfile models.UserProfile
	if err := db.DB.Where("user_id = ?", userID).First(&userProfile).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User profile not found"})
	}

	interaction := models.UserArticleInteraction{
		UserProfileID:   userProfile.ID,
		ArticleID:        input.ArticleID,
		InteractionType:  input.InteractionType,
		Duration:         input.Duration,
	}

	result := db.DB.Create(&interaction)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create interaction"})
	}

	return c.Status(fiber.StatusCreated).JSON(interaction)
}

func CreateProductInteraction(c fiber.Ctx) error {
	var input struct {
		ProductID       uint   `json:"productId"`
		InteractionType string `json:"interactionType"`
	}
	if err := c.Bind().JSON(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
	}

	userID := c.Locals("userId").(string)

	var userProfile models.UserProfile
	if err := db.DB.Where("user_id = ?", userID).First(&userProfile).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User profile not found"})
	}

	interaction := models.UserProductInteraction{
		UserProfileID:   userProfile.ID,
		ProductID:        input.ProductID,
		InteractionType:   input.InteractionType,
	}

	result := db.DB.Create(&interaction)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create interaction"})
	}

	return c.Status(fiber.StatusCreated).JSON(interaction)
}