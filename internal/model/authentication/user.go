package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64         `gorm:"primarykey"`
	Code         string         `gorm:"column:code;size:50;uniqueIndex" json:"code"`
	RoleId       uint64         `gorm:"column:role_id;index" json:"role_id"`
	MerchantId   uint64         `gorm:"column:merchant_id;index" json:"merchant_id"`
	MerchantCode *string        `gorm:"column:merchant_code;size:50;index" json:"merchant_code"`
	MerchantName *string        `gorm:"column:merchant_name;size:255;index" json:"merchant_name"`
	Role         *Role          `gorm:"foreignKey:RoleId;references:ID" json:"role"`
	Token        *string        `gorm:"-" json:"token"`
	Name         string         `gorm:"column:name;size:255;uniqueIndex" json:"name"`
	Email        string         `gorm:"column:email;size:255;uniqueIndex" json:"email"`
	Username     string         `gorm:"column:username;size:100;uniqueIndex" json:"username"`
	Password     string         `gorm:"column:password;size:255" json:"-"`
	Address      *string        `gorm:"column:address;type:text" json:"address"`
	Phone        *string        `gorm:"column:phone;size:20" json:"phone"`
	Active       string         `gorm:"column:active;size:20" json:"active"`
	Status       string         `gorm:"column:status;size:50;index;type:enum('active','inactive','suspended','deleted')" json:"status"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
