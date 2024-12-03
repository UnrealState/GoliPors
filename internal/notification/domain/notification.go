package domain

import "time"

type Notification struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null"`
}
