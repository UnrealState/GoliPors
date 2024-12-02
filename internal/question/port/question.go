package port

import "golipors/internal/question/domain"

type QuestionRepository interface {
	CreateQuestion(question *domain.Question) error
	GetQuestionByID(id uint) (*domain.Question, error)
	UpdateQuestion(question *domain.Question) error
	DeleteQuestion(id uint) error
	ListQuestionsBySurvey(surveyID uint) ([]*domain.Question, error)
}
