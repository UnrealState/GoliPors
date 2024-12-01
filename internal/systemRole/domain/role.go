package domain

import (
	userDomain "golipors/internal/user/domain"
)

type Role struct {
	ID    uint
	Name  string
	Users []*userDomain.User
}
