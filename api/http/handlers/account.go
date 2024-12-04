package handlers

import (
	"github.com/gofiber/fiber/v2"
	app "golipors/app"
	"golipors/config"
)

func RegisterAccountHandlers(router fiber.Router, appContainer app.App, cfg config.ServerConfig) {
	accountGroup := router.Group("/account")

	accountGroup.Post("/login", Login)
	accountGroup.Get("/logout", Logout)
	accountGroup.Post("/register", Register)
	accountGroup.Post("/verify-otp", VerifyOtp)
	accountGroup.Post("/reset-password", ResetPassword)
	accountGroup.Post("/reset-password/verify", ResetPasswordVerify)
}

func Login(c *fiber.Ctx) error {
	return c.SendString("")
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

func Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

func VerifyOtp(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

func ResetPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

func ResetPasswordVerify(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
