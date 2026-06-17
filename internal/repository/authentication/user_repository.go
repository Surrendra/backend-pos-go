package repository

import (
	"fmt"

	"BackendPOS/internal/model/authentication"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetData() ([]model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByUsernameOrEmail(term string) (*model.User, error)
	GetUserPermissions(userId uint64) ([]model.Permission, error)
	Update(userCode string, user model.User) (model.User, error)
	FindByCode(userCode string) (model.User, error)
	FindById(userId uint64) (model.User, error)
	Delete(userCode string) error
	Create(user model.User) (model.User, error)
	AssignRole(userID uint64, roleID uint64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetData() ([]model.User, error) {
	var users []model.User
	err := r.db.Preload("Role").Select("id, code, organization_id, role_id, name, email, password, address, phone, status, created_at, updated_at, deleted_at").Find(&users).Error
	return users, err
}

func (r *userRepository) FindByUsernameOrEmail(term string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ? OR email = ?", term, term).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	fmt.Println(username)
	err := r.db.Where("email = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) FindById(userID uint64) (model.User, error) {
	var user model.User
	err := r.db.Preload("Role").Preload("Unit").First(&user, userID).Error
	return user, err
}

func (r *userRepository) GetUserPermissions(UserId uint64) ([]model.Permission, error) {
	var permissions []model.Permission
	r.db.Find(&permissions).Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Joins("JOIN user_roles ON user_roles.role_id = role_permissions.role_id").
		Where("user_roles.user_id = ?", UserId)
	return permissions, nil
}

func (r *userRepository) Update(userCode string, user model.User) (model.User, error) {
	fmt.Println("userCode : ", userCode)
	err := r.db.Where("code = ?", userCode).Updates(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByCode(userCode string) (model.User, error) {
	var user model.User
	err := r.db.Preload("Role").Preload("Unit").Preload("Organization").Where("code = ?", userCode).First(&user).Error
	return user, err
}

func (r *userRepository) Delete(userCode string) error {
	err := r.db.Where("code = ?", userCode).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

func (r *userRepository) AssignRole(userID uint64, roleID uint64) error {
	//ctx := gin.Context{}
	//userRole := &model.UserRole{}
	//userRole.UserId = userID
	//userRole.RoleId = roleID
	//userRole.CreatedUserId = ctx.GetUint64("user_id")
	//userRole.CreatedUserName = ctx.GetString("user_name")
	//err := r.db.Create(&userRole).Error
	//if err != nil {
	//	return err
	//}
	return nil
}
