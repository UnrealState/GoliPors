package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golipors/api/http/handlers/helpers"
	middlerwares "golipors/api/http/middlewares"
	"golipors/api/http/services"
	"golipors/api/http/types"
	"golipors/app"
	"golipors/config"
	"net/http"
)

func RegisterAccountHandlers(router fiber.Router, appContainer app.App, cfg config.ServerConfig) {
	accountGroup := router.Group("/account")
	accountSvcGetter := services.AccountServiceGetter(appContainer, cfg)

	accountGroup.Use(middlerwares.SetTransaction(appContainer.DB()))

	accountGroup.Post("/login", Login(accountSvcGetter))
	accountGroup.Post("/register", Register(accountSvcGetter))
	accountGroup.Post("/verify-otp", VerifyOtp(accountSvcGetter))
	accountGroup.Post("/reset-password", ResetPassword(accountSvcGetter))
	accountGroup.Post("/reset-password/verify", ResetPasswordVerify(accountSvcGetter))
}

func Login(svcGetter helpers.ServiceGetter[*services.AccountService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		body := new(types.LoginRequest)

		if err := helpers.ParseRequestBody(c, body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		if !helpers.IsValidEmail(body.Email) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid email",
			})
		}

		response, err := svc.Login(c.UserContext(), *body)

		if err != nil {

			switch {
			case errors.Is(err, services.ErrUserNotFound):
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"error": "username or password incorrect",
				})
			default:
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Internal server error",
					"message": err.Error(),
				})
			}
		}

		return c.JSON(response)
	}
}

func Register(svcGetter helpers.ServiceGetter[*services.AccountService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		body := new(types.RegisterRequest)

		if err := helpers.ParseRequestBody(c, body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		if !helpers.IsValidEmail(body.Email) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid email",
			})
		}

		if len(body.Password) > 72 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "password too long",
			})
		}

		if _, err := helpers.IsValidDate(body.Birthday); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "invalid birthday",
				"message": err.Error(),
			})
		}

		if len(body.NationalID) != 10 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid nationalID",
			})
		}

		err := svc.Register(c.UserContext(), *body)

		if err != nil {
			switch {
			case errors.Is(err, services.ErrUserOnCreate):
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Error while creating new user",
					"message": err.Error(),
				})
			case errors.Is(err, services.ErrUserAlreadyExists):
				return c.Status(http.StatusConflict).JSON(fiber.Map{
					"error": "User already exists",
				})
			case errors.Is(err, services.ErrBirthdayInvalid):
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid birthday",
				})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Internal server error",
					"message": err.Error(),
				})
			}
		}

		return c.JSON(fiber.Map{
			"message": "User registered successfully",
		})
	}
}

func VerifyOtp(svcGetter helpers.ServiceGetter[*services.AccountService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		body := new(types.VerifyOTPRequest)

		if err := helpers.ParseRequestBody[*types.VerifyOTPRequest](c, &body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		if !helpers.IsValidEmail(body.Email) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid email",
			})
		}

		response, err := svc.VerifyOtp(c.UserContext(), *body)

		if err != nil {
			switch {
			case errors.Is(err, services.ErrUserNotFound):
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"error": "Incorrect code, session or email",
				})
			default:
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error": "Internal server error",
					"msg":   err.Error(),
				})
			}
		}

		return c.JSON(response)
	}
}

func ResetPassword(svcGetter helpers.ServiceGetter[*services.AccountService]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func ResetPasswordVerify(svcGetter helpers.ServiceGetter[*services.AccountService]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
