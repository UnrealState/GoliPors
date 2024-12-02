package storage

import (
	"golipors/internal/system_role/domain"
	"golipors/internal/system_role/port"
	"gorm.io/gorm"
)

type systemRoleRepository struct {
	db *gorm.DB
}

func NewSystemRoleRepository(db *gorm.DB) port.SystemRoleRepository {
	return &systemRoleRepository{db: db}
}

func (r *systemRoleRepository) CreateSystemRole(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *systemRoleRepository) GetSystemRoleByID(id uint) (*domain.Role, error) {
	var role domain.Role
	err := r.db.First(&role, id).Error
	return &role, err
}

func (r *systemRoleRepository) ListSystemRoles() ([]*domain.Role, error) {
	var roles []*domain.Role
	err := r.db.Find(&roles).Error
	return roles, err
}
