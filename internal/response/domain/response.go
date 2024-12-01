package domain

import (
	questionDomain "golipors/internal/question/domain"
	surveyDomain "golipors/internal/survey/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Response struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint
	User              *userDomain.User `gorm:"foreignKey:UserID"`
	SurveyID          uint
	Survey            *surveyDomain.Survey `gorm:"foreignKey:SurveyID"`
	QuestionID        uint
	Question          *questionDomain.Question `gorm:"foreignKey:QuestionID"`
	ResponseText      string                   `gorm:"type:text"`
	SelectedOptionIDs string                   `gorm:"type:text"`
	EncryptedData     []byte                   `gorm:"type:bytea"`
	Secret            string                   `gorm:"not null"`
	ResponseTime      time.Time                `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
