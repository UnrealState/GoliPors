// pkg/adapters/storage/models/survey.go
package types

import (
	"time"
)

type Survey struct {
	ID                       uint      `gorm:"primaryKey"`
	Title                    string    `gorm:"not null"`
	CreationTime             time.Time `gorm:"not null"`
	StartTime                *time.Time
	EndTime                  *time.Time
	RandomOrder              bool   `gorm:"default:false"`
	AllowReturn              bool   `gorm:"default:false"`
	NumParticipationAttempts int    `gorm:"default:1"`
	ResponseTime             int    // in seconds
	AnonymityLevel           string `gorm:"not null"`
	OwnerID                  uint
	DemographicRestrictions  string `gorm:"type:text"`
	ResponseModification     bool   `gorm:"default:false"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
