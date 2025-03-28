package controllers

import (
	"dairy-management-backend/entities"
	"dairy-management-backend/usecases"

	"github.com/gofiber/fiber/v2"
)

// AuthController handles authentication and user creation
type AuthController struct {
	AuthUC *usecases.AuthUseCase
	UserUC *usecases.UserUseCase // ✅ Added UserUseCase for CreateUser
}

// NewAuthController initializes AuthController with AuthUseCase and UserUseCase
func NewAuthController(authUC *usecases.AuthUseCase, userUC *usecases.UserUseCase) *AuthController {
	return &AuthController{
		AuthUC: authUC,
		UserUC: userUC, // ✅ Ensure UserUseCase is properly injected
	}
}

// Login endpoint for users to authenticate and receive JWT token
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := ac.AuthUC.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// CreateUser allows only admins to create new users
func (ac *AuthController) CreateUser(c *fiber.Ctx) error {
	adminRole, ok := c.Locals("userRole").(string)
	if !ok || adminRole != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Only admins can create users"})
	}

	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// ✅ Use UserUseCase for creating users
	err := ac.UserUC.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}
func (ac *AuthController) GetCurrentUser(ctx *fiber.Ctx) error {
	// Extract user email from context (set by middleware)
	userEmail, ok := ctx.Locals("userEmail").(string) // Ensure it's a string
	if !ok || userEmail == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Fetch user details from the repository via UseCase
	user, err := ac.UserUC.GetUserByEmail(userEmail) // ✅ Fetch user by email
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Return user data
	return ctx.JSON(user)
}
