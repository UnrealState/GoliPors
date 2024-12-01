package main

import (
	"flag"
	"golipors/app"
	"golipors/config"
	"os"
)

var configPath = flag.String("config", "config.json", "Path to service config file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	_ = app.MustNewApp(c)

	// ToDo Http bootstrap
	/*err := http.Bootstrap(appContainer, c.Server)

	if err != nil {
		return
	}*/
}
