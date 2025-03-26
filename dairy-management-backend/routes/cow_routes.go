// routes/cow_routes.go - API route definitions for cows
package routes

import (
	"dairy-management-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// RegisterCowRoutes defines all routes related to cows
func RegisterCowRoutes(app *fiber.App, controller *controllers.CowController) {
	routes := app.Group("/api/admin/cows") // Ensure it starts with `/api/admin`

	routes.Post("/", controller.CreateCow)
	routes.Get("/", controller.GetAllCows)
	routes.Get("/:tag_id", controller.GetCowByID) // Now uses tag_id instead of id
	routes.Put("/:id", controller.UpdateCow)
	routes.Delete("/:id", controller.DeleteCow)
}
