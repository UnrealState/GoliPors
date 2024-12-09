package port

import (
	"context"
	"golipors/internal/user/domain"
)

type Repo interface {
	FindByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Insert(ctx context.Context, user *domain.User) (domain.UserID, error)
	RunMigrations() error
}
