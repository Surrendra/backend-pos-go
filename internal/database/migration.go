package database

import (
	authModel "BackendPOS/internal/model/authentication"
	maintenanceModel "BackendPOS/internal/model/maintenance"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		authModel.Permission{},
		authModel.Role{},
		authModel.RolePermission{},
		authModel.User{},
		authModel.UserMerchant{},
		authModel.UserRole{},

		maintenanceModel.Merchant{},
	)
	if err != nil {
		logrus.Error("Failed to auto migrate database", err)
	}
	logrus.Info("Database schema migrated")
}
