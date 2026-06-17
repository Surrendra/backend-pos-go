package middleware

import (
	"BackendPOS/internal/config"
	"BackendPOS/internal/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var cfg = config.Load()
var jwtSecret = []byte(cfg.JwtSecret)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.Replace(
			ctx.GetHeader("Authorization"),
			"Bearer ",
			"",
			1,
		)

		token, err := jwt.ParseWithClaims(
			tokenString,
			&util.UserClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			},
		)

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		// claims := token.Claims.(*util.UserClaims)

		claims, err := util.GetUserClaimsFromAuthHeader(ctx.GetHeader("Authorization"))
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"message": err.Error()})
			return
		}
		ctx.Set("user_id", claims.UserId)
		ctx.Set("name", claims.Name)
		ctx.Set("code", claims.Code)

		ctx.Set(AuthUserContextKey, claims)

		ctx.Next()
	}
}
