package domain

import (
	surveyDomain "golipors/internal/survey/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Role struct {
	ID             uint
	SurveyID       uint
	Survey         *surveyDomain.Survey
	UserID         uint
	User           *userDomain.User
	RoleName       string
	IsTemporary    bool
	ExpiryTime     *time.Time
	CanViewSurvey  bool
	CanAssignVotes bool
	CanCastVotes   bool
	CanEditSurvey  bool
	CanAddVotes    bool
	CanAssignRoles bool
	CanViewReports bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
