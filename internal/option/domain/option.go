package domain

import (
	questionDomain "golipors/internal/question/domain"
	"time"
)

type OptionID uint

type Option struct {
	ID         OptionID
	QuestionID questionDomain.QuestionID
	Question   *questionDomain.Question
	Text       string
	IsCorrect  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
