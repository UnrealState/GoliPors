package models

import "time"

type Option struct {
	ID         uint `gorm:"primaryKey"`
	QuestionID uint
	Question   *Question `gorm:"foreignKey:QuestionID"`
	Text       string    `gorm:"not null"`
	IsCorrect  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
