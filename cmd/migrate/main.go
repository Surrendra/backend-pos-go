package main

import (
	"BackendPOS/internal/config"
	"BackendPOS/internal/database"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	logrus.Info("Running database migration...")
	database.AutoMigrate(db)
	logrus.Info("Migration finished.")
}
