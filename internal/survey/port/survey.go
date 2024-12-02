package port

import "golipors/internal/survey/domain"

type SurveyRepository interface {
	CreateSurvey(survey *domain.Survey) error
	GetSurveyByID(id uint) (*domain.Survey, error)
	UpdateSurvey(survey *domain.Survey) error
	DeleteSurvey(id uint) error
	ListSurveysByOwner(ownerID uint) ([]*domain.Survey, error)
}
