package port

import "golipors/internal/question/domain"

type QuestionService interface {
	AddQuestion(question *domain.Question) error
	GetQuestionDetails(id uint) (*domain.Question, error)
	UpdateQuestion(question *domain.Question) error
	DeleteQuestion(id uint) error
	ListQuestionsBySurvey(surveyID uint) ([]*domain.Question, error)
}
