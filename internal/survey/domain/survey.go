package domain

import (
	chatRoomDomain "golipors/internal/chatroom/domain"
	questionDomain "golipors/internal/question/domain"
	responseRoomDomain "golipors/internal/response/domain"
	surveyRoleDomain "golipors/internal/survey_role/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Survey struct {
	ID                       uint
	Title                    string
	CreationTime             time.Time
	StartTime                *time.Time
	EndTime                  *time.Time
	RandomOrder              bool
	AllowReturn              bool
	NumParticipationAttempts int
	ResponseTime             time.Duration
	AnonymityLevel           string
	OwnerID                  uint
	Owner                    *userDomain.User
	DemographicRestrictions  string
	ResponseModification     bool
	Questions                []*questionDomain.Question
	Chatroom                 *chatRoomDomain.Chatroom
	SurveyRoles              []*surveyRoleDomain.Role
	Responses                []*responseRoomDomain.Response
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
