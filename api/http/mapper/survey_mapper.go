// api/http/mapper/survey_mapper.go
package mapper

import (
	"golipors/api/http/handlers/presenter"
	"golipors/internal/survey/domain"
)

func CreateSurveyRequestToDomain(req presenter.CreateSurveyRequest, ownerID uint) domain.Survey {
	return domain.Survey{
		Title:                    req.Title,
		CreationTime:             req.CreationTime,
		StartTime:                req.StartTime,
		EndTime:                  req.EndTime,
		RandomOrder:              req.RandomOrder,
		AllowReturn:              req.AllowReturn,
		NumParticipationAttempts: req.NumParticipationAttempts,
		ResponseTime:             req.ResponseTime,
		AnonymityLevel:           req.AnonymityLevel,
		DemographicRestrictions:  req.DemographicRestrictions,
		ResponseModification:     req.ResponseModification,
		OwnerID:                  ownerID,
	}
}

func DomainSurveyToCreateSurveyResponse(survey domain.Survey) presenter.CreateSurveyResponse {
	return presenter.CreateSurveyResponse{
		ID:      survey.ID,
		Title:   survey.Title,
		OwnerID: survey.OwnerID,
	}
}

func DomainSurveyToGetSurveyResponse(survey domain.Survey) presenter.GetSurveyResponse {
	return presenter.GetSurveyResponse{
		ID:                       survey.ID,
		Title:                    survey.Title,
		CreationTime:             survey.CreationTime,
		StartTime:                survey.StartTime,
		EndTime:                  survey.EndTime,
		RandomOrder:              survey.RandomOrder,
		AllowReturn:              survey.AllowReturn,
		NumParticipationAttempts: survey.NumParticipationAttempts,
		ResponseTime:             survey.ResponseTime,
		AnonymityLevel:           survey.AnonymityLevel,
		DemographicRestrictions:  survey.DemographicRestrictions,
		ResponseModification:     survey.ResponseModification,
	}
}

func UpdateSurveyRequestToDomain(req presenter.UpdateSurveyRequest) domain.Survey {
	survey := domain.Survey{}

	if req.Title != nil {
		survey.Title = *req.Title
	}
	if req.StartTime != nil {
		survey.StartTime = req.StartTime
	}
	if req.EndTime != nil {
		survey.EndTime = req.EndTime
	}
	if req.RandomOrder != nil {
		survey.RandomOrder = *req.RandomOrder
	}
	if req.AllowReturn != nil {
		survey.AllowReturn = *req.AllowReturn
	}
	if req.NumParticipationAttempts != nil {
		survey.NumParticipationAttempts = *req.NumParticipationAttempts
	}
	if req.ResponseTime != nil {
		survey.ResponseTime = *req.ResponseTime
	}
	if req.AnonymityLevel != nil {
		survey.AnonymityLevel = *req.AnonymityLevel
	}
	if req.DemographicRestrictions != nil {
		survey.DemographicRestrictions = *req.DemographicRestrictions
	}
	if req.ResponseModification != nil {
		survey.ResponseModification = *req.ResponseModification
	}

	return survey
}

func DomainSurveyToUpdateSurveyResponse(survey domain.Survey) presenter.UpdateSurveyResponse {
	return presenter.UpdateSurveyResponse{
		ID:    survey.ID,
		Title: survey.Title,
	}
}
