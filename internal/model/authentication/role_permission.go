package model

import "time"

type RolePermission struct {
	RoleId       uint64 `gorm:"uniqueIndex:idx_role_permission_role_id_permission_id" json:"role_id"`
	PermissionId uint64 `gorm:"uniqueIndex:idx_role_permission_role_id_permission_id" json:"permission_id"`
	//CreatedUserId   int64     `gorm:"index" json:"created_user_id"`
	//CreatedUserName string    `gorm:"size:255" json:"created_user_name"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
	UpdatedAt time.Time `gorm:"index" json:"updated_at"`
}
