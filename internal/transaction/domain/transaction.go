package domain

import (
	responseDomain "golipors/internal/response/domain"
	surveyDomain "golipors/internal/survey/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Transaction struct {
	ID         uint
	UserID     uint
	User       *userDomain.User
	Amount     float64
	VoteCount  int
	Type       string
	SurveyID   *uint
	Survey     *surveyDomain.Survey `gorm:"foreignKey:SurveyID"`
	ResponseID *uint
	Response   *responseDomain.Response
	Timestamp  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
