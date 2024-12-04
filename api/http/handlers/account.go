package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golipors/api/http/handlers/helpers"
	"golipors/api/http/handlers/presenter"
	"golipors/app"
	"golipors/config"
	userPort "golipors/internal/user"
	userDomain "golipors/internal/user/domain"
	"golipors/pkg/adapters/email"
	"golipors/pkg/cache"
	jwt2 "golipors/pkg/jwt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	ErrUserCreationValidation = userPort.ErrUserCreationValidation
	ErrUserOnCreate           = userPort.ErrUserOnCreate
	ErrUserNotFound           = userPort.ErrUserNotFound
	ErrInvalidPassword        = userPort.ErrInvalidPassword
)

func RegisterAccountHandlers(router fiber.Router, appContainer app.App, cfg config.ServerConfig) {
	accountGroup := router.Group("/account")

	accountGroup.Post("/login", Login(appContainer, cfg))
	accountGroup.Post("/register", Register)
	accountGroup.Post("/verify-otp", VerifyOtp(appContainer, cfg))
	accountGroup.Post("/reset-password", ResetPassword)
	accountGroup.Post("/reset-password/verify", ResetPasswordVerify)
}

func Login(appContainer app.App, cfg config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userService := appContainer.UserService()
		body := new(presenter.LoginInput)

		if err := c.BodyParser(body); err != nil || body.Email == "" || body.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Missing required body data",
			})
		}

		if !helpers.IsValidEmail(body.Email) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid email",
			})
		}

		user, err := userService.GetUserByUsernamePassword(c.UserContext(), body.Email, body.Password)

		if err != nil {
			switch {
			case errors.Is(err, ErrUserNotFound):
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"error": "User not found",
				})
			default:
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
		}

		code, err := helpers.GenerateOTP()

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error generating OTP",
			})
		}

		log.Println("OTP sent for user", user.ID, "code:", code)

		oc := cache.NewJsonObjectCache[presenter.LoginCacheSession](appContainer.Cache(), "auth_session_")

		emailService := email.NewEmailAdapter(appContainer.Config().SMTP)

		err = emailService.SendText(
			body.Email,
			fmt.Sprintf("GoliPors OTP code for %s", body.Email),
			fmt.Sprintf("GoliPors OTP code: %s", code),
		)

		log.Println(err)

		err = nil

		err = oc.Set(c.UserContext(), strconv.Itoa(int(user.ID)), time.Minute*time.Duration(cfg.OtpTtlMinutes), presenter.LoginCacheSession{
			SessionID: uuid.New(),
			UserID:    user.ID,
			Code:      code,
		})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
				"error":   err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"code":       code,
			"session_id": uuid.New(),
		})
	}
}

func Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

func VerifyOtp(appContainer app.App, cfg config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {

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
		})
	}
}

func ResetPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

func ResetPasswordVerify(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
