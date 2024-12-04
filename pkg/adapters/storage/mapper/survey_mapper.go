// pkg/adapters/storage/mapper/survey_mapper.go
package mapper

import (
	"golipors/internal/survey/domain"
	"golipors/pkg/adapters/storage/models"
)

func DomainToModel(survey domain.Survey) models.Survey {
	return models.Survey{
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
		OwnerID:                  survey.OwnerID,
		DemographicRestrictions:  survey.DemographicRestrictions,
		ResponseModification:     survey.ResponseModification,
		CreatedAt:                survey.CreatedAt,
		UpdatedAt:                survey.UpdatedAt,
	}
}

func ModelToDomain(survey models.Survey) domain.Survey {
	return domain.Survey{
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
		OwnerID:                  survey.OwnerID,
		DemographicRestrictions:  survey.DemographicRestrictions,
		ResponseModification:     survey.ResponseModification,
		CreatedAt:                survey.CreatedAt,
		UpdatedAt:                survey.UpdatedAt,
	}
}
