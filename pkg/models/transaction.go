package models

import "time"

type Transaction struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	User       *User   `gorm:"foreignKey:UserID"`
	Amount     float64 `gorm:"not null"`
	VoteCount  int     `gorm:"default:0"`
	Type       string  `gorm:"not null"`
	SurveyID   *uint
	Survey     *Survey `gorm:"foreignKey:SurveyID"`
	ResponseID *uint
	Response   *Response `gorm:"foreignKey:ResponseID"`
	Timestamp  time.Time `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
