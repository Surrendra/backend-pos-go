package middleware

import (
	"BackendPOS/internal/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type JwtAuthMiddlewareInterface interface {
	GetAuthUser() *util.UserClaims
}

type jwtAuthMiddleware struct {
	db  *gorm.DB
	ctx *gin.Context
}

func NewJwtAuthMiddleware(db *gorm.DB, ctx *gin.Context) *jwtAuthMiddleware {
	return &jwtAuthMiddleware{
		db:  db,
		ctx: ctx,
	}
}

func (m *jwtAuthMiddleware) GetAuthUser() *util.UserClaims {
	tokenString := strings.Replace(
		m.ctx.GetHeader("Authorization"),
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
		return nil
	}

	claims := token.Claims.(*util.UserClaims)
	return claims
}
