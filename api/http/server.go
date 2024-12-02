package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	middlerwares "golipors/api/http/middlewares"
	"golipors/app"
	"golipors/config"
)

func Bootstrap(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	router.Group("/api/v1", middlerwares.RateLimiter())

	return router.Listen(fmt.Sprintf(":%d", cfg.Port))
}
