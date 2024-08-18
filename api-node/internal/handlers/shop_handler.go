package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/kafka"
)

func GetProductsByShopID(c fiber.Ctx) error {
    shopID, err := strconv.Atoi(c.Params("shopId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid shop ID"})
    }

    var products []models.Product
    result := db.DB.Table("products").
        Select("products.*, shops.*, creators.*, user_profiles.*, categories.*").
        Joins("LEFT JOIN shops ON shops.id = products.shop_id").
        Joins("LEFT JOIN creators ON creators.id = shops.creator_id").
        Joins("LEFT JOIN user_profiles ON user_profiles.id = creators.user_profile_id").
        Joins("LEFT JOIN categories ON categories.id = products.category_id").
        Where("products.shop_id = ?", shopID).
        Find(&products)

    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
    }

    // Debug: Print product details
    for _, product := range products {
        fmt.Printf("Product ID: %d, Shop ID: %d, Creator ID: %d, UserProfile ID: %d\n",
            product.ID, product.Shop.ID, product.Shop.Creator.ID, product.Shop.Creator.UserProfile.ID)
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
    userID, ok := c.Locals("userId").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
    }
    shopID, err := strconv.Atoi(c.Params("shopId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid shop ID"})
    }
    var orderItems []models.OrderItemInput
    if err := c.Bind().JSON(&orderItems); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
    }

    // Create a new order
    order := models.Order{
        UserID: userID,
        ShopID: uint(shopID),
        Status: "pending",
    }

    // Start a transaction
    tx := db.DB.Begin()
    if tx.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to start transaction"})
    }

    // Create the order
    if err := tx.Create(&order).Error; err != nil {
        tx.Rollback()
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
    }

    // Create order items and product interactions
    var interactions []models.UserProductInteraction
    for _, item := range orderItems {
        var product models.Product
        if err := tx.First(&product, item.ProductID).Error; err != nil {
            tx.Rollback()
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Product with ID %d not found", item.ProductID)})
        }

        orderItem := models.OrderItem{
            OrderID:   order.ID,
            ProductID: item.ProductID,
            Quantity:  item.Quantity,
            Price:     product.Price,
        }

        if err := tx.Create(&orderItem).Error; err != nil {
            tx.Rollback()
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order item"})
        }

        // use the getUserProfileByUserID function to get the user profile
        userProfile, err := models.GetUserProfileByUserID(db.DB, order.UserID)
        if err != nil {
            tx.Rollback()
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user profile"})
        }

        interaction := models.UserProductInteraction{
            UserProfileID:   userProfile.ID,
            ProductID:       item.ProductID,
            InteractionType: "purchase",
        }

        if err := tx.Create(&interaction).Error; err != nil {
            tx.Rollback()
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product interaction"})
        }

        interactions = append(interactions, interaction)
    }

    // Commit the transaction
    if err := tx.Commit().Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
    }

    // Send interactions to Kafka asynchronously
    go func(interactions []models.UserProductInteraction, userID string) {
        for _, interaction := range interactions {
            kafka.WriteProductInteraction(context.Background(), userID, "purchase", interaction)
        }
    }(interactions, userID)

    fmt.Printf("Order placed - userID: %s, shopID: %d, orderID: %d\n", userID, shopID, order.ID)
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Order placed successfully",
        "orderID": order.ID,
    })
}