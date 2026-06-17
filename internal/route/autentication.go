package route

import (
	"BackendPOS/internal/controller"
	authRepository "BackendPOS/internal/repository/authentication"
	maintenanceRepository "BackendPOS/internal/repository/maintenance"
	service "BackendPOS/internal/service/authentication"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthenticationRoute(r *gin.RouterGroup, db *gorm.DB) {
	userRepository := authRepository.NewUserRepository(db)
	roleRepository := authRepository.NewRoleRepository(db)
	merchantRepository := maintenanceRepository.NewMerchantRepository(db)
	userService := service.NewUserService(userRepository, roleRepository, merchantRepository)
	authController := controller.NewAuthController(userService)

	authentication := r.Group("/authentication")
	{
		authentication.POST("login", authController.Login)
		authentication.POST("register", authController.Register)
	}
}
