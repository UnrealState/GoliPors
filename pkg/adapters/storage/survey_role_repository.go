package storage

import (
	"golipors/internal/survey_role/domain"
	"golipors/internal/survey_role/port"
	"gorm.io/gorm"
)

type surveyRoleRepository struct {
	db *gorm.DB
}

func NewSurveyRoleRepository(db *gorm.DB) port.SurveyRoleRepository {
	return &surveyRoleRepository{db: db}
}

func (r *surveyRoleRepository) AssignRole(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *surveyRoleRepository) GetRolesBySurvey(surveyID uint) ([]*domain.Role, error) {
	var roles []*domain.Role
	err := r.db.Where("survey_id = ?", surveyID).Find(&roles).Error
	return roles, err
}

func (r *surveyRoleRepository) GetRolesByUser(userID uint) ([]*domain.Role, error) {
	var roles []*domain.Role
	err := r.db.Where("user_id = ?", userID).Find(&roles).Error
	return roles, err
}
