package route

import (
	maintenanceRepository "BackendPOS/internal/repository/maintenance"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterMaintenanceRoute(r *gin.RouterGroup, db *gorm.DB) {
	merchantRepository := maintenanceRepository.NewMerchantRepository(db)
}
