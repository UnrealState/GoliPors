package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	userService "golipors/internal/user"
	"golipors/internal/user/domain"
	"golipors/internal/user/port"
	"golipors/pkg/adapters/storage/mapper"
	"golipors/pkg/adapters/storage/migrations"
	"golipors/pkg/adapters/storage/types"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	secret string
}

func NewUserRepo(db *gorm.DB, secret string) port.Repo {
	return &userRepo{db, secret}
}

type BcryptHasher struct{}

func (b *BcryptHasher) hashPassword(password string) (string, error) {
	if len(password) > 72 {
		return "", userService.ErrPasswordTooLong
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (b *BcryptHasher) validate(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func (r *userRepo) FindByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error) {
	var user types.User

	// Retrieve the user by username
	err := r.db.WithContext(ctx).Where("email = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Validate the plain password against the hashed password
	bcryptHasher := BcryptHasher{}
	if !bcryptHasher.validate(user.Password, password) {
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

func (r *userRepo) FindByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	return &user, nil
}

func (r *userRepo) Update(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
