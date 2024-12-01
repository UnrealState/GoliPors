package port

import "golipors/internal/system_role/domain"

type SystemRoleService interface {
	AddSystemRole(role *domain.Role) error
	GetSystemRoleDetails(id uint) (*domain.Role, error)
	ListAllSystemRoles() ([]*domain.Role, error)
}
