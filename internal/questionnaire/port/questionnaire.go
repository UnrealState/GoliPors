package port

type Service interface {
	RunMigrations() error
}
