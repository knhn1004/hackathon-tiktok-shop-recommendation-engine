package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
)
func GetProductsByShopID(c fiber.Ctx) error {
	shopID, err := strconv.Atoi(c.Params("shopId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid shop ID"})
	}

	var products []models.Product
	result := db.DB.Where("shop_id = ?", shopID).Find(&products)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}

	return c.JSON(products)
}


func GetProductDetail(c fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("productId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	var product models.Product
	result := db.DB.First(&product, productID)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch product"})
	}

	return c.JSON(product)
}


func OrderCart(c fiber.Ctx) error {
	userID := c.Locals("userId").(string)
	shopID, err := strconv.Atoi(c.Params("shopId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid shop ID"})
	}


	fmt.Printf("userID: %s, shopID: %d\n", userID, shopID)
	// just a placeholder that returns a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Order placed successfully"})
}