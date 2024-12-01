package port

import "golipors/internal/survey_role/domain"

type SurveyRoleRepository interface {
	AssignRole(role *domain.Role) error
	GetRolesBySurvey(surveyID uint) ([]*domain.Role, error)
	GetRolesByUser(userID uint) ([]*domain.Role, error)
}
