package domain

import (
	responseDomain "golipors/internal/response/domain"
	surveyDomain "golipors/internal/survey/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type TransactionID uint

type Transaction struct {
	ID         TransactionID
	UserID     userDomain.UserID
	User       *userDomain.User
	Amount     float64
	VoteCount  int
	Type       string
	SurveyID   *surveyDomain.SurveyID
	Survey     *surveyDomain.Survey
	ResponseID *responseDomain.ResponseID
	Response   *responseDomain.Response
	Timestamp  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
