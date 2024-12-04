// api/http/dto/update_survey.go
package dto

import "time"

type UpdateSurveyRequest struct {
	Title                    *string    `json:"title,omitempty"`
	StartTime                *time.Time `json:"start_time,omitempty"`
	EndTime                  *time.Time `json:"end_time,omitempty"`
	RandomOrder              *bool      `json:"random_order,omitempty"`
	AllowReturn              *bool      `json:"allow_return,omitempty"`
	NumParticipationAttempts *int       `json:"num_participation_attempts,omitempty"`
	ResponseTime             *int       `json:"response_time,omitempty"`
	AnonymityLevel           *string    `json:"anonymity_level,omitempty"`
	DemographicRestrictions  *string    `json:"demographic_restrictions,omitempty"`
	ResponseModification     *bool      `json:"response_modification,omitempty"`
}

type UpdateSurveyResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
