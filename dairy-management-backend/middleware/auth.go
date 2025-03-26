package middleware

import (
	"dairy-management-backend/entities"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware validates JWT and checks user role
func AuthMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")

		// Check if token is present
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Missing token",
			})
		}

		// Ensure correct Bearer format
		tokenParts := strings.Split(tokenStr, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid token format",
			})
		}

		tokenString := tokenParts[1] // Extract actual token

		// Get JWT secret key
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			jwtSecret = "default-secret"
		}

		// Parse and validate token
		token, err := jwt.ParseWithClaims(tokenString, &entities.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid token",
			})
		}

		// Extract claims and check role
		claims, ok := token.Claims.(*entities.JWTClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid token claims",
			})
		}

		fmt.Println("User Role:", claims.Role) // Debugging

		// Store user role in context
		c.Locals("userRole", claims.Role)

		// Check required role
		if requiredRole != "" && claims.Role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden: Insufficient permissions",
			})
		}

		// Proceed if authorized
		return c.Next()
	}
}
