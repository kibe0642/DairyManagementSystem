package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)
	log.Printf("[%s]%s%s | %dms", c.Method(), c.Path(), c.IP(), duration.Milliseconds())
	return err
}
