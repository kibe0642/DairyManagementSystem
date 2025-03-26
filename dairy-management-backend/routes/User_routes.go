package routes

import (
	"dairy-management-backend/controllers"
	"dairy-management-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

// RegisterUserRoutes defines all user-related API endpoints.
func RegisterUserRoutes(app *fiber.App, userController *controllers.UserController) {
	api := app.Group("/api/admin") // All routes start with `/api/admin`

	// Protect routes with admin middleware
	api.Post("/users", middleware.AuthMiddleware("admin"), userController.CreateUser)
	api.Get("/users", middleware.AuthMiddleware("admin"), userController.GetAllUsers)
	api.Get("/users/:id", middleware.AuthMiddleware("admin"), userController.GetUserByEmail)
	api.Put("/users/:id", middleware.AuthMiddleware("admin"), userController.UpdateUser)
	api.Delete("/users/:id", middleware.AuthMiddleware("admin"), userController.DeleteUser)
}
