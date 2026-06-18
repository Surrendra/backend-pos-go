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
	permissionRepository := authRepository.NewPermissionRepository(db)

	userService := service.NewUserService(userRepository, roleRepository, merchantRepository)
	roleService := service.NewRoleService(roleRepository, permissionRepository)

	authController := controller.NewAuthController(userService)
	roleController := authenticationController.NewRoleController(roleService)
	userController := authenticationController.NewUserController(userService)

	authentication := r.Group("/authentication")
	{
		authentication.POST("login", authController.Login)
		authentication.POST("register", authController.Register)

		authentication.Use(middleware.JWTAuth())
		role := authentication.Group("/role")
		{
			role.GET("get_data", middleware.RequirePermission(authz, "role.index"), roleController.GetData)
			role.GET("datatable", middleware.RequirePermission(authz, "role.index"), roleController.Datatable)
			role.PUT("update/:code", middleware.RequirePermission(authz, "role.edit"), roleController.Update)
			role.POST("create", middleware.RequirePermission(authz, "role.create"), roleController.Create)
		}
		user := authentication.Group("/user")
		{
			user.GET("get_data", middleware.RequirePermission(authz, "user.index"), userController.GetData)
			user.GET("datatable", middleware.RequirePermission(authz, "user.index"), userController.Datatable)
			user.POST("create", middleware.RequirePermission(authz, "user.store"), userController.Create)
			user.PUT("update/:code", middleware.RequirePermission(authz, "user.edit"), userController.Update)
			user.DELETE("delete/:code", middleware.RequirePermission(authz, "user.delete"), userController.Delete)
		}
	}

}
