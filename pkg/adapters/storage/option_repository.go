package storage

import (
	"golipors/internal/option/domain"
	"golipors/internal/option/port"
	"gorm.io/gorm"
)

type optionRepository struct {
	db *gorm.DB
}

func NewOptionRepository(db *gorm.DB) port.OptionRepository {
	return &optionRepository{db: db}
}

func (r *optionRepository) CreateOption(option *domain.Option) error {
	return r.db.Create(option).Error
}

func (r *optionRepository) GetOptionByID(id uint) (*domain.Option, error) {
	var option domain.Option
	err := r.db.First(&option, id).Error
	return &option, err
}

func (r *optionRepository) UpdateOption(option *domain.Option) error {
	return r.db.Save(option).Error
}

func (r *optionRepository) DeleteOption(id uint) error {
	return r.db.Delete(&domain.Option{}, id).Error
}

func (r *optionRepository) ListOptionsByQuestion(questionID uint) ([]*domain.Option, error) {
	var options []*domain.Option
	err := r.db.Where("question_id = ?", questionID).Find(&options).Error
	return options, err
}
