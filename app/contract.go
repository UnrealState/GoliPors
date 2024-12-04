package app

import (
	"golipors/config"
	userPort "golipors/internal/user/port"
)

type App interface {
	Config() config.Config
	UserService() userPort.Service
}
