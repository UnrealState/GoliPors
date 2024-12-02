package domain

import (
	messageDomain "golipors/internal/message/domain"
	surveyDomain "golipors/internal/survey/domain"
	"time"
)

type ChatroomID uint

type Chatroom struct {
	ID        ChatroomID
	SurveyID  surveyDomain.SurveyID
	Survey    *surveyDomain.Survey
	Messages  []*messageDomain.Message
	CreatedAt time.Time
	UpdatedAt time.Time
}
