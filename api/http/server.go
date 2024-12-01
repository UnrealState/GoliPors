// api/http/server.go
package http

import (
	"golipors/api/http/handlers"
	"golipors/api/http/middleware"
	"golipors/config"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg config.Config) {
	// Apply rate limiter middleware globally
	app.Use(middleware.RateLimiter())

	// Public routes
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/verify-otp", handlers.VerifyOTP)
	app.Post("/request-password-reset", handlers.RequestPasswordReset)
	app.Post("/reset-password", handlers.ResetPassword)

	// Protected routes
	authGroup := app.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware)
	authGroup.Post("/logout", handlers.Logout)
	// Add more protected routes here
}
