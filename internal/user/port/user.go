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
	FindByID(ctx context.Context, id domain.UserID) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
}
