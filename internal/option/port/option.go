package port

import "golipors/internal/option/domain"

type OptionRepository interface {
	CreateOption(option *domain.Option) error
	GetOptionByID(id uint) (*domain.Option, error)
	UpdateOption(option *domain.Option) error
	DeleteOption(id uint) error
	ListOptionsByQuestion(questionID uint) ([]*domain.Option, error)
}
