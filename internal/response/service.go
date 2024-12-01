package response

import (
	"golipors/internal/response/domain"
	"golipors/internal/response/port"
)

type responseService struct {
	repo port.ResponseRepository
}

func NewResponseService(repo port.ResponseRepository) port.ResponseService {
	return &responseService{repo: repo}
}

func (s *responseService) SubmitResponse(response *domain.Response) error {
	return s.repo.CreateResponse(response)
}

func (s *responseService) GetResponseDetails(id uint) (*domain.Response, error) {
	return s.repo.GetResponseByID(id)
}

func (s *responseService) ListResponsesBySurvey(surveyID uint) ([]*domain.Response, error) {
	return s.repo.ListResponsesBySurvey(surveyID)
}
