package repository

import (
	"BackendPOS/internal/helper"
	model "BackendPOS/internal/model/authentication"
	"BackendPOS/internal/request"
	"BackendPOS/internal/response"

	"gorm.io/gorm"
)

type RoleRepositoryInterface interface {
	FindByName(name string) (*model.Role, error)
	CreateDirect(userID uint64, roleID uint64, createdUserID uint64, createdUserName string)
	findByRoleCodes(roleCodes []string) ([]model.Role, error)
	GetData() (*gorm.DB, error)
	Create(role model.Role) (model.Role, error)
	Update(roleID uint64, role model.Role) (model.Role, error)
	Delete(roleID uint64) error
	Datatable(req request.DataTableRequest) ([]model.Role, response.DataTableMeta, error)
	FindByID(roleID uint64) (*model.Role, error)
	FindByCode(code string) (*model.Role, error)
	ReAttachPermissions(permissionNames []string, roleID uint64) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindByName(name string) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return &role, err
}

func (r *roleRepository) CreateDirect(userID uint64, roleID uint64, createdUserID uint64, createdUserName string) {
	userRole := &model.UserRole{
		UserID:          userID,
		RoleID:          roleID,
		CreatedUserID:   createdUserID,
		CreatedUserName: createdUserName,
	}
	_ = r.db.Create(userRole).Error
}

func (r *roleRepository) findByRoleCodes(roleCodes []string) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.Where("code IN ?", roleCodes).Find(&roles).Error
	return roles, err
}

func (r *roleRepository) GetData() (*gorm.DB, error) {
	roles := r.db.Model(&model.Role{})
	return roles, nil
}

func (r *roleRepository) Create(role model.Role) (model.Role, error) {
	err := r.db.Create(&role).Error
	return role, err
}

func (r *roleRepository) Update(roleID uint64, role model.Role) (model.Role, error) {
	err := r.db.Where("id = ?", roleID).Updates(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}

func (r *roleRepository) Delete(roleID uint64) error {
	err := r.db.Where("id = ?", roleID).Delete(&model.Role{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Datatable(req request.DataTableRequest) ([]model.Role, response.DataTableMeta, error) {
	baseDB := r.db.Model(&model.Role{})
	return helper.Paginate[model.Role](baseDB, req, []string{"name", "code", "description"})
}

func (r *roleRepository) FindByID(roleID uint64) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("id = ?", roleID).First(&role).Error
	return &role, err
}

func (r *roleRepository) FindByCode(code string) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("code = ?", code).First(&role).Error
	return &role, err
}

func (r *roleRepository) ReAttachPermissions(permissionNames []string, roleID uint64) error {
	var permissionIDs []uint64
	// get permission IDs from names
	err := r.db.Model(&model.Permission{}).Where("name IN ?", permissionNames).Pluck("id", &permissionIDs).Error
	if err != nil {
		return err
	}
	// delete role_permission where not in permissionIDs
	err = r.db.Where("role_id = ? AND permission_id NOT IN ?", roleID, permissionIDs).Delete(&model.RolePermission{}).Error
	if err != nil {
		return err
	}

	// attach new permissions to role
	for _, permissionID := range permissionIDs {
		var rolePermission model.RolePermission
		err := r.db.Where("role_id = ? AND permission_id = ?", roleID, permissionID).First(&rolePermission).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				rolePermission = model.RolePermission{
					RoleId:       roleID,
					PermissionId: permissionID,
				}
				err = r.db.Create(&rolePermission).Error
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
