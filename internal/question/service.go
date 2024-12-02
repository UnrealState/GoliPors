package question

import (
	"golipors/internal/question/domain"
	"golipors/internal/question/port"
)

type questionService struct {
	repo port.QuestionRepository
}

func NewQuestionService(repo port.QuestionRepository) port.QuestionService {
	return &questionService{repo: repo}
}

func (s *questionService) AddQuestion(question *domain.Question) error {
	return s.repo.CreateQuestion(question)
}

func (s *questionService) GetQuestionDetails(id uint) (*domain.Question, error) {
	return s.repo.GetQuestionByID(id)
}

func (s *questionService) UpdateQuestion(question *domain.Question) error {
	return s.repo.UpdateQuestion(question)
}

func (s *questionService) DeleteQuestion(id uint) error {
	return s.repo.DeleteQuestion(id)
}

func (s *questionService) ListQuestionsBySurvey(surveyID uint) ([]*domain.Question, error) {
	return s.repo.ListQuestionsBySurvey(surveyID)
}
