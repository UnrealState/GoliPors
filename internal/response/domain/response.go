package domain

import (
	questionDomain "golipors/internal/question/domain"
	surveyDomain "golipors/internal/survey/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type ResponseID uint

type Response struct {
	ID                ResponseID
	UserID            userDomain.UserID
	User              *userDomain.User
	SurveyID          surveyDomain.SurveyID
	Survey            *surveyDomain.Survey
	QuestionID        questionDomain.QuestionID
	Question          *questionDomain.Question
	ResponseText      string
	SelectedOptionIDs string
	EncryptedData     []byte
	Secret            string
	ResponseTime      time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
