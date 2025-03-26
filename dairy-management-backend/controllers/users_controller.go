package controllers

import (
	"dairy-management-backend/entities"
	"dairy-management-backend/usecases"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserUC *usecases.UserUseCase
}

func NewUserController(userUC *usecases.UserUseCase) *UserController {
	return &UserController{UserUC: userUC}
}

// Create User (Admin)
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := uc.UserUC.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

// Get All Users
func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := uc.UserUC.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// Get User By Email
func (uc *UserController) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	// Call the use case function, which now returns user and error
	user, err := uc.UserUC.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User retrieved successfully",
		"user":    user,
	})
}

// Update User
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id") // Get user ID from URL parameter
	var userUpdate entities.User

	if err := c.BodyParser(&userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Now passing (id, &userUpdate) correctly
	err := uc.UserUC.UpdateUser(id, &userUpdate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

// Delete User
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id") // Get the ID from the URL

	err := uc.UserUC.DeleteUser(id) // Pass only the string ID
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
