package app

import (
	"golipors/config"
	userPort "golipors/internal/user/port"
	"golipors/pkg/cache"
)

type App interface {
	Config() config.Config
	Cache() cache.Provider
	UserService() userPort.Service
}
