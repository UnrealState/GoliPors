package user

import (
	"golipors/internal/user/domain"
	"golipors/internal/user/port"
)

type userService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(user *domain.User) error {
	//TODO Implement registration logic, e.g., hash password, validate input
	return s.repo.CreateUser(user)
}

func (s *userService) AuthenticateUser(email, password string) (*domain.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	//if user.Password != hashPassword(password) { // TODO Replace with real hashing
	//	return nil, errors.New("invalid credentials")
	//}

	return user, nil
}

func (s *userService) UpdateProfile(user *domain.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) GetUserDetails(id uint) (*domain.User, error) {
	return s.repo.GetUserByID(id)
}
