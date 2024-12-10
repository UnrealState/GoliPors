// internal/survey/domain/survey.go
package domain

import (
	"errors"
	"time"
)

type Survey struct {
	ID                       uint
	Title                    string
	CreationTime             time.Time
	StartTime                *time.Time
	EndTime                  *time.Time
	RandomOrder              bool
	AllowReturn              bool
	NumParticipationAttempts int
	ResponseTime             int // in seconds
	AnonymityLevel           string
	OwnerID                  uint
	DemographicRestrictions  string
	ResponseModification     bool
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

func (s *Survey) Validate() error {
	if s.Title == "" {
		return errors.New("title cannot be empty")
	}
	if s.CreationTime.IsZero() {
		return errors.New("creation_time is required")
	}
	if s.AnonymityLevel == "" {
		return errors.New("anonymity_level cannot be empty")
	}
	return nil
}
