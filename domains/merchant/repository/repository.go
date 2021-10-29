package repository

import (
	Config "github.com/hrz8/go-pos-mini/config"
	"github.com/hrz8/go-pos-mini/helpers"
	"github.com/hrz8/go-pos-mini/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		CountAll(trx *gorm.DB) (*int64, error)
		Create(trx *gorm.DB, Merchant *models.Merchant) (*models.Merchant, error)
		GetBy(trx *gorm.DB, payload *models.Merchant) (*models.Merchant, error)
		GetAll(trx *gorm.DB, payload *models.MerchantPayloadGetAll) (*[]models.Merchant, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) CountAll(trx *gorm.DB) (*int64, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	var total int64 = 0
	if err := trx.Model(&models.Merchant{}).Count(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (i *impl) Create(trx *gorm.DB, merchant *models.Merchant) (*models.Merchant, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&merchant).Error; err != nil {
		return nil, err
	}

	return merchant, nil
}

func (i *impl) GetBy(trx *gorm.DB, payload *models.Merchant) (*models.Merchant, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Merchant{}
	if err := trx.Debug().Where(payload).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *impl) GetAll(trx *gorm.DB, payload *models.MerchantPayloadGetAll) (*[]models.Merchant, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.Merchant{}
	executor := trx.Debug()

	if payload.Limit != nil {
		executor = executor.Limit(int(*payload.Limit))
	}
	if payload.Limit != nil && payload.Page != nil {
		executor = executor.Offset(helpers.GetOffset(int(*payload.Page), int(*payload.Limit)))
	}

	if err := executor.Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func NewRepository(db *gorm.DB, appConfig *Config.AppConfig) RepositoryInterface {
	db.AutoMigrate(&models.Merchant{})
	db.Debug().Create(&models.Merchant{
		ID:          8912,
		Name:        "Merchant Pamungkas",
		Description: "Merchant di jalan pamungkas",
	})
	return &impl{
		db: db,
	}
}
