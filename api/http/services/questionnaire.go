package services

import (
	"context"
	"golipors/api/http/handlers/helpers"
	"golipors/app"
	"golipors/config"
	questionnairPort "golipors/internal/questionnaire/port"
)

type QuestionnaireService struct {
	svc questionnairPort.Service
}

func NewQuestionnaireService(
	svc questionnairPort.Service,
) *QuestionnaireService {
	return &QuestionnaireService{
		svc: svc,
	}
}

func QuestionnaireServiceGetter(appContainer app.App, cfg config.ServerConfig) helpers.ServiceGetter[*QuestionnaireService] {
	return func(ctx context.Context) *QuestionnaireService {
		return NewQuestionnaireService(
			appContainer.QuestionnaireService(ctx),
		)
	}
}
