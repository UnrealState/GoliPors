package types

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	NationalID    string `gorm:"size:16;uniqueIndex"`
	Email         string `gorm:"size:255;uniqueIndex"`
	Password      string `gorm:"size:255"`
	FirstName     string `gorm:"size:100"`
	LastName      string `gorm:"size:100"`
	Birthday      time.Time
	City          string `gorm:"size:100"`
	WalletBalance float64
	VoteBalance   int
}
