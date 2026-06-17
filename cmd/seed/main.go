package main

import (
	"BackendPOS/internal/cache"
	"BackendPOS/internal/config"
	"BackendPOS/internal/database"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	logrus.Infof("Running database seeder (env: %s)...", cfg.AppEnv)
	database.Seed(db)
	cache.InitRedis(cfg)
	logrus.Info("Seeding finished.")
}
