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
