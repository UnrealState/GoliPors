package port

import (
	"context"
	"golipors/internal/user/domain"
)

type Service interface {
	GetUserByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (domain.UserID, error)
	RunMigrations() error
	AssignRole(ctx context.Context, userID domain.UserID, role string) error
	GetUserByID(ctx context.Context, id domain.UserID) (*domain.User, error)
}
