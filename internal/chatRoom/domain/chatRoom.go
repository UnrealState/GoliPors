package domain

import (
	messageDomain "golipors/internal/message/domain"
	surveyDomain "golipors/internal/survey/domain"
	"time"
)

type Chatroom struct {
	ID        uint `gorm:"primaryKey"`
	SurveyID  uint
	Survey    *surveyDomain.Survey     `gorm:"foreignKey:SurveyID"`
	Messages  []*messageDomain.Message `gorm:"foreignKey:ChatroomID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
