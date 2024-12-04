package migrations

import (
	"golipors/pkg/adapters/storage/helpers"
	"golipors/pkg/adapters/storage/types"
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func GetUserMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: helpers.GenerateMigrationID("add_users_table"),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&types.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
	}
}
