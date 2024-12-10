package services

import (
	"context"
	"golipors/api/http/handlers/helpers"
	"golipors/app"
	"golipors/config"
	userPort "golipors/internal/user/port"
)

type QuestionnaireService struct {
	svc userPort.Service
}

func NewQuestionnaireService(
	svc userPort.Service,
) *QuestionnaireService {
	return &QuestionnaireService{
		svc: svc,
	}
}

func QuestionnaireServiceGetter(appContainer app.App, cfg config.ServerConfig) helpers.ServiceGetter[*QuestionnaireService] {
	return func(ctx context.Context) *QuestionnaireService {
		return NewQuestionnaireService(
			appContainer.UserService(ctx),
		)
	}
}
