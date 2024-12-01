package models

import "time"

type Notification struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      *User     `gorm:"foreignKey:UserID"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null"`
}
