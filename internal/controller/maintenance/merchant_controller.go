package controller

import (
	request2 "BackendPOS/internal/request/maintenance"
	"BackendPOS/internal/response"
	maintencnceService "BackendPOS/internal/service/maintenance"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantService *maintencnceService.MerchantService
}

func NewMerchantController(merchantService *maintencnceService.MerchantService) *MerchantController {
	return &MerchantController{
		merchantService: merchantService,
	}
}

func (c *MerchantController) Create(ctx *gin.Context) {
	var req request2.CreateMerchantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, "Invalid request body", err)
		return
	}
	merchant, err := c.merchantService.Create(req)
	if err != nil {
		response.BadRequest(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "Merchant created successfully", merchant)
}
