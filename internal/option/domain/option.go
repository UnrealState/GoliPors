package domain

import (
	questionDomain "golipors/internal/question/domain"
	"time"
)

type Option struct {
	ID         uint `gorm:"primaryKey"`
	QuestionID uint
	Question   *questionDomain.Question `gorm:"foreignKey:QuestionID"`
	Text       string                   `gorm:"not null"`
	IsCorrect  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
