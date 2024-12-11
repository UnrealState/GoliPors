package main

import (
	"flag"
	"golipors/api/http"
	"golipors/app"
	"golipors/config"
	"os"
)

var configPath = flag.String("config", "config.json", "Path to service config file")

// @title GoliPors Api Doc
// @version 1.0
// @description An online form platform
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.MustNewApp(c)

	err := http.Bootstrap(appContainer, c.Server)

	if err != nil {
		return
	}
}
