package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/theamniel/sentry-webhook/middleware"
	"github.com/theamniel/sentry-webhook/route"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(200)
	})
	// ik, is POST :p
	app.Get("/sentry", middleware.VerifySentrySignature, route.SentryRoute)

	log.Fatal(app.Listen(":1809"))
}
