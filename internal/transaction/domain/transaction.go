package domain

import (
	responseDomain "golipors/internal/response/domain"
	surveyDomain "golipors/internal/survey/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Transaction struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	User       *userDomain.User `gorm:"foreignKey:UserID"`
	Amount     float64          `gorm:"not null"`
	VoteCount  int              `gorm:"default:0"`
	Type       string           `gorm:"not null"`
	SurveyID   *uint
	Survey     *surveyDomain.Survey `gorm:"foreignKey:SurveyID"`
	ResponseID *uint
	Response   *responseDomain.Response `gorm:"foreignKey:ResponseID"`
	Timestamp  time.Time                `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
