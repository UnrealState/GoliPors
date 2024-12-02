package port

import "golipors/internal/response/domain"

type ResponseService interface {
	SubmitResponse(response *domain.Response) error
	GetResponseDetails(id uint) (*domain.Response, error)
	ListResponsesBySurvey(surveyID uint) ([]*domain.Response, error)
}
