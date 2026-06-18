package service

import (
	"BackendPOS/internal/middleware"
	model "BackendPOS/internal/model/maintenance"
	maintenanceRepository "BackendPOS/internal/repository/maintenance"
	maintenanceRequest "BackendPOS/internal/request/maintenance"
	"BackendPOS/internal/response"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantService struct {
	merchantRepository maintenanceRepository.MerchantRepositoryInterface
}

func NewMerchantService(merchantRepository maintenanceRepository.MerchantRepositoryInterface) *MerchantService {
	return &MerchantService{
		merchantRepository: merchantRepository,
	}
}

func (s *MerchantService) GetData(req maintenanceRequest.MerchantDataTableRequest) ([]model.Merchant, error) {
	merchants, err := s.merchantRepository.GetData()
	if err != nil {
		return nil, fmt.Errorf("failed to get merchants: %v", err)
	}
	if req.Active != "" {
		merchants = merchants.Where("active = ?", req.Active)
	}
	if req.Limit > 0 {
		merchants = merchants.Limit(req.Limit)
	}
	if req.Page > 0 {
		offset := (req.Page - 1) * req.Limit
		merchants = merchants.Offset(offset)
	}
	var result []model.Merchant
	err = merchants.Find(&result).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get merchants: %v", err)
	}
	return result, nil
}

func (s *MerchantService) Create(req maintenanceRequest.CreateMerchantRequest, ctx *gin.Context) (*model.Merchant, error) {
	checkMerchantDomain, err := s.merchantRepository.FindByDomain(req.Domain)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("something wrong when checking merchant domain: %v", err)
	}
	if checkMerchantDomain != nil {
		return nil, fmt.Errorf("merchant with domain %s already exists", req.Domain)
	}
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
	// logrus.Info("Merchant Service Create: ", merchant)
	newMerchant, err := s.merchantRepository.Create(&merchant)
	if err != nil {
		return nil, fmt.Errorf("something wrong when creating merchant: %v", err)
	}
	return newMerchant, nil
}

func (s *MerchantService) Update(req maintenanceRequest.UpdateMerchantRequest, ctx *gin.Context) (*model.Merchant, error) {
	var merchant model.Merchant
	merchant.Name = req.Name
	merchant.Description = &req.Description
	merchant.Address = &req.Address
	merchant.PicName = &req.PicName
	merchant.Domain = req.Domain
	merchant.Active = req.Active
	merchant.PicEmail = req.PicEmail
	merchant.PicAddress = &req.PicAddress
	merchant.PicContact = &req.PicContact
	var merchantCode = ctx.Param("code")
	merchantFind, err := s.merchantRepository.FindByCode(merchantCode)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("merchant with code %s not found", merchantCode)
		}
		return nil, fmt.Errorf("something wrong when finding merchant: %v", err)
	}
	newMerchant, err := s.merchantRepository.Update(merchantFind.ID, merchant)
	if err != nil {
		return nil, fmt.Errorf("something wrong when updating merchant: %v", err)
	}
	return &newMerchant, nil
}

func (s *MerchantService) Delete(ctx *gin.Context) error {
	var merchantCode = ctx.Param("code")
	merchantFind, err := s.merchantRepository.FindByCode(merchantCode)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("merchant with code %s not found", merchantCode)
		}
		return fmt.Errorf("something wrong when finding merchant: %v", err)
	}
	err = s.merchantRepository.Delete(merchantFind.ID)
	if err != nil {
		return fmt.Errorf("something wrong when deleting merchant: %v", err)
	}
	return nil
}

func (s *MerchantService) Datatable(req maintenanceRequest.MerchantDataTableRequest) ([]model.Merchant, response.DataTableMeta, error) {
	return s.merchantRepository.Datatable(req)
}
