package domain

import (
	"time"

	"golipors/pkg/utils"

	"gorm.io/gorm"
)

type User struct {
	ID               uint      `gorm:"primaryKey"`
	NationalID       string    `gorm:"size:10;not null;unique"`
	Email            string    `gorm:"not null;unique"`
	Password         string    `gorm:"not null"` // Hashed password
	FirstName        string    `gorm:"not null"`
	LastName         string    `gorm:"not null"`
	DateOfBirth      time.Time `gorm:"not null"`
	RegistrationDate time.Time `gorm:"not null"`
	City             string
	WalletBalance    float64 `gorm:"default:0"`
	VoteBalance      int     `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
