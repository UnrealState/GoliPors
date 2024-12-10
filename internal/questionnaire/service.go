package questionnaire

import (
	"golipors/internal/questionnaire/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) RunMigrations() error {
	return s.repo.RunMigrations()
}
