package service

import (
	"BackendPOS/internal/middleware"
	model "BackendPOS/internal/model/maintenance"
	maintenanceRepository "BackendPOS/internal/repository/maintenance"
	request "BackendPOS/internal/request/maintenance"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type MerchantService struct {
	merchantRepository maintenanceRepository.MerchantRepositoryInterface
}

func NewMerchantService(merchantRepository maintenanceRepository.MerchantRepositoryInterface) *MerchantService {
	return &MerchantService{
		merchantRepository: merchantRepository,
	}
}

func (s *MerchantService) Create(req request.CreateMerchantRequest, ctx *gin.Context) (*model.Merchant, error) {
	var merchant model.Merchant
	authCtx := middleware.NewGinAuthContext(ctx)
	merchant.Code = uuid.NewString()
	merchant.Name = req.Name
	merchant.Description = &req.Description
	merchant.Address = &req.Address
	merchant.PicName = &req.PicName
	merchant.Domain = req.Domain
	merchant.Active = req.Active
	merchant.PicEmail = req.PicEmail
	merchant.PicAddress = &req.PicAddress
	merchant.PicContact = &req.PicContact
	var authUser = authCtx.GetAuthUser()
	merchant.CreatedUserID = authUser.UserId
	merchant.CreatedUserName = authUser.Name
	logrus.Info("Merchant Service Create: ", merchant)
	newMerchant, err := s.merchantRepository.Create(&merchant)
	if err != nil {
		return nil, fmt.Errorf("something wrong when creating merchant: %v", err)
	}
	return newMerchant, nil
}
