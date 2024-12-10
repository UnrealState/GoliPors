package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"golipors/internal/user/domain"
	"golipors/internal/user/port"
	"golipors/pkg/adapters/rbac"
	"gorm.io/gorm"
	"log"
)

var (
	ErrUserOnCreate      = errors.New("error on creating new user")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidPassword   = errors.New("password is invalid")
	ErrPasswordTooLong   = errors.New("password too long")
)

type service struct {
	repo          port.Repo
	casbinAdapter rbac.CasbinAdapter
}

func NewService(repo port.Repo, casbinAdapter rbac.CasbinAdapter) port.Service {
	return &service{
		repo:          repo,
		casbinAdapter: casbinAdapter,
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

func (s *service) CreateUser(ctx context.Context, user *domain.User) (domain.UserID, error) {
	userID, err := s.repo.Insert(ctx, user)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, ErrUserAlreadyExists
		}

		log.Println("error on creating new user : ", err.Error())

		return 0, ErrUserOnCreate
	}

	return userID, nil
}
func (s *service) AssignRole(ctx context.Context, userID domain.UserID, role string) error {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	// Correctly handle AddRoleForUser
	added, err := s.casbinAdapter.Enforcer.AddRoleForUser(user.Email, role)
	if err != nil {
		return fmt.Errorf("failed to assign role: %v", err)
	}

	// Optional: Log or act on the `added` boolean
	if !added {
		fmt.Printf("Role '%s' was already assigned to user '%s'\n", role, user.Email)
	}

	user.Role = role
	return s.repo.Update(ctx, user)
}

func (s *service) GetUserByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return user, nil
}
