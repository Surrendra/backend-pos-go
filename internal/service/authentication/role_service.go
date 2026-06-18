package service

import (
	model "BackendPOS/internal/model/authentication"
	authRepository "BackendPOS/internal/repository/authentication"
	"BackendPOS/internal/request"
	authRequest "BackendPOS/internal/request/authentication"
	"BackendPOS/internal/response"
	"fmt"

	"github.com/google/uuid"
)

type RoleService struct {
	roleRepository       authRepository.RoleRepositoryInterface
	permissionRepository authRepository.PermissionRepositoryInterface
}

func NewRoleService(
	roleRepository authRepository.RoleRepositoryInterface,
	permissionRepository authRepository.PermissionRepositoryInterface,
) *RoleService {
	return &RoleService{
		roleRepository:       roleRepository,
		permissionRepository: permissionRepository,
	}
}

func (s *RoleService) GetData(req request.DataTableRequest) ([]model.Role, error) {
	roles, err := s.roleRepository.GetData()
	if err != nil {
		return nil, fmt.Errorf("failed to get roles: %v", err)
	}
	if req.Limit > 0 {
		roles = roles.Limit(req.Limit)
	}
	if req.Page > 0 {
		offset := (req.Page - 1) * req.Limit
		roles = roles.Offset(offset)
	}
	var result []model.Role
	err = roles.Preload("Permissions").Find(&result).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get roles: %v", err)
	}
	return result, nil
}

func (s *RoleService) Create(req authRequest.CreateRoleRequest) (*model.Role, error) {
	checkRoleName, err := s.roleRepository.FindByName(req.Name)
	if err != nil && err.Error() != "record not found" {
		return nil, fmt.Errorf("something wrong when checking role name: %v", err)
	}
	if checkRoleName.ID != 0 {
		return nil, fmt.Errorf("role with name %s already exists", req.Name)
	}
	var role model.Role
	role.Name = req.Name
	role.Code = uuid.NewString()
	role.Description = req.Description
	role, err = s.roleRepository.Create(role)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %v", err)
	}
	return &role, nil
}

func (s *RoleService) Update(code string, req authRequest.UpdateRoleRequest) (*model.Role, error) {
	findRRole, err := s.roleRepository.FindByCode(code)
	if err != nil {
		return nil, fmt.Errorf("failed to find role by code: %v", err)
	}
	var role model.Role
	role.Name = req.Name
	role.Description = req.Description
	role.Active = req.Active
	role, err = s.roleRepository.Update(findRRole.ID, role)
	if err != nil {
		return nil, fmt.Errorf("failed to update role: %v", err)
	}
	s.roleRepository.ReAttachPermissions(req.Permissions, findRRole.ID)
	return findRRole, nil
}

func (s *RoleService) Delete(code string) error {
	findRole, err := s.roleRepository.FindByCode(code)
	if err != nil {
		return fmt.Errorf("failed to find role by code: %v", err)
	}
	err = s.roleRepository.Delete(findRole.ID)
	if err != nil {
		return fmt.Errorf("failed to delete role: %v", err)
	}
	return nil
}

func (s *RoleService) FindByID(roleID uint64) (*model.Role, error) {
	role, err := s.roleRepository.FindByID(roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to find role by ID: %v", err)
	}
	return role, nil
}

func (s *RoleService) FindByCode(code string) (*model.Role, error) {
	role, err := s.roleRepository.FindByCode(code)
	if err != nil {
		return nil, fmt.Errorf("failed to find role by code: %v", err)
	}
	return role, nil
}

func (s *RoleService) Datatable(req request.DataTableRequest) ([]model.Role, response.DataTableMeta, error) {
	return s.roleRepository.Datatable(req)
}
