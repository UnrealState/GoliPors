package domain

import (
	surveyDomain "golipors/internal/survey/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Role struct {
	ID             uint `gorm:"primaryKey"`
	SurveyID       uint
	Survey         *surveyDomain.Survey `gorm:"foreignKey:SurveyID"`
	UserID         uint
	User           *userDomain.User `gorm:"foreignKey:UserID"`
	RoleName       string           `gorm:"not null"`
	IsTemporary    bool             `gorm:"default:false"`
	ExpiryTime     *time.Time
	CanViewSurvey  bool `gorm:"default:false"`
	CanAssignVotes bool `gorm:"default:false"`
	CanCastVotes   bool `gorm:"default:false"`
	CanEditSurvey  bool `gorm:"default:false"`
	CanAddVotes    bool `gorm:"default:false"`
	CanAssignRoles bool `gorm:"default:false"`
	CanViewReports bool `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
