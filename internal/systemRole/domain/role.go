package domain

import (
	userDomain "golipors/internal/user/domain"
)

type Role struct {
	ID    uint               `gorm:"primaryKey"`
	Name  string             `gorm:"not null;unique"`
	Users []*userDomain.User `gorm:"many2many:user_system_roles;"`
}
