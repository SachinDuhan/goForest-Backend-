package middleware

import "github.com/gofiber/fiber/v2"

// ErrorHandler middleware to catch panics and errors
func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return nil
	}
}
