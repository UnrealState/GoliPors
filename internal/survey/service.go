// internal/survey/service/service.go
package service

import (
	"context"
	"errors"
	"golipors/internal/survey/domain"
	"golipors/internal/survey/port"
)

type surveyService struct {
	repository port.Repository
}

func NewService(repository port.Repository) port.Service {
	return &surveyService{
		repository: repository,
	}
}

func (s *surveyService) CreateSurvey(ctx context.Context, survey domain.Survey) (uint, error) {
	if err := survey.Validate(); err != nil {
		return 0, err
	}
	return s.repository.CreateSurvey(ctx, survey)
}

func (s *surveyService) GetSurveyByID(ctx context.Context, id uint) (*domain.Survey, error) {
	return s.repository.GetSurveyByID(ctx, id)
}

func (s *surveyService) UpdateSurvey(ctx context.Context, survey domain.Survey) error {
	existingSurvey, err := s.repository.GetSurveyByID(ctx, survey.ID)
	if err != nil {
		return err
	}
	if existingSurvey == nil {
		return errors.New("survey not found")
	}

	// Merge existing and updated survey data
	if survey.Title != "" {
		existingSurvey.Title = survey.Title
	}
	if survey.StartTime != nil {
		existingSurvey.StartTime = survey.StartTime
	}
	if survey.EndTime != nil {
		existingSurvey.EndTime = survey.EndTime
	}
	if survey.RandomOrder != existingSurvey.RandomOrder {
		existingSurvey.RandomOrder = survey.RandomOrder
	}
	if survey.AllowReturn != existingSurvey.AllowReturn {
		existingSurvey.AllowReturn = survey.AllowReturn
	}
	if survey.NumParticipationAttempts != 0 {
		existingSurvey.NumParticipationAttempts = survey.NumParticipationAttempts
	}
	if survey.ResponseTime != 0 {
		existingSurvey.ResponseTime = survey.ResponseTime
	}
	if survey.AnonymityLevel != "" {
		existingSurvey.AnonymityLevel = survey.AnonymityLevel
	}
	if survey.DemographicRestrictions != "" {
		existingSurvey.DemographicRestrictions = survey.DemographicRestrictions
	}
	if survey.ResponseModification != existingSurvey.ResponseModification {
		existingSurvey.ResponseModification = survey.ResponseModification
	}

	if err := existingSurvey.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateSurvey(ctx, *existingSurvey)
}

func (s *surveyService) DeleteSurvey(ctx context.Context, id uint) error {
	return s.repository.DeleteSurvey(ctx, id)
}
