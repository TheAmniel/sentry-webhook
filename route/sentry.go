package route

import "github.com/gofiber/fiber/v2"

// TODO
// - Sentry-Hook-Resource: 'error'
// - Sentry-Hook-Resource: 'issue'?
func SentryRoute(c *fiber.Ctx) error {
	return c.Send([]byte("Hello world!"))
}
