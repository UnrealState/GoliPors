package port

import (
	"context"
	"golipors/internal/user/domain"
)

type Service interface {
	GetUserByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	RunMigrations() error
}
