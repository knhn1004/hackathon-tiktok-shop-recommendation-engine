package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/user"
	"gorm.io/gorm"
)


func CreateUserProfile(c fiber.Ctx) error {
    var newUser models.UserProfile
    if err := c.Bind().JSON(&newUser); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
    }

    // Check if the newUser struct is empty
    if newUser == (models.UserProfile{}) {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
    }

    err := user.CreateUser(&newUser)
    if err != nil {
        switch {
        case errors.Is(err, gorm.ErrDuplicatedKey):
            return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
        default:
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
        }
    }

    return c.Status(fiber.StatusCreated).JSON(newUser)
}


