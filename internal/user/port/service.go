package port

import "golipors/internal/user/domain"

type UserService interface {
	RegisterUser(user *domain.User) error
	AuthenticateUser(email, password string) (*domain.User, error)
	UpdateProfile(user *domain.User) error
	GetUserDetails(id uint) (*domain.User, error)
}
