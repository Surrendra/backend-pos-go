package main

import (
	"BackendPOS/internal/config"
	"BackendPOS/internal/database"
	"BackendPOS/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	database.AutoMigrate(db)

	r := gin.Default()
	api := r.Group("/api")
	route.RegisterAuthenticationRoute(api, db)
	r.Run(":" + cfg.AppPort)
}
