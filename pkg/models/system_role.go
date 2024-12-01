package models

type SystemRole struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"not null;unique"`
	Users []*User `gorm:"many2many:user_system_roles;"`
}
