// pkg/adapters/storage/survey_repository.go
package storage

import (
	"context"
	"errors"
	"golipors/internal/survey/domain"
	"golipors/internal/survey/port"
	"golipors/pkg/adapters/storage/mapper"
	"golipors/pkg/adapters/storage/migrations"
	"golipors/pkg/adapters/storage/types"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type surveyRepository struct {
	db *gorm.DB
}

func NewSurveyRepository(db *gorm.DB) port.Repository {
	return &surveyRepository{
		db: db,
	}
}

func (r *surveyRepository) CreateSurvey(ctx context.Context, survey domain.Survey) (uint, error) {
	model := mapper.DomainToModel(survey)
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return 0, err
	}
	return model.ID, nil
}

func (r *surveyRepository) GetSurveyByID(ctx context.Context, id uint) (*domain.Survey, error) {
	var model types.Survey
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	survey := mapper.ModelToDomain(model)
	return &survey, nil
}

func (r *surveyRepository) UpdateSurvey(ctx context.Context, survey domain.Survey) error {
	model := mapper.DomainToModel(survey)
	if err := r.db.WithContext(ctx).Model(&types.Survey{}).Where("id = ?", survey.ID).Updates(&model).Error; err != nil {
		return err
	}
	return nil
}

func (r *surveyRepository) DeleteSurvey(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&types.Survey{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *surveyRepository) RunMigrations() error {
	migrator := gormigrate.New(r.db, gormigrate.DefaultOptions, migrations.GetUserMigrations())
	return migrator.Migrate()
}
