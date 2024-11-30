package main

import (
	"flag"
	"golipors/config"
	"os"
)

var configPath = flag.String("config", "config.json", "Path to service config file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	_ = config.MustReadConfig(*configPath)
}
