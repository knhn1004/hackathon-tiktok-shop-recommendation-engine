package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/kafka"
	"gorm.io/gorm"
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

	var interaction models.UserArticleInteraction
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		var userProfile models.UserProfile
		if err := tx.Where("user_id = ?", userID).First(&userProfile).Error; err != nil {
			return err
		}

		interaction = models.UserArticleInteraction{
			UserProfileID:   userProfile.ID,
			ArticleID:       input.ArticleID,
			InteractionType: input.InteractionType,
			Duration:        input.Duration,
		}

		return tx.Create(&interaction).Error
	})

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User profile not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create interaction"})
	}

	// Send interaction to Kafka asynchronously
	go func() {
		kafka.WriteArticleInteraction(context.Background(), userID, interaction.InteractionType, interaction)
	}()

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

	var interaction models.UserProductInteraction
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		var userProfile models.UserProfile
		if err := tx.Where("user_id = ?", userID).First(&userProfile).Error; err != nil {
			return err
		}

		interaction = models.UserProductInteraction{
			UserProfileID:   userProfile.ID,
			ProductID:       input.ProductID,
			InteractionType: input.InteractionType,
		}

		return tx.Create(&interaction).Error
	})

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User profile not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create interaction"})
	}

	// Send interaction to Kafka asynchronously
	go func() {
		kafka.WriteProductInteraction(context.Background(), userID, interaction.InteractionType, interaction)
	}()

	return c.Status(fiber.StatusCreated).JSON(interaction)
}