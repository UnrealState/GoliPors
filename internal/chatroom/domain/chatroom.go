package domain

import "time"

type Chatroom struct {
	ID        uint      `gorm:"primaryKey"`
	SurveyID  uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
