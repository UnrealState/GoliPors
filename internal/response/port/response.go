package port

import "golipors/internal/response/domain"

type ResponseRepository interface {
	CreateResponse(response *domain.Response) error
	GetResponseByID(id uint) (*domain.Response, error)
	ListResponsesBySurvey(surveyID uint) ([]*domain.Response, error)
}
