package domain

import (
	userDomain "golipors/internal/user/domain"
	"time"
)

type LogEntryID uint

type LogEntry struct {
	ID        LogEntryID
	Timestamp time.Time
	Level     string
	Service   string
	Endpoint  string
	UserID    *userDomain.UserID
	User      *userDomain.User
	Message   string
	Context   string
}
