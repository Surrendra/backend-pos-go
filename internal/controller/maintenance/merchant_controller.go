package controller

import (
	maintenanceRequest "BackendPOS/internal/request/maintenance"
	"BackendPOS/internal/response"
	maintenanceService "BackendPOS/internal/service/maintenance"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantService *maintenanceService.MerchantService
}

func NewMerchantController(merchantService *maintenanceService.MerchantService) *MerchantController {
	return &MerchantController{
		merchantService: merchantService,
	}
}

func (c *MerchantController) GetData(ctx *gin.Context) {
	var req maintenanceRequest.MerchantDataTableRequest
	var err = ctx.ShouldBindQuery(&req)
	if err != nil {
		response.BadRequest(ctx, "Invalid query parameters", err.Error())
		return
	}
	req.SetDefaults()
	merchants, err := c.merchantService.GetData(req)
	if err != nil {
		response.Error(ctx, 500, "Failed to retrieve data", err.Error())
		return
	}
	response.Success(ctx, "Data retrieved successfully", merchants)
}

func (c *MerchantController) Datatable(ctx *gin.Context) {
	var req maintenanceRequest.MerchantDataTableRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.BadRequest(ctx, "Invalid query parameters", err.Error())
		return
	}
	req.SetDefaults()
	merchants, meta, err := c.merchantService.Datatable(req)
	if err != nil {
		response.Error(ctx, 500, "Failed to retrieve data", err.Error())
		return
	}
	response.SuccessDatatable(ctx, "Data retrieved successfully", merchants, meta)
}

func (c *MerchantController) Create(ctx *gin.Context) {
	var req maintenanceRequest.CreateMerchantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, "Invalid request body", err)
		return
	}
	merchant, err := c.merchantService.Create(req, ctx)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "Merchant created successfully", merchant)
}

func (c *MerchantController) Update(ctx *gin.Context) {
	var req maintenanceRequest.UpdateMerchantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, "Invalid request body", err)
		return
	}
	merchant, err := c.merchantService.Update(req, ctx)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "Merchant updated successfully", merchant)
}

func (c *MerchantController) Delete(ctx *gin.Context) {
	err := c.merchantService.Delete(ctx)
	if err != nil {
		response.ValidationError(ctx, "Invalid merchant ID", err)
		return
	}
	response.Success(ctx, "Merchant deleted successfully", nil)
}
