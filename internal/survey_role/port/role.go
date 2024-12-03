package port

import "golipors/internal/survey_role/domain"

type SurveyRoleRepository interface {
	AssignRole(role *domain.SurveyRole) error
	GetRolesBySurvey(surveyID uint) ([]*domain.SurveyRole, error)
	GetRolesByUser(userID uint) ([]*domain.SurveyRole, error)
}
