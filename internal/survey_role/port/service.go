package port

import "golipors/internal/survey_role/domain"

type SurveyRoleService interface {
	AssignSurveyRole(role *domain.Role) error
	ListSurveyRolesBySurvey(surveyID uint) ([]*domain.Role, error)
	ListSurveyRolesByUser(userID uint) ([]*domain.Role, error)
}
