package middleware

import "github.com/gofiber/fiber/v2"

func VerifySentrySignature(c *fiber.Ctx) error {
	return c.Next()
}
