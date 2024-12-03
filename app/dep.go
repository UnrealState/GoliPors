package app

import (
	"golipors/config"
	"gorm.io/gorm"
)

type app struct {
	db  *gorm.DB
	cfg config.Config
	// ToDo define services
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	// ToDo Initialize db connection
	/*db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
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

	a.db = db*/
	return nil
}

func (a *app) setRedis() {
	// ToDo Initialize redis connection
	/*a.redisCache = redisAdapter.NewRedisProvider(fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port))*/
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.setRedis()

	return a, nil
}

func MustNewApp(cfg config.Config) App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return a
}