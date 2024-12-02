package storage

import (
	"golipors/internal/response/domain"
	"golipors/internal/response/port"
	"gorm.io/gorm"
)

type responseRepository struct {
	db *gorm.DB
}

func NewResponseRepository(db *gorm.DB) port.ResponseRepository {
	return &responseRepository{db: db}
}

func (r *responseRepository) CreateResponse(response *domain.Response) error {
	return r.db.Create(response).Error
}

func (r *responseRepository) GetResponseByID(id uint) (*domain.Response, error) {
	var response domain.Response
	err := r.db.First(&response, id).Error
	return &response, err
}

func (r *responseRepository) ListResponsesBySurvey(surveyID uint) ([]*domain.Response, error) {
	var responses []*domain.Response
	err := r.db.Where("survey_id = ?", surveyID).Find(&responses).Error
	return responses, err
}
