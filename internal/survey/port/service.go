package port

import "golipors/internal/survey/domain"

type SurveyService interface {
	CreateSurvey(survey *domain.Survey) error
	GetSurveyDetails(id uint) (*domain.Survey, error)
	UpdateSurvey(survey *domain.Survey) error
	DeleteSurvey(id uint) error
	ListSurveysByOwner(ownerID uint) ([]*domain.Survey, error)
}
