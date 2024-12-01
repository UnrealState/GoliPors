package domain

import (
	userDomain "golipors/internal/user/domain"
	"time"
)

type LogEntry struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"not null"`
	Level     string    `gorm:"not null"`
	Service   string
	Endpoint  string
	UserID    *uint
	User      *userDomain.User `gorm:"foreignKey:UserID"`
	Message   string           `gorm:"type:text;not null"`
	Context   string           `gorm:"type:text"`
}
