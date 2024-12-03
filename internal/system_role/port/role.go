package port

import "golipors/internal/system_role/domain"

type SystemRoleRepository interface {
	CreateSystemRole(role *domain.SystemRole) error
	GetSystemRoleByID(id uint) (*domain.SystemRole, error)
	ListSystemRoles() ([]*domain.SystemRole, error)
}
