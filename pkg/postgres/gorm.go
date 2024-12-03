package postgres

import (
	"fmt"

	// Importing domain models with unique aliases to avoid naming conflicts
	chatroomDomain "golipors/internal/chatroom/domain"
	logEntryDomain "golipors/internal/log_entry/domain"
	messageDomain "golipors/internal/message/domain"
	notificationDomain "golipors/internal/notification/domain"
	optionDomain "golipors/internal/option/domain"
	questionDomain "golipors/internal/question/domain"
	responseDomain "golipors/internal/response/domain"
	surveyDomain "golipors/internal/survey/domain"
	surveyRoleDomain "golipors/internal/survey_role/domain"
	systemRoleDomain "golipors/internal/system_role/domain"
	transactionDomain "golipors/internal/transaction/domain"
	userDomain "golipors/internal/user/domain"
	// userSystemRoleDomain "golipors/internal/user_system_role/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConnOptions holds the configuration for the PostgreSQL connection.
type DBConnOptions struct {
	Host   string
	Port   uint
	User   string
	Pass   string
	Name   string
	Schema string
}

// PostgresDSN constructs the Data Source Name for PostgreSQL connection.
func (o DBConnOptions) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		o.Host, o.Port, o.User, o.Pass, o.Name, o.Schema)
}

// NewPsqlGormConnection establishes a new GORM connection to PostgreSQL and migrates the schema.
func NewPsqlGormConnection(opt DBConnOptions) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // Adjust logging level as needed
	})

	if err != nil {
		return db, err
	}

	// Migrate the schema using the refactored domain models.
	err = db.AutoMigrate(
		&userDomain.User{},
		&systemRoleDomain.SystemRole{},
		// &userSystemRoleDomain.UserSystemRole{},
		&surveyDomain.Survey{},
		&surveyRoleDomain.SurveyRole{},
		&questionDomain.Question{},
		&optionDomain.Option{},
		&responseDomain.Response{},
		&chatroomDomain.Chatroom{},
		&messageDomain.Message{},
		&notificationDomain.Notification{},
		&transactionDomain.Transaction{},
		&logEntryDomain.LogEntry{},
	)
	if err != nil {
		return db, err
	}

	return db, nil
}
