package port

import (
	"context"
	"golipors/internal/user/domain"
)

type Repo interface {
	FindByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error)
	RunMigrations() error
}
