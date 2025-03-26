// controllers/cow_controller.go - API handlers for cows
package controllers

import (
	"dairy-management-backend/entities"
	"dairy-management-backend/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CowController struct {
	CowUC *usecases.CowUseCase
}

func NewCowController(usecase *usecases.CowUseCase) *CowController {
	return &CowController{CowUC: usecase}
}

func (cc *CowController) CreateCow(c *fiber.Ctx) error {
	var cow entities.Cow

	// Parse request body
	if err := c.BodyParser(&cow); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// 🛠 Ensure ID is NOT manually assigned
	cow.ID = 0 // Auto-increment will handle this

	// 🛠 Correct Age Conversion
	if cow.Age > 0 && cow.Age < 1 {
		// If Age is given as a fraction of a year, convert to months
		cow.AgeUnit = "months"
		cow.Age = cow.Age * 12 // Convert fractional years to months
	} else {
		cow.AgeUnit = "years"
	}

	// 🛠 Call UseCase to add the cow
	if err := cc.CowUC.AddCow(&cow); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 🛠 Return the created cow as response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Cow created successfully",
		"cow":     cow,
	})
}

func (cc *CowController) GetAllCows(c *fiber.Ctx) error {
	cows, err := cc.CowUC.GetCows()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cows)
}

func (cc *CowController) GetCowByID(c *fiber.Ctx) error {
	// Get tag_id from URL parameters
	tagID := c.Params("tag_id")

	// Fetch cow by tag_id
	cow, err := cc.CowUC.GetCowByTagID(tagID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cow not found",
		})
	}

	// Return cow data
	return c.Status(fiber.StatusOK).JSON(cow)
}

func (cc *CowController) UpdateCow(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid cow ID"})
	}
	var cow entities.Cow
	if err := c.BodyParser(&cow); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	cow.ID = uint(id)
	if err := cc.CowUC.UpdateCow(&cow); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cow)
}

func (cc *CowController) DeleteCow(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid cow ID"})
	}
	if err := cc.CowUC.RemoveCow(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Cow deleted successfully"})
}
