package route

import (
	controller "BackendPOS/internal/controller/maintenance"
	"BackendPOS/internal/middleware"
	maintenanceRepository "BackendPOS/internal/repository/maintenance"
	maintencnceService "BackendPOS/internal/service/maintenance"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterMaintenanceRoute(r *gin.RouterGroup, db *gorm.DB) {
	merchantRepository := maintenanceRepository.NewMerchantRepository(db)
	merchantService := maintencnceService.NewMerchantService(merchantRepository)
	merchantController := controller.NewMerchantController(merchantService)

	r.Use(middleware.JWTAuth())
	maintenance := r.Group("/maintenance")
	{
		merchant := maintenance.Group("/merchant")
		{
			merchant.POST("create", merchantController.Create)
		}
	}

}
