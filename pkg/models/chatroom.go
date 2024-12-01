package models

import "time"

type Chatroom struct {
	ID        uint `gorm:"primaryKey"`
	SurveyID  uint
	Survey    *Survey    `gorm:"foreignKey:SurveyID"`
	Messages  []*Message `gorm:"foreignKey:ChatroomID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
