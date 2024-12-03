package survey

import (
	"golipors/internal/survey/domain"
	"golipors/internal/survey/port"
)

type surveyService struct {
	repo port.SurveyRepository
}

func NewSurveyService(repo port.SurveyRepository) port.SurveyService {
	return &surveyService{repo: repo}
}

func (s *surveyService) CreateSurvey(survey *domain.Survey) error {
	//TODO Add business logic (e.g., validate fields)
	return s.repo.CreateSurvey(survey)
}

func (s *surveyService) GetSurveyDetails(id uint) (*domain.Survey, error) {
	return s.repo.GetSurveyByID(id)
}

func (s *surveyService) UpdateSurvey(survey *domain.Survey) error {
	return s.repo.UpdateSurvey(survey)
}

func (s *surveyService) DeleteSurvey(id uint) error {
	return s.repo.DeleteSurvey(id)
}

func (s *surveyService) ListSurveysByOwner(ownerID uint) ([]*domain.Survey, error) {
	return s.repo.ListSurveysByOwner(ownerID)
}
