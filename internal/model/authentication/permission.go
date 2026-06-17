package model

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID          uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentID    *uint64 `gorm:"column:parent_id;index" json:"parent_id"`
	Code        string  `gorm:"column:code;size:50;uniqueIndex" json:"code"`
	Name        string  `gorm:"column:name;size:255;uniqueIndex" json:"name"`
	Description string  `gorm:"column:description;size:255" json:"description"`
	Active      string  `gorm:"column:active;size:20" json:"active"`

	Parent   *Permission  `gorm:"foreignKey:ParentID;references:ID" json:"parent,omitempty"`
	Children []Permission `gorm:"foreignKey:ParentID;references:ID" json:"children,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
