package domain

import (
	chatRoomDomain "golipors/internal/chatRoom/domain"
	questionDomain "golipors/internal/question/domain"
	responseRoomDomain "golipors/internal/response/domain"
	surveyRoleDomain "golipors/internal/surveyRole/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Survey struct {
	ID                       uint      `gorm:"primaryKey"`
	Title                    string    `gorm:"not null"`
	CreationTime             time.Time `gorm:"not null"`
	StartTime                *time.Time
	EndTime                  *time.Time
	RandomOrder              bool `gorm:"default:false"`
	AllowReturn              bool `gorm:"default:false"`
	NumParticipationAttempts int  `gorm:"default:1"`
	ResponseTime             time.Duration
	AnonymityLevel           string `gorm:"not null"`
	OwnerID                  uint
	Owner                    *userDomain.User               `gorm:"foreignKey:OwnerID"`
	DemographicRestrictions  string                         `gorm:"type:text"`
	ResponseModification     bool                           `gorm:"default:false"`
	Questions                []*questionDomain.Question     `gorm:"foreignKey:SurveyID"`
	Chatroom                 *chatRoomDomain.Chatroom       `gorm:"foreignKey:SurveyID"`
	SurveyRoles              []*surveyRoleDomain.Role       `gorm:"foreignKey:SurveyID"`
	Responses                []*responseRoomDomain.Response `gorm:"foreignKey:SurveyID"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
