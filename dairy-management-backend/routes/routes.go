package routes

import (
	"dairy-management-backend/controllers"
	"dairy-management-backend/usecases"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes initializes all routes
func RegisterRoutes(app *fiber.App, cowUC *usecases.CowUseCase) {
	// Initialize controllers with required dependencies
	cowController := controllers.NewCowController(cowUC)

	// Register cow routes
	RegisterCowRoutes(app, cowController)
}
