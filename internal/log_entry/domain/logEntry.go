package domain

import (
	userDomain "golipors/internal/user/domain"
	"time"
)

type LogEntry struct {
	ID        uint
	Timestamp time.Time
	Level     string
	Service   string
	Endpoint  string
	UserID    *uint
	User      *userDomain.User
	Message   string
	Context   string
}
