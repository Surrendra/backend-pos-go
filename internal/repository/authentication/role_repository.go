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
