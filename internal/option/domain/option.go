package domain

import (
	questionDomain "golipors/internal/question/domain"
	"time"
)

type Option struct {
	ID         uint
	QuestionID uint
	Question   *questionDomain.Question
	Text       string
	IsCorrect  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
