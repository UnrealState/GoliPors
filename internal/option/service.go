package option

import (
	"golipors/internal/option/domain"
	"golipors/internal/option/port"
)

type optionService struct {
	repo port.OptionRepository
}

func NewOptionService(repo port.OptionRepository) port.OptionService {
	return &optionService{repo: repo}
}

func (s *optionService) AddOption(option *domain.Option) error {
	return s.repo.CreateOption(option)
}

func (s *optionService) GetOptionDetails(id uint) (*domain.Option, error) {
	return s.repo.GetOptionByID(id)
}

func (s *optionService) UpdateOption(option *domain.Option) error {
	return s.repo.UpdateOption(option)
}

func (s *optionService) DeleteOption(id uint) error {
	return s.repo.DeleteOption(id)
}

func (s *optionService) ListOptionsByQuestion(questionID uint) ([]*domain.Option, error) {
	return s.repo.ListOptionsByQuestion(questionID)
}
