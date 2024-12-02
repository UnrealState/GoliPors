package domain

import (
	userDomain "golipors/internal/user/domain"
	"time"
)

type NotificationID uint

type Notification struct {
	ID        NotificationID
	UserID    userDomain.UserID
	User      *userDomain.User
	Message   string
	CreatedAt time.Time
}
