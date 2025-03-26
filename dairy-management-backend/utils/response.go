package utils

import "github.com/gofiber/fiber/v2"

type StandardResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccesResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(StandardResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}
func ErrorResponse(c *fiber.Ctx, message string) error {
	return c.JSON(StandardResponse{
		Success: false,
		Message: message,
	})
}
