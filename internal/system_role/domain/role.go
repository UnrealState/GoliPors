package domain

import (
	userDomain "golipors/internal/user/domain"
)

type RoleID uint

type Role struct {
	ID    RoleID
	Name  string
	Users []*userDomain.User
}
