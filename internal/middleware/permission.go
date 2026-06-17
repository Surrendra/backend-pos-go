package middleware

import (
	"BackendPOS/internal/config"
	auth "BackendPOS/internal/service/authentication"
	"fmt"

	//"github.com/Surrendra/go-kanal-web-service/internal/services/auth"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequirePermissions(authz *auth.AuthorizationService, reqPermissions []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetUint64("user_id")
		cfg := config.Load()
		if cfg.AppDebug {
			logrus.Info("checking permission for user id: ", userID, " permission: ", reqPermissions)
		}

		//ok, err := authz.HasPermission(userID, permission)
		userPermissions, err := authz.GetUserPermissions(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(403, gin.H{
				"message": "forbidden | failed to get user permissions",
			})
			return
		}

		if cfg.AppDebug {
			fmt.Println("User Permissions:", userPermissions)
		}

		if !HasAllPermissions(userPermissions, reqPermissions) {
			ctx.AbortWithStatusJSON(403, gin.H{"message": "User does not have required permission"})
			return
		}

		ctx.Next()
	}
}

func RequirePermission(authz *auth.AuthorizationService, reqPermission string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetUint64("user_id")
		cfg := config.Load()
		if cfg.AppDebug {
			logrus.Info("checking permission for user id: ", userID, " permission: ", reqPermission)
		}

		//ok, err := authz.HasPermission(userID, permission)
		userPermissions, err := authz.GetUserPermissions(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(403, gin.H{
				"message": "forbidden | failed to get user permissions",
			})
			return
		}

		for _, userPermission := range userPermissions {
			if userPermission == reqPermission {
				ctx.Next()
				return
			}
		}
		// fmt.Println("required permission:", reqPermission)
		ctx.AbortWithStatusJSON(403, gin.H{"message": "User does not have required permission " + reqPermission})
	}
}

func HasAllPermissions(userPerms []string, reqPerms []string) bool {
	set := make(map[string]struct{}, len(userPerms))

	for _, p := range userPerms {
		set[p] = struct{}{}
	}

	for _, rp := range reqPerms {
		if _, ok := set[rp]; !ok {
			return false
		}
	}

	return true
}
