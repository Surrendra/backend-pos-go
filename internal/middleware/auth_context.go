package middleware

import (
	"BackendPOS/internal/util"

	"github.com/gin-gonic/gin"
)

const AuthUserContextKey = "auth_user"

type GinAuthContext struct {
	ctx *gin.Context
}

func NewGinAuthContext(ctx *gin.Context) *GinAuthContext {
	return &GinAuthContext{ctx: ctx}
}

func (g *GinAuthContext) GetAuthUser() *util.UserClaims {
	if g == nil || g.ctx == nil {
		return nil
	}
	v, ok := g.ctx.Get(AuthUserContextKey)
	if !ok {
		return nil
	}
	claims, _ := v.(*util.UserClaims)
	return claims
}
