package storage

import (
	"golipors/internal/survey/domain"
	"golipors/internal/survey/port"
	"gorm.io/gorm"
)

type surveyRepository struct {
	db *gorm.DB
}

func NewSurveyRepository(db *gorm.DB) port.SurveyRepository {
	return &surveyRepository{db: db}
}

func (r *surveyRepository) CreateSurvey(survey *domain.Survey) error {
	return r.db.Create(survey).Error
}

func (r *surveyRepository) GetSurveyByID(id uint) (*domain.Survey, error) {
	var survey domain.Survey
	err := r.db.Preload("Questions.Options").
		Preload("SurveyRoles").
		Preload("Responses").
		Preload("Chatroom.Messages").
		First(&survey, id).Error
	return &survey, err
}

func (r *surveyRepository) UpdateSurvey(survey *domain.Survey) error {
	return r.db.Save(survey).Error
}

func (r *surveyRepository) DeleteSurvey(id uint) error {
	return r.db.Delete(&domain.Survey{}, id).Error
}

func (r *surveyRepository) ListSurveysByOwner(ownerID uint) ([]*domain.Survey, error) {
	var surveys []*domain.Survey
	err := r.db.Where("owner_id = ?", ownerID).
		Preload("Questions").
		Preload("SurveyRoles").
		Find(&surveys).Error
	return surveys, err
}
