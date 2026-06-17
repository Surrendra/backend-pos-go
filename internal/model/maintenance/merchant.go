package model

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	ID              uint64          `gorm:"primarykey"`
	Code            string          `gorm:"column:code;size:50;uniqueIndex" json:"code"`
	Name            string          `gorm:"column:name;size:255;uniqueIndex" json:"name"`
	Description     *string         `gorm:"column:description;size:255" json:"description"`
	Address         *string         `gorm:"column:address;size:255" json:"address"`
	PicName         *string         `gorm:"column:pic_name;size:255" json:"pic_name"`
	Domain          string          `gorm:"column:domain;size:255;uniqueIndex" json:"domain"`
	Active          string          `gorm:"column:active;size:20;index" json:"active"`
	PicEmail        string          `gorm:"column:pic_email;size:255" json:"pic_email"`
	PicAddress      *string         `gorm:"column:pic_address;size:255" json:"pic_address"`
	PicContact      *string         `gorm:"column:pic_contact;size:255" json:"pic_contact"`
	CreatedUserID   uint64          `gorm:"column:created_user_id" json:"created_user_id"`
	CreatedUserName string          `gorm:"column:created_user_name" json:"created_user_name"`
	CreatedAt       time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
