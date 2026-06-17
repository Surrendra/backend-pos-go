package maintenance

import (
	modelAuth "BackendPOS/internal/model/authentication"
	model "BackendPOS/internal/model/maintenance"

	"gorm.io/gorm"
)

type MerchantRepositoryInterface interface {
	FindByCode(code string) (*model.Merchant, error)
	CreateDirect(userID uint64, MerchantID uint64, createdUserID uint64, createdUserName string)
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
	return merchant, nil
}

func (r *merchantRepository) Dpdate(MerchantID uint64) error {
	err := r.db.Where("id = ?", MerchantID).Delete(&model.Merchant{}).Error
	return err
}
