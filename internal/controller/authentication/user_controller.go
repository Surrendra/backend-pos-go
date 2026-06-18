package controller

import (
	authRequest "BackendPOS/internal/request/authentication"
	authService "BackendPOS/internal/service/authentication"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *authService.UserService
}

func NewUserController(userService *authService.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) Create(ctx *gin.Context) {
	var req authRequest.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// user, err := c.UserService.crea(req)
	// if err != nil {
	// 	ctx.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// ctx.JSON(201, user)
}
