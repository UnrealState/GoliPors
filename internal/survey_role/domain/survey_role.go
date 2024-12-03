package domain

import "time"

type SurveyRole struct {
	ID             uint   `gorm:"primaryKey"`
	SurveyID       uint   `gorm:"not null"`
	UserID         uint   `gorm:"not null"`
	RoleName       string `gorm:"not null"`
	IsTemporary    bool   `gorm:"default:false"`
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
