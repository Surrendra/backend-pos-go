package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	UserID          uint64 `gorm:"uniqueIndex:idx_user_role_user_id_role_id"`
	RoleID          uint64 `gorm:"uniqueIndex:idx_user_role_user_id_role_id"`
	CreatedUserID   uint64 `gorm:"column:created_user_id;index" json:"created_user_id"`
	CreatedUserName string `gorm:"column:created_user_name;size:255;index" json:"created_user_name"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
