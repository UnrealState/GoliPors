package survey_role

import (
	"golipors/internal/survey_role/domain"
	"golipors/internal/survey_role/port"
)

type surveyRoleService struct {
	repo port.SurveyRoleRepository
}

func NewSurveyRoleService(repo port.SurveyRoleRepository) port.SurveyRoleService {
	return &surveyRoleService{repo: repo}
}

func (s *surveyRoleService) AssignSurveyRole(role *domain.SurveyRole) error {
	return s.repo.AssignRole(role)
}

func (s *surveyRoleService) ListSurveyRolesBySurvey(surveyID uint) ([]*domain.SurveyRole, error) {
	return s.repo.GetRolesBySurvey(surveyID)
}

func (s *surveyRoleService) ListSurveyRolesByUser(userID uint) ([]*domain.SurveyRole, error) {
	return s.repo.GetRolesByUser(userID)
}
