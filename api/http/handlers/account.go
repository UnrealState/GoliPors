package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golipors/api/http/handlers/helpers"
	"golipors/api/http/services"
	"golipors/api/http/types"
	"golipors/app"
	"golipors/config"
	"net/http"
)

func RegisterAccountHandlers(router fiber.Router, appContainer app.App, cfg config.ServerConfig) {
	accountGroup := router.Group("/account")
	accountSvcGetter := services.AccountServiceGetter(appContainer, cfg)

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

		err := helpers.ParseRequestBody[*types.LoginRequest](c, &body)

		if err != nil {
			return err
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
					"error": "Internal server error",
					"msg":   err.Error(),
				})
			}
		}

		return c.JSON(response)
	}
}

func Register(svcGetter helpers.ServiceGetter[*services.AccountService]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return nil
	}
}

func VerifyOtp(svcGetter helpers.ServiceGetter[*services.AccountService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		/*
			var (
				authExp    = time.Now().Add(time.Minute * time.Duration(cfg.AuthExpirationMinutes))
				refreshExp = time.Now().Add(time.Minute * time.Duration(cfg.AuthRefreshMinutes))
			)

			user := &userDomain.User{}

			accessToken, err := jwt2.CreateToken([]byte(cfg.Secret), jwt2.GenerateUserClaims(user, authExp))
			refreshToken, err := jwt2.CreateToken([]byte(cfg.Secret), jwt2.GenerateUserClaims(user, refreshExp))

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Cannot create token",
				})
			}

			// Successful authentication
			return c.Status(http.StatusOK).JSON(&presenter.UserToken{
				AuthorizationToken: accessToken,
				RefreshToken:       refreshToken,
				ExpiresAt:          authExp.Unix(),
			})*/
		return nil
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
