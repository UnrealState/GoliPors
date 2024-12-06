package user

import (
	"context"
	"errors"
	"fmt"
	"golipors/internal/user/domain"
	"golipors/internal/user/port"
	"gorm.io/gorm"
)

var (
	ErrUserOnCreate           = errors.New("error on creating new user")
	ErrUserCreationValidation = errors.New("validation failed while creating new user")
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidPassword        = errors.New("password is invalid")
	ErrPasswordTooLong        = errors.New("password too long")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUserByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error) {
	user, err := s.repo.FindByUsernamePassword(ctx, username, password)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, ErrInvalidPassword):
			return nil, ErrUserNotFound
		default:
			return nil, errors.New(fmt.Sprintf("failed to authenticate user: %s", err))
		}
	}

	return user, nil
}
func (s *service) RunMigrations() error {
	return s.repo.RunMigrations()
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, ErrInvalidPassword):
			return nil, ErrUserNotFound
		default:
			return nil, errors.New(fmt.Sprintf("failed to authenticate user: %s", err))
		}
	}

	return user, nil
}
