// cmd/main.go
package main

import (
	"flag"
	"fmt"
	httpServer "golipors/api/http"
	"golipors/config"
	"golipors/pkg/postgres"
	"golipors/pkg/redis"
	"golipors/pkg/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var configPath = flag.String("config", "config/sample-config.json", "Path to service config file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	// Initialize Database
	err = postgres.InitDB(cfg.DB)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Initialize Redis
	err = redis.InitRedis(cfg.Redis)
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	// Set JWT Secret from Config
	utils.SetJWTSecret(cfg.Server.Secret)

	// Set up Fiber app
	app := fiber.New()

	// Set up routes
	httpServer.SetupRoutes(app, cfg)

	// Start server
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := app.Listen(serverAddr); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
