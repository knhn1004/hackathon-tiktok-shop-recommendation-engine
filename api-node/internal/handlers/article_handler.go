package handlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/kafka"
)

func GetArticles(c fiber.Ctx) error {
	page := c.Query("page", "1")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	limit := 10
	offset := (pageInt - 1) * limit

	var articles []models.Article
	result := db.DB.Preload("Creator").Preload("Likes").
		Offset(offset).Limit(limit).Find(&articles)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch articles",
		})
	}

	return c.JSON(articles)
}

func GetArticleComments(c fiber.Ctx) error {
	articleID := c.Params("id")

	var comments []models.Comment
	result := db.DB.Where("article_id = ?", articleID).
		Preload("UserProfile").Find(&comments)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch comments",
		})
	}

	return c.JSON(comments)
}

func LikeArticle(c fiber.Ctx) error {
	articleID := c.Params("id")
	userID := c.Locals("userId").(string)

	articleIDUint, err := strconv.ParseUint(articleID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid article ID",
		})
	}

	userProfile, err := models.GetUserProfileByUserID(db.DB, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user profile",
		})
	}

	like := models.ArticleLike{
		UserProfileID: userProfile.ID,
		ArticleID:     uint(articleIDUint),
	}

	result := db.DB.FirstOrCreate(&like, like)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to like article",
		})
	}

	// Send interaction to Kafka asynchronously only if a new like was created
	if result.RowsAffected > 0 {
		go func() {
			kafka.WriteArticleInteraction(context.Background(), userID, "like", map[string]interface{}{
				"articleId": articleIDUint,
			})
		}()
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Article liked successfully",
	})
}

func UnlikeArticle(c fiber.Ctx) error {
	articleID := c.Params("id")
	userID := c.Locals("userId").(string)

	userProfile, err := models.GetUserProfileByUserID(db.DB, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get user profile",
			"userId": userID,
		})
	}

	result := db.DB.Where("user_profile_id = ? AND article_id = ?", userProfile.ID, articleID).
		Delete(&models.ArticleLike{})

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func CreateComment(c fiber.Ctx) error {
	articleID := c.Params("id")
	userID := c.Locals("userId").(string)
	var comment models.Comment
	if err := c.Bind().JSON(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	// Check if the comment struct is empty
	if comment.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	articleIDUint, err := strconv.ParseUint(articleID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid article ID",
		})
	}

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	comment.ArticleID = uint(articleIDUint)
	comment.UserProfileID = uint(userIDUint)

	result := db.DB.Create(&comment)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create comment",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}