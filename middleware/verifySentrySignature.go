package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

func VerifySignature(secret string) fiber.Handler {
	mac := hmac.New(sha256.New, []byte(secret))
	return func(c *fiber.Ctx) error {
		if _, err := mac.Write(c.Body()); err != nil {
			return err
		}

		digested := hex.EncodeToString(mac.Sum(nil))
		if digested != c.Get("Sentry-Hook-Signature") {
			return c.SendStatus(401)
		}

		return c.Next()
	}
}
