// api/http/dto/create_survey.go
package dto

import "time"

type CreateSurveyRequest struct {
	Title                    string     `json:"title" validate:"required,min=3,max=100"`
	CreationTime             time.Time  `json:"creation_time" validate:"required"`
	StartTime                *time.Time `json:"start_time" validate:"required"`
	EndTime                  *time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	RandomOrder              bool       `json:"random_order"`
	AllowReturn              bool       `json:"allow_return"`
	NumParticipationAttempts int        `json:"num_participation_attempts" validate:"gte=1,lte=10"`
	ResponseTime             int        `json:"response_time" validate:"gte=60,lte=86400"` // Between 1 minute and 1 day
	AnonymityLevel           string     `json:"anonymity_level" validate:"required,oneof=visible_to_creator visible_to_creator_and_admins anonymous"`
	DemographicRestrictions  string     `json:"demographic_restrictions" validate:"omitempty,json"`
	ResponseModification     bool       `json:"response_modification"`
}

type CreateSurveyResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	OwnerID uint   `json:"owner_id"`
}
