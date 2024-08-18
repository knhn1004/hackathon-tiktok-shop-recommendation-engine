package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func JWTMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// TODO: verify clerk JWT

		c.Locals("userID", "dummy-user-123")

		return c.Next()
	}
}
