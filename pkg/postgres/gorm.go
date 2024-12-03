package postgres

import (
	"fmt"
	"golipors/pkg/models"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnOptions struct {
	Host   string
	Port   uint
	User   string
	Pass   string
	Name   string
	Schema string
}

func (o DBConnOptions) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		o.Host, o.Port, o.User, o.Pass, o.Name, o.Schema)
}

func NewPsqlGormConnection(opt DBConnOptions) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{
		Logger: logger.Discard,
	})

	if err != nil {
		return db, err
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&models.User{},
		&models.SystemRole{},
		&models.Survey{},
		&models.SurveyRole{},
		&models.Question{},
		&models.Option{},
		&models.Response{},
		&models.Chatroom{},
		&models.Message{},
		&models.Notification{},
		&models.Transaction{},
		&models.LogEntry{},
	)
	if err != nil {
		return db, err
	}

	return db, nil
}
