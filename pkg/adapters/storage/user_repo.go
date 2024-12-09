package storage

import (
	"context"
	"github.com/go-gormigrate/gormigrate/v2"
	userService "golipors/internal/user"
	"golipors/internal/user/domain"
	"golipors/internal/user/port"
	"golipors/pkg/adapters/storage/mapper"
	"golipors/pkg/adapters/storage/migrations"
	"golipors/pkg/adapters/storage/types"
	"golipors/pkg/hash"

	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	secret string
}

func NewUserRepo(db *gorm.DB, secret string) port.Repo {
	return &userRepo{db, secret}
}

func (r *userRepo) FindByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error) {
	var user types.User

	// Retrieve the user by username
	err := r.db.WithContext(ctx).Where("email = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Validate the plain password against the hashed password
	bcryptHasher := hash.NewBcryptHasher()
	if !bcryptHasher.Validate(user.Password, password) {
		return nil, userService.ErrInvalidPassword
	}

	// Map the user to domain and return
	return mapper.ToDomainUser(&user), nil
}

func (r *userRepo) RunMigrations() error {
	migrator := gormigrate.New(r.db, gormigrate.DefaultOptions, migrations.GetUserMigrations())
	return migrator.Migrate()
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user types.User

	// Retrieve the user by username
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return mapper.ToDomainUser(&user), nil
}

func (r *userRepo) Insert(ctx context.Context, user *domain.User) (domain.UserID, error) {
	newU := mapper.ToModelUser(user)

	return domain.UserID(newU.ID), r.db.WithContext(ctx).Create(newU).Error
}
