package storage

import (
	"golipors/internal/question/domain"
	"golipors/internal/question/port"
	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) port.QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) CreateQuestion(question *domain.Question) error {
	return r.db.Create(question).Error
}

func (r *questionRepository) GetQuestionByID(id uint) (*domain.Question, error) {
	var question domain.Question
	err := r.db.Preload("Options").First(&question, id).Error
	return &question, err
}

func (r *questionRepository) UpdateQuestion(question *domain.Question) error {
	return r.db.Save(question).Error
}

func (r *questionRepository) DeleteQuestion(id uint) error {
	return r.db.Delete(&domain.Question{}, id).Error
}

func (r *questionRepository) ListQuestionsBySurvey(surveyID uint) ([]*domain.Question, error) {
	var questions []*domain.Question
	err := r.db.Where("survey_id = ?", surveyID).
		Preload("Options").
		Order("order").
		Find(&questions).Error
	return questions, err
}
