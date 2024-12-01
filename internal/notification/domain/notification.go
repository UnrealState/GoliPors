package domain

import (
	userDomain "golipors/internal/user/domain"
	"time"
)

type Notification struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      *userDomain.User `gorm:"foreignKey:UserID"`
	Message   string           `gorm:"type:text;not null"`
	CreatedAt time.Time        `gorm:"not null"`
}
