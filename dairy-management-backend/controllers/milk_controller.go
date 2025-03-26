package controllers
import (
	"dairy-management-backend/entities"
	"dairy-management-backend/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MilkController struct {
	MilkUC *usecases.MilkUseCase
}

func NewMilkController(usecase *usecases.MilkUseCase) *MilkController {
	return &MilkController{MilkUC: usecase}
}

func (mc *MilkController) CreateMilkRecord(c *fiber.Ctx) error {
	var milk entities.MilkCollection
	if err := c.BodyParser(&milk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := mc.MilkUC.AddMilkRecord(&milk); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(milk)
}

func (mc *MilkController) GetAllMilkRecords(c *fiber.Ctx) error {
	records, err := mc.MilkUC.GetMilkRecords()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(records)
}

func (mc *MilkController) GetMilkRecordByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	record, err := mc.MilkUC.GetMilkRecordByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Milk record not found"})
	}
	return c.JSON(record)
}

func (mc *MilkController) UpdateMilkRecord(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var milk entities.MilkCollection
	if err := c.BodyParser(&milk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	milk.ID = uint(id)
	if err := mc.MilkUC.UpdateMilkRecord(&milk); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(milk)
}

func (mc *MilkController) DeleteMilkRecord(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := mc.MilkUC.RemoveMilkRecord(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Milk record deleted successfully"})
}
