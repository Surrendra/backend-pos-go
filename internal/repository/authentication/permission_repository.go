package repository

import (
	model "BackendPOS/internal/model/authentication"

	"gorm.io/gorm"
)

type PermissionRepositoryInterface interface {
	WhereInNames(names []string) ([]model.Permission, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *permissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) WhereInNames(names []string) ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.db.Where("name IN ?", names).Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
