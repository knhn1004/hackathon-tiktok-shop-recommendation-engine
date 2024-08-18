package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
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
	userID := c.Locals("userID").(string)
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	articleIDUint, err := strconv.ParseUint(articleID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid article ID",
		})
	}

	like := models.ArticleLike{
		UserProfileID: uint(userIDUint),
		ArticleID:     uint(articleIDUint),
	}

	result := db.DB.Create(&like)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to like article",
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func UnlikeArticle(c fiber.Ctx) error {
	articleID := c.Params("id")
	userID := c.Locals("userID").(string)

	result := db.DB.Where("user_profile_id = ? AND article_id = ?", userID, articleID).
		Delete(&models.ArticleLike{})

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to unlike article",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func CreateComment(c fiber.Ctx) error {
	articleID := c.Params("id")
	userID := c.Locals("userID").(string)
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
