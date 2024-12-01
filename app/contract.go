package app

import (
	"golipors/config"
)

type App interface {
	Config() config.Config
	// ToDo Define services
}
