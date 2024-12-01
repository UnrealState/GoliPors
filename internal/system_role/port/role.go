package port

import "golipors/internal/system_role/domain"

type SystemRoleRepository interface {
	CreateSystemRole(role *domain.Role) error
	GetSystemRoleByID(id uint) (*domain.Role, error)
	ListSystemRoles() ([]*domain.Role, error)
}
