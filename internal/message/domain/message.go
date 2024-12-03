package domain

import "time"

type Message struct {
	ID         uint      `gorm:"primaryKey"`
	ChatroomID uint      `gorm:"not null"`
	UserID     uint      `gorm:"not null"`
	Content    string    `gorm:"type:text;not null"`
	CreatedAt  time.Time `gorm:"not null"`
}
