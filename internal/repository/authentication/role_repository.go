package repository

import (
	model "BackendPOS/internal/model/authentication"

	"gorm.io/gorm"
)

type RoleRepositoryInterface interface {
	FindByName(name string) (*model.Role, error)
	CreateDirect(userID uint64, roleID uint64, createdUserID uint64, createdUserName string)
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
