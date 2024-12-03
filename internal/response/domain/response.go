package domain

import "time"

type Response struct {
	ID                uint      `gorm:"primaryKey"`
	UserID            uint      `gorm:"not null"`
	SurveyID          uint      `gorm:"not null"`
	QuestionID        uint      `gorm:"not null"`
	ResponseText      string    `gorm:"type:text"`
	SelectedOptionIDs string    `gorm:"type:text"`
	EncryptedData     []byte    `gorm:"type:bytea"`
	Secret            string    `gorm:"not null"`
	ResponseTime      time.Time `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
