package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		
		err := c.Next()
		
		log.Printf(
			"%s %s %s %v",
			c.Method(),
			c.Path(),
			c.IP(),
			time.Since(start),
		)
		
		return err
	}
} 