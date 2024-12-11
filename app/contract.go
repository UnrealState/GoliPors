package app

import (
	"context"
	"golipors/config"
	userPort "golipors/internal/user/port"
	"golipors/pkg/adapters/email"
	"golipors/pkg/cache"
	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	Cache() cache.Provider
	MailService() email.Adapter
	UserService(ctx context.Context) userPort.Service
}
