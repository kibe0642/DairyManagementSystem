// routes/milk_routes.go - API routes for milk collection
package routes

import (
	"dairy-management-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterMilkRoutes(app *fiber.App, controller *controllers.MilkController) {
	routes := app.Group("/milk")
	routes.Post("/", controller.CreateMilkRecord)
	routes.Get("/", controller.GetAllMilkRecords)
	routes.Get("/:id", controller.GetMilkRecordByID)
	routes.Put("/:id", controller.UpdateMilkRecord)
	routes.Delete("/:id", controller.DeleteMilkRecord)
}
