package models

import (
	"time"
)

type Response struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint
	User              *User `gorm:"foreignKey:UserID"`
	SurveyID          uint
	Survey            *Survey `gorm:"foreignKey:SurveyID"`
	QuestionID        uint
	Question          *Question `gorm:"foreignKey:QuestionID"`
	ResponseText      string    `gorm:"type:text"`
	SelectedOptionIDs string    `gorm:"type:text"`
	EncryptedData     []byte    `gorm:"type:bytea"`
	Secret            string    `gorm:"not null"`
	ResponseTime      time.Time `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
