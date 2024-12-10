// internal/survey/port/repository.go
package port

import (
	"context"
	"golipors/internal/survey/domain"
)

type Repository interface {
	CreateSurvey(ctx context.Context, survey domain.Survey) (uint, error)
	GetSurveyByID(ctx context.Context, id uint) (*domain.Survey, error)
	UpdateSurvey(ctx context.Context, survey domain.Survey) error
	DeleteSurvey(ctx context.Context, id uint) error
}
