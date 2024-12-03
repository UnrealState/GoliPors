package domain

import "time"

type Question struct {
	ID               uint   `gorm:"primaryKey"`
	SurveyID         uint   `gorm:"not null"`
	Text             string `gorm:"not null"`
	Type             string `gorm:"not null"`
	Order            int
	AttachmentURL    string
	CorrectOptionIDs string `gorm:"type:text"`
	IsConditional    bool
	Condition        string `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
