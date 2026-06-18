package controller

import (
	request "BackendPOS/internal/request"
	authRequest "BackendPOS/internal/request/authentication"
	"BackendPOS/internal/response"
	authService "BackendPOS/internal/service/authentication"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RoleController struct {
	roleService *authService.RoleService
}

func NewRoleController(roleService *authService.RoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (c *RoleController) GetData(ctx *gin.Context) {
	var req request.DataTableRequest
	roles, err := c.roleService.GetData(req)
	if err != nil {
		response.BadRequest(ctx, "failed binding data", err)
		return
	}
	response.Success(ctx, "Data retrieved successfully", roles)
}

func (c *RoleController) Datatable(ctx *gin.Context) {
	var req request.DataTableRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.BadRequest(ctx, "Invalid query parameters", err.Error())
		return
	}
	req.SetDefaults()
	roles, meta, err := c.roleService.Datatable(req)
	if err != nil {
		response.Error(ctx, 500, "Failed to retrieve data", err.Error())
		return
	}
	response.SuccessDatatable(ctx, "Data retrieved successfully", roles, meta)
}

func (c *RoleController) Create(ctx *gin.Context) {
	var req authRequest.CreateRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "Invalid request body", err.Error())
		return
	}
	role, err := c.roleService.Create(req)
	if err != nil {
		response.Error(ctx, 500, "Failed to create role", err.Error())
		return
	}
	response.Success(ctx, "Role created successfully", role)
}

func (c *RoleController) Update(ctx *gin.Context) {
	var req authRequest.UpdateRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "Invalid request body", err.Error())
		return
	}
	logrus.Info(req)
	role, err := c.roleService.Update(ctx.Param("code"), req)
	if err != nil {
		response.Error(ctx, 500, "Failed to update role", err.Error())
		return
	}
	response.Success(ctx, "Role updated successfully", role)
}
func (c *RoleController) Delete(ctx *gin.Context) {
	err := c.roleService.Delete(ctx.Param("code"))
	if err != nil {
		response.Error(ctx, 500, "Failed to delete role", err.Error())
		return
	}
	response.Success(ctx, "Role deleted successfully", nil)
}
