package domain

type SystemRole struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null;unique"`
}
