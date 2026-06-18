package controller

import (
	baseRequest "BackendPOS/internal/request"
	authRequest "BackendPOS/internal/request/authentication"
	"BackendPOS/internal/response"
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

func (c *UserController) GetData(ctx *gin.Context) {
	users, err := c.UserService.GetData()
	if err != nil {
		response.Error(ctx, 500, "Failed to retrieve data", err.Error())
		return
	}
	response.Success(ctx, "Data retrieved successfully", users)
}

func (c *UserController) Datatable(ctx *gin.Context) {
	var req baseRequest.DataTableRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.BadRequest(ctx, "Invalid query parameters", err.Error())
		return
	}
	req.SetDefaults()
	users, meta, err := c.UserService.Datatable(req)
	if err != nil {
		response.Error(ctx, 500, "Failed to retrieve data", err.Error())
		return
	}
	response.SuccessDatatable(ctx, "Data retrieved successfully", users, meta)
}

func (c *UserController) Create(ctx *gin.Context) {
	var req authRequest.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "Invalid request body", err.Error())
		return
	}
	user, err := c.UserService.Create(req)
	if err != nil {
		response.Error(ctx, 500, "Failed to create user", err.Error())
		return
	}
	response.Success(ctx, "User created successfully", user)
}

func (c *UserController) Update(ctx *gin.Context) {
	var req authRequest.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "Invalid request body", err.Error())
		return
	}
	user, err := c.UserService.Update(ctx.Param("code"), req)
	if err != nil {
		response.Error(ctx, 500, "Failed to update user", err.Error())
		return
	}
	response.Success(ctx, "User updated successfully", user)
}

func (c *UserController) Delete(ctx *gin.Context) {
	err := c.UserService.Delete(ctx.Param("code"))
	if err != nil {
		response.Error(ctx, 500, "Failed to delete user", err.Error())
		return
	}
	response.Success(ctx, "User deleted successfully", nil)
}
