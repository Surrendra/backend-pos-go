package route

import (
	controller "BackendPOS/internal/controller/maintenance"
	"BackendPOS/internal/middleware"
	maintenanceRepository "BackendPOS/internal/repository/maintenance"
	authService "BackendPOS/internal/service/authentication"
	maintenanceService "BackendPOS/internal/service/maintenance"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterMaintenanceRoute(r *gin.RouterGroup, db *gorm.DB, authz *authService.AuthorizationService) {
	merchantRepository := maintenanceRepository.NewMerchantRepository(db)
	merchantService := maintenanceService.NewMerchantService(merchantRepository)
	merchantController := controller.NewMerchantController(merchantService)

	r.Use(middleware.JWTAuth())
	maintenance := r.Group("/maintenance")
	{
		merchant := maintenance.Group("/merchant")
		{
			merchant.GET("datatable", middleware.RequirePermission(authz, "merchant.index"), merchantController.Datatable)
			merchant.POST("create", middleware.RequirePermission(authz, "merchant.create"), merchantController.Create)
			merchant.PUT("update/:code", middleware.RequirePermission(authz, "merchant.update"), merchantController.Update)
			merchant.DELETE("delete/:code", middleware.RequirePermission(authz, "merchant.delete"), merchantController.Delete)
		}
	}

}
