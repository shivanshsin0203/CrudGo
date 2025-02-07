package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Printf("%s - %s %s", c.IP(), c.Method(), c.OriginalURL())
		return c.Next()
	}
}
