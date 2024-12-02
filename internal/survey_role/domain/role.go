package domain

import (
	surveyDomain "golipors/internal/survey/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type RoleID uint

type Role struct {
	ID             RoleID
	SurveyID       surveyDomain.SurveyID
	Survey         *surveyDomain.Survey
	UserID         userDomain.UserID
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
