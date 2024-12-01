package domain

import (
	questionDomain "golipors/internal/question/domain"
	surveyDomain "golipors/internal/survey/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Response struct {
	ID                uint
	UserID            uint
	User              *userDomain.User
	SurveyID          uint
	Survey            *surveyDomain.Survey
	QuestionID        uint
	Question          *questionDomain.Question
	ResponseText      string
	SelectedOptionIDs string
	EncryptedData     []byte
	Secret            string
	ResponseTime      time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
