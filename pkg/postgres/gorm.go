package postgres

import (
	"fmt"
	"golipors/config"
	"golipors/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg config.DBConfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
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
		return err
	}

	DB = db
	return nil
}
