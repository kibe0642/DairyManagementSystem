package routes

import (
	"dairy-management-backend/controllers"
	"dairy-management-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

// Register authentication routes
func RegisterAuthRoutes(app *fiber.App, authController *controllers.AuthController) {
	routes := app.Group("/api/auth")
	routes.Post("/login", authController.Login)
	routes.Post("/create-user", middleware.AuthMiddleware("admin"), authController.CreateUser) // ✅ Fixed: Use authController instead of controller
}
