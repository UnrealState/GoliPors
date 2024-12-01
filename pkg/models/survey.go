package models

import (
	"time"
)

type Survey struct {
	ID                       uint      `gorm:"primaryKey"`
	Title                    string    `gorm:"not null"`
	CreationTime             time.Time `gorm:"not null"`
	StartTime                *time.Time
	EndTime                  *time.Time
	RandomOrder              bool `gorm:"default:false"`
	AllowReturn              bool `gorm:"default:false"`
	NumParticipationAttempts int  `gorm:"default:1"`
	ResponseTime             time.Duration
	AnonymityLevel           string `gorm:"not null"`
	OwnerID                  uint
	Owner                    *User         `gorm:"foreignKey:OwnerID"`
	DemographicRestrictions  string        `gorm:"type:text"`
	ResponseModification     bool          `gorm:"default:false"`
	Questions                []*Question   `gorm:"foreignKey:SurveyID"`
	Chatroom                 *Chatroom     `gorm:"foreignKey:SurveyID"`
	SurveyRoles              []*SurveyRole `gorm:"foreignKey:SurveyID"`
	Responses                []*Response   `gorm:"foreignKey:SurveyID"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
