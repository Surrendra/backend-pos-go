package route

import (
	controller "BackendPOS/internal/controller"
	authenticationController "BackendPOS/internal/controller/authentication"
	"BackendPOS/internal/middleware"
	authRepository "BackendPOS/internal/repository/authentication"
	maintenanceRepository "BackendPOS/internal/repository/maintenance"
	authService "BackendPOS/internal/service/authentication"
	service "BackendPOS/internal/service/authentication"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthenticationRoute(r *gin.RouterGroup, db *gorm.DB, authz *authService.AuthorizationService) {
	userRepository := authRepository.NewUserRepository(db)
	roleRepository := authRepository.NewRoleRepository(db)
	merchantRepository := maintenanceRepository.NewMerchantRepository(db)
	userService := service.NewUserService(userRepository, roleRepository, merchantRepository)
	roleService := service.NewRoleService(roleRepository)

	authController := controller.NewAuthController(userService)
	roleController := authenticationController.NewRoleController(roleService)

	authentication := r.Group("/authentication")
	{
		authentication.POST("login", authController.Login)
		authentication.POST("register", authController.Register)

		authentication.Use(middleware.JWTAuth())
		role := authentication.Group("/role")
		{
			role.GET("get_data", middleware.RequirePermission(authz, "role.index"), roleController.GetData)
			role.GET("datatable", middleware.RequirePermission(authz, "role.index"), roleController.Datatable)
		}
	}

}
