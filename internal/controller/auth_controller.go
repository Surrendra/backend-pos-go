package controller

import (
	request "BackendPOS/internal/request/authentication"
	"BackendPOS/internal/response"
	service "BackendPOS/internal/service/authentication"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService *service.UserService
}

func NewAuthController(userService *service.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, "Invalid request body", err)
		return
	}
	user, err := c.userService.Login(req.Username, req.Password)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "login successful", user)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, "Invalid request body", err)
		return
	}
	user, err := c.userService.Register(req)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "Register Success!", user)
}
