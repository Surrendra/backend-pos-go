package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint64         `gorm:"primarykey"`
	Code        string         `gorm:"column:code;size:50;uniqueIndex" json:"code"`
	Name        string         `gorm:"column:name;size:255;uniqueIndex" json:"name"`
	Description *string        `gorm:"column:description;size:255" json:"description"`
	Active      string         `gorm:"column:active;size:20" json:"active"`
	Permissions []Permission   `gorm:"many2many:role_permissions;joinForeignKey:RoleId;JoinReferences:PermissionId" json:"permissions"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
