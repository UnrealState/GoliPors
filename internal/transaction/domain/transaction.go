package domain

import "time"

type Transaction struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `gorm:"not null"`
	Amount     float64 `gorm:"not null"`
	VoteCount  int     `gorm:"default:0"`
	Type       string  `gorm:"not null"`
	SurveyID   *uint
	ResponseID *uint
	Timestamp  time.Time `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
