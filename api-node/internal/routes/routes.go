package routes

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	SetupUserRoutes(app)
	// Add other route setup functions here as needed
}
