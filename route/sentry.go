package route

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/theamniel/sentry-webhook/sentry"
)

// TODO
// - Sentry-Hook-Resource: 'error'
// - Sentry-Hook-Resource: 'event_alert'
// - Sentry-Hook-Resource: 'issue'?
func SentryRoute(c *fiber.Ctx) error {
	var body sentry.Sentry
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(400)
	}

	resource := c.Get("Sentry-Hook-Resource")
	if body.Action == "" || body.Data == nil || resource == "" {
		return c.SendStatus(400)
	}
	log.Printf("Received '%s.%s' webhook from Sentry.\n", resource, body.Action)

	if resource == "error" {
		return c.SendStatus(200)
	}

	switch resource {
	case "error":
		return c.SendStatus(200)
	case "event_alert":
		return c.SendStatus(200)
	case "issue":
		return c.SendStatus(200)
	default:
		log.Printf("Unknown resource '%s.%s'", resource, body.Action)
		break
	}
	return c.Send([]byte("Hello world!"))
}
