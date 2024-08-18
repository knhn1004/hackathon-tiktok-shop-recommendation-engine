package middleware

import (
	"encoding/pem"
	"fmt"
	"strings"

	"crypto/x509"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/config"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/models"
	"github.com/knhn1004/hackathon-tiktok-shop-recommendation-engine/api-node/internal/services/db"
)

func JWTMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid Authorization header"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		 publicKeyPEM := []byte(config.ClerkPublicKey)

		block, _ := pem.Decode(publicKeyPEM)
		if block == nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse PEM block containing the public key"})
		}

		publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse public key"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId, ok := claims["sub"].(string)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid sub claim"})
			}
			// TODO: also get username from claims after clerk updates

			fmt.Printf("user logged in: %s\n", userId)

			var user models.UserProfile
			result := db.DB.Where("user_id = ?", userId).First(&user)
			if result.Error != nil {
				// Create new user if not exists
				user = models.UserProfile{
					UserID: userId,
					// Set default values for other fields if needed
				}
				result = db.DB.Create(&user)
				if result.Error != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
				}
			}

			c.Locals("user", user)
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}
}