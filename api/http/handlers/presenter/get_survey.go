// api/http/dto/get_survey.go
package presenter

import "time"

type GetSurveyResponse struct {
	ID                       uint       `json:"id"`
	Title                    string     `json:"title"`
	CreationTime             time.Time  `json:"creation_time"`
	StartTime                *time.Time `json:"start_time"`
	EndTime                  *time.Time `json:"end_time"`
	RandomOrder              bool       `json:"random_order"`
	AllowReturn              bool       `json:"allow_return"`
	NumParticipationAttempts int        `json:"num_participation_attempts"`
	ResponseTime             int        `json:"response_time"`
	AnonymityLevel           string     `json:"anonymity_level"`
	DemographicRestrictions  string     `json:"demographic_restrictions"`
	ResponseModification     bool       `json:"response_modification"`
}
