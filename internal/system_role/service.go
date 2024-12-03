package system_role

import (
	"golipors/internal/system_role/domain"
	"golipors/internal/system_role/port"
)

type systemRoleService struct {
	repo port.SystemRoleRepository
}

func NewSystemRoleService(repo port.SystemRoleRepository) port.SystemRoleService {
	return &systemRoleService{repo: repo}
}

func (s *systemRoleService) AddSystemRole(role *domain.SystemRole) error {
	return s.repo.CreateSystemRole(role)
}

func (s *systemRoleService) GetSystemRoleDetails(id uint) (*domain.SystemRole, error) {
	return s.repo.GetSystemRoleByID(id)
}

func (s *systemRoleService) ListAllSystemRoles() ([]*domain.SystemRole, error) {
	return s.repo.ListSystemRoles()
}
