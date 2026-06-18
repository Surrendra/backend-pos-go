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
	authRepository authRepository.RoleRepositoryInterface
}

func NewRoleService(authRepository authRepository.RoleRepositoryInterface) *RoleService {
	return &RoleService{authRepository: authRepository}
}

func (s *RoleService) GetData(req request.DataTableRequest) ([]model.Role, error) {
	roles, err := s.authRepository.GetData()
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
	checkRoleName, err := s.authRepository.FindByName(req.Name)
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
	role, err = s.authRepository.Create(role)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %v", err)
	}
	return &role, nil
}

func (s *RoleService) Update(roleID uint64, req authRequest.UpdateRoleRequest) (*model.Role, error) {
	var role model.Role
	role.Name = req.Name
	role.Description = req.Description
	role.Active = req.Active
	role, err := s.authRepository.Update(roleID, role)
	if err != nil {
		return nil, fmt.Errorf("failed to update role: %v", err)
	}
	return &role, nil
}

func (s *RoleService) Delete(roleID uint64) error {
	err := s.authRepository.Delete(roleID)
	if err != nil {
		return fmt.Errorf("failed to delete role: %v", err)
	}
	return nil
}

func (s *RoleService) FindByID(roleID uint64) (*model.Role, error) {
	role, err := s.authRepository.FindByID(roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to find role by ID: %v", err)
	}
	return role, nil
}

func (s *RoleService) FindByCode(code string) (*model.Role, error) {
	role, err := s.authRepository.FindByCode(code)
	if err != nil {
		return nil, fmt.Errorf("failed to find role by code: %v", err)
	}
	return role, nil
}

func (s *RoleService) Datatable(req request.DataTableRequest) ([]model.Role, response.DataTableMeta, error) {
	return s.authRepository.Datatable(req)
}
