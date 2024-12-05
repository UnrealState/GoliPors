package app

import (
	"golipors/config"
	userPort "golipors/internal/user/port"
	"golipors/pkg/adapters/email"
	"golipors/pkg/cache"
)

type App interface {
	Config() config.Config
	Cache() cache.Provider
	MailService() email.Adapter
	UserService() userPort.Service
}
