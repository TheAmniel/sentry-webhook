package route

import "github.com/gofiber/fiber/v2"

// TODO
func SentryRoute(c *fiber.Ctx) error {
	return c.Send([]byte("Hello world!"))
}
