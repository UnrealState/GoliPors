package domain

import (
	messageDomain "golipors/internal/message/domain"
	notificationDomain "golipors/internal/notification/domain"
	responseDomain "golipors/internal/response/domain"
	surveyDomain "golipors/internal/survey/domain"
	surveyRoleDomain "golipors/internal/surveyRole/domain"
	systemRoleDomain "golipors/internal/systemRole/domain"
	transactionDomain "golipors/internal/transaction/domain"
	"time"
)

type User struct {
	ID               uint
	NationalID       string
	Email            string
	Password         string // Hashed password
	FirstName        string
	LastName         string
	DateOfBirth      time.Time
	RegistrationDate time.Time
	City             string
	WalletBalance    float64
	VoteBalance      int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	SystemRoles      []*systemRoleDomain.Role
	SurveyRoles      []*surveyRoleDomain.Role
	Surveys          []*surveyDomain.Survey
	Notifications    []*notificationDomain.Notification
	Messages         []*messageDomain.Message
	Responses        []*responseDomain.Response
	Transactions     []*transactionDomain.Transaction
}
