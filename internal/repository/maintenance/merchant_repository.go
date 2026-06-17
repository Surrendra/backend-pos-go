package maintenance

import (
	"BackendPOS/internal/helper"
	modelAuth "BackendPOS/internal/model/authentication"
	model "BackendPOS/internal/model/maintenance"
	maintenanceRequest "BackendPOS/internal/request/maintenance"
	"BackendPOS/internal/response"

	"gorm.io/gorm"
)

var merchantSearchColumns = []string{"name", "code", "domain", "pic_email", "pic_name", "description"}

type MerchantRepositoryInterface interface {
	FindByCode(code string) (*model.Merchant, error)
	CreateDirect(userID uint64, MerchantID uint64, createdUserID uint64, createdUserName string)
	Create(merchant *model.Merchant) (*model.Merchant, error)
	Update(MerchantID uint64, merchant model.Merchant) (model.Merchant, error)
	Delete(MerchantID uint64) error
	Datatable(req maintenanceRequest.MerchantDataTableRequest) ([]model.Merchant, response.DataTableMeta, error)
}

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *merchantRepository {
	return &merchantRepository{db}
}

func (r *merchantRepository) FindByCode(code string) (*model.Merchant, error) {
	var merchant model.Merchant
	err := r.db.Where("code = ?", code).First(&merchant).Error
	return &merchant, err
}

func (r *merchantRepository) FindByID(id uint64) (*model.Merchant, error) {
	var merchant model.Merchant
	err := r.db.Where("id = ?", id).First(&merchant).Error
	return &merchant, err
}

func (r *merchantRepository) Create(merchant *model.Merchant) (*model.Merchant, error) {
	err := r.db.Create(merchant).Error
	return merchant, err
}

func (r *merchantRepository) CreateDirect(userID uint64, MerchantID uint64, createdUserID uint64, createdUserName string) {
	userMerchant := &modelAuth.UserMerchant{
		UserID:          userID,
		MerchantID:      MerchantID,
		CreatedUserID:   createdUserID,
		CreatedUserName: createdUserName,
	}
	_ = r.db.Create(userMerchant).Error
}

func (r *merchantRepository) Update(MerchantID uint64, merchant model.Merchant) (model.Merchant, error) {
	r.db.Where("id = ?", MerchantID).Updates(&merchant)
	newMerchant, _ := r.FindByID(MerchantID)
	return *newMerchant, nil
}

func (r *merchantRepository) Delete(MerchantID uint64) error {
	err := r.db.Where("id = ?", MerchantID).Delete(&model.Merchant{}).Error
	return err
}

func (r *merchantRepository) Datatable(req maintenanceRequest.MerchantDataTableRequest) ([]model.Merchant, response.DataTableMeta, error) {
	baseDB := r.db.Model(&model.Merchant{})
	if req.Active != "" {
		baseDB = baseDB.Where("active = ?", req.Active)
	}
	return helper.Paginate[model.Merchant](baseDB, req.DataTableRequest, merchantSearchColumns)
}
