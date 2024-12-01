package domain

import (
	messageDomain "golipors/internal/message/domain"
	surveyDomain "golipors/internal/survey/domain"
	"time"
)

type Chatroom struct {
	ID        uint
	SurveyID  uint
	Survey    *surveyDomain.Survey
	Messages  []*messageDomain.Message
	CreatedAt time.Time
	UpdatedAt time.Time
}
