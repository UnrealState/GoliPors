package port

import "golipors/internal/survey_role/domain"

type SurveyRoleService interface {
	AssignSurveyRole(role *domain.SurveyRole) error
	ListSurveyRolesBySurvey(surveyID uint) ([]*domain.SurveyRole, error)
	ListSurveyRolesByUser(userID uint) ([]*domain.SurveyRole, error)
}
