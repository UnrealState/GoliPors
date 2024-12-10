package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"golipors/pkg/adapters/storage/helpers"
	"golipors/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

func GetQuestionnaireMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: helpers.GenerateMigrationID("add_questionnaire_table"),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&types.Questionnaire{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("questionnaires")
			},
		},
	}
}
