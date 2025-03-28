package routes

import (
	"dairy-management-backend/controllers"
	"dairy-management-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

// RegisterCowRoutes defines all routes related to cows
func RegisterCowRoutes(app *fiber.App, controller *controllers.CowController) {
	routes := app.Group("/api/admin/cows")

	routes.Post("/", middleware.AuthMiddleware("admin"), controller.CreateCow)      // Only admins can create a cow
	routes.Get("/", middleware.AuthMiddleware("admin"), controller.GetAllCows)      // Only admins can get all cows
	routes.Get("/:tag_id", controller.GetCowByID)                                   // Anyone can view cows by tag_id
	routes.Put("/:id", middleware.AuthMiddleware("admin"), controller.UpdateCow)    // Only admins can update cows
	routes.Delete("/:id", middleware.AuthMiddleware("admin"), controller.DeleteCow) // Only admins can delete cows
}
