package main

import (
	"BackendPOS/internal/config"
	"BackendPOS/internal/database"
	"BackendPOS/internal/route"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	database.AutoMigrate(db)

	r := gin.Default()
	//if err := utils.InitUniPDF(cfg); err != nil {
	//	logrus.Errorf("Failed to initialize UniPDF: %v", err)
	//}
	whiteListUrl := []string{"http://localhost:3000", "http://127.0.0.1:3000", cfg.FrontendURL}
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
	route.RegisterMaintenanceRoute(api, db)
	r.Run(":" + cfg.AppPort)
}
