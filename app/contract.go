package app

import (
	"context"
	"golipors/config"
	"golipors/pkg/adapters/email"
	"golipors/pkg/cache"
	"gorm.io/gorm"

	questionnaireService "golipors/internal/questionnaire/port"
	userPort "golipors/internal/user/port"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	Cache() cache.Provider
	MailService() email.Adapter
	UserService(ctx context.Context) userPort.Service
	QuestionnaireService(ctx context.Context) questionnaireService.Service
}
