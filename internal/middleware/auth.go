package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/zaidalmaghfur/go-app/pkg"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return pkg.Error(c, fiber.StatusUnauthorized, "Missing or invalid token", nil)
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		valid, err := pkg.ValidateToken(token)
		if err != nil || !valid {
			return pkg.Error(c, fiber.StatusUnauthorized, "Invalid token", nil)
		}

		return c.Next()
	}
}
