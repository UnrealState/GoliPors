package port

type Repo interface {
	RunMigrations() error
}
