package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golipors/app"
)

func RegisterAccountHandlers(router fiber.Router, app app.App) {
	accountGroup := router.Group("/account")

	accountGroup.Post("/login", Login)
	accountGroup.Get("/mmd", mmd)
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

func mmd(c *fiber.Ctx) error {
	return c.SendString("mmd")
}
