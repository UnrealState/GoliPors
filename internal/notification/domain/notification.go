package domain

import (
	userDomain "golipors/internal/user/domain"
	"time"
)

type Notification struct {
	ID        uint
	UserID    uint
	User      *userDomain.User
	Message   string
	CreatedAt time.Time
}
