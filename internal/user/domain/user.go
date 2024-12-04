package domain

import (
	"time"
)

type UserID uint

type User struct {
	ID            UserID
	NationalID    string
	Email         string
	Password      string
	FirstName     string
	LastName      string
	Birthday      time.Time
	City          string
	WalletBalance float64
	VoteBalance   int
	CreatedAt     time.Time
	DeletedAt     time.Time
}
