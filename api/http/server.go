package http

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"golipors/api/http/handlers"
	middlerwares "golipors/api/http/middlewares"
	di "golipors/app"
	"golipors/config"
)

func Bootstrap(appContainer di.App, cfg config.ServerConfig) error {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(
		swagger.New(swagger.Config{
			BasePath: "/",
			FilePath: "./docs/api/swagger.json",
			Path:     "swagger",
		}),
		middlerwares.RateLimiter(),
	)

	api := app.Group("/api/v1", middlerwares.SetUserContext)

	handlers.RegisterAccountHandlers(api, appContainer, cfg)

	return app.Listen(fmt.Sprintf(":%d", cfg.Port))
}
