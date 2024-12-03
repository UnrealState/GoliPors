package domain

import "time"

type Option struct {
	ID         uint   `gorm:"primaryKey"`
	QuestionID uint   `gorm:"not null"`
	Text       string `gorm:"not null"`
	IsCorrect  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
