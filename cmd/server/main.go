package main

import (
	"BackendPOS/internal/cache"
	"BackendPOS/internal/config"
	constan "BackendPOS/internal/constant"
	"BackendPOS/internal/database"
	"BackendPOS/internal/route"
	"time"

	authService "BackendPOS/internal/service/authentication"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)
	cache.InitRedis(cfg)

	database.AutoMigrate(db)
	authz := authService.NewAuthorizationService(db)

	r := gin.Default()
	//if err := utils.InitUniPDF(cfg); err != nil {
	//	logrus.Errorf("Failed to initialize UniPDF: %v", err)
	//}
	whiteListUrl := []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	if cfg.FrontendURL != "" {
		whiteListUrl = append(whiteListUrl, cfg.FrontendURL)
	}
	if cfg.AppEnv == constan.EnvLocal || cfg.AppEnv == constan.EnvDevelopment {
		authz.DeleteUserPermissionCache(1)
		logrus.Info("Successfully DeleteUserPermissionCache for user ID 1")
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     whiteListUrl,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	api := r.Group("/api")
	route.RegisterAuthenticationRoute(api, db)
	route.RegisterMaintenanceRoute(api, db, authz)
	r.Run(":" + cfg.AppPort)
}
