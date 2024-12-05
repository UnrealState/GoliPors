package app

import (
	"errors"
	"fmt"
	"golipors/config"
	"golipors/internal/user"
	"golipors/pkg/adapters/email"
	"golipors/pkg/adapters/storage"
	"golipors/pkg/cache"
	"golipors/pkg/postgres"
	"gorm.io/gorm"

	userPort "golipors/internal/user/port"
	redisAdapter "golipors/pkg/adapters/cache"
)

type app struct {
	db          *gorm.DB
	redis       cache.Provider
	cfg         config.Config
	userService userPort.Service
	mailService email.Adapter
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) UserService() userPort.Service {
	return a.userService
}

func (a *app) Cache() cache.Provider {
	return a.redis
}

func (a *app) MailService() email.Adapter {
	return a.mailService
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Pass,
		Name:   a.cfg.DB.Name,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setRedis() {
	a.redis = redisAdapter.NewRedisProvider(fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port))
}

func (a *app) setEmailService() {
	a.mailService = email.NewEmailAdapter(a.cfg.SMTP)
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{cfg: cfg}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.setRedis()
	a.setEmailService()

	a.userService = user.NewService(storage.NewUserRepo(a.db, cfg.Server.PasswordSecret))

	if err := a.userService.RunMigrations(); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to run migrations: %v", err))
	}

	return a, nil
}

func MustNewApp(cfg config.Config) App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return a
}
