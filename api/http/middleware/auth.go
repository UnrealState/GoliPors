package middleware

import (
    "golipors/pkg/utils"
    "strings"

    "github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
    }

    tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
    claims, err := utils.ValidateToken(tokenString)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
    }

    // Store user ID in Locals for access in handlers
    c.Locals("userID", claims.UserID)
    return c.Next()
}
