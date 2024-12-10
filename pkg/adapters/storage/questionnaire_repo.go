package storage

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"golipors/internal/questionnaire/port"
	"golipors/pkg/adapters/storage/migrations"
	"gorm.io/gorm"
)

type questionnaireRepo struct {
	db *gorm.DB
}

func NewQuestionnaireRepo(db *gorm.DB) port.Repo {
	return &questionnaireRepo{db}
}

func (r *questionnaireRepo) RunMigrations() error {
	migrator := gormigrate.New(r.db, gormigrate.DefaultOptions, migrations.GetQuestionnaireMigrations())
	return migrator.Migrate()
}
