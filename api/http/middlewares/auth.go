package middlerwares

import (
	"github.com/gofiber/fiber/v2"
	"golipors/config"
	"golipors/pkg/jwt"
)

func Authorization(c *fiber.Ctx, cfg config.ServerConfig) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
	}

	clams, err := jwt.ParseToken(authHeader, []byte(cfg.Secret))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Wrong authorization header",
			"msg":   err.Error(),
		})
	}

	c.Locals("clams", clams)
	return c.Next()
}
