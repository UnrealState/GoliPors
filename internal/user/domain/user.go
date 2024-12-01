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
	ID               uint      `gorm:"primaryKey"`
	NationalID       string    `gorm:"size:10;not null;unique"`
	Email            string    `gorm:"not null;unique"`
	Password         string    `gorm:"not null"` // Hashed password
	FirstName        string    `gorm:"not null"`
	LastName         string    `gorm:"not null"`
	DateOfBirth      time.Time `gorm:"not null"`
	RegistrationDate time.Time `gorm:"not null"`
	City             string
	WalletBalance    float64 `gorm:"default:0"`
	VoteBalance      int     `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	SystemRoles      []*systemRoleDomain.Role           `gorm:"many2many:user_system_roles;"`
	SurveyRoles      []*surveyRoleDomain.Role           `gorm:"foreignKey:UserID"`
	Surveys          []*surveyDomain.Survey             `gorm:"foreignKey:OwnerID"`
	Notifications    []*notificationDomain.Notification `gorm:"foreignKey:UserID"`
	Messages         []*messageDomain.Message           `gorm:"foreignKey:UserID"`
	Responses        []*responseDomain.Response         `gorm:"foreignKey:UserID"`
	Transactions     []*transactionDomain.Transaction   `gorm:"foreignKey:UserID"`
}
