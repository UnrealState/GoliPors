package port

import "golipors/internal/system_role/domain"

type SystemRoleService interface {
	AddSystemRole(role *domain.SystemRole) error
	GetSystemRoleDetails(id uint) (*domain.SystemRole, error)
	ListAllSystemRoles() ([]*domain.SystemRole, error)
}
