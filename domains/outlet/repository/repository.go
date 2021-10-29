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
		Create(trx *gorm.DB, Outlet *models.Outlet) (*models.Outlet, error)
		GetBy(trx *gorm.DB, payload *models.Outlet) (*models.Outlet, error)
		GetAll(trx *gorm.DB, payload *models.OutletPayloadGetAll) (*[]models.Outlet, error)
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
	if err := trx.Model(&models.Outlet{}).Count(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (i *impl) Create(trx *gorm.DB, merchant *models.Outlet) (*models.Outlet, error) {
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

func (i *impl) GetBy(trx *gorm.DB, payload *models.Outlet) (*models.Outlet, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Outlet{}
	if err := trx.Debug().Where(payload).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *impl) GetAll(trx *gorm.DB, payload *models.OutletPayloadGetAll) (*[]models.Outlet, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.Outlet{}
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
	db.AutoMigrate(&models.Outlet{})
	db.Debug().Create(&models.Outlet{
		ID:          3892,
		Name:        "Outlet Pamungkas A",
		Description: "Outlet di jalan pamungkas 1",
		Merchant: &models.Merchant{
			ID:          8912,
			Name:        "Merchant Pamungkas",
			Description: "Merchant di jalan pamungkas",
		},
	})
	db.Debug().Create(&models.Outlet{
		ID:          3892,
		Name:        "Outlet Pamungkas B",
		Description: "Outlet di jalan pamungkas 2",
		Merchant: &models.Merchant{
			ID:          8912,
			Name:        "Merchant Pamungkas",
			Description: "Merchant di jalan pamungkas",
		},
	})
	db.Debug().Create(&models.Outlet{
		ID:          3892,
		Name:        "Outlet Pamungkas C",
		Description: "Outlet di jalan pamungkas 3",
		Merchant: &models.Merchant{
			ID:          8912,
			Name:        "Merchant Pamungkas",
			Description: "Merchant di jalan pamungkas",
		},
	})
	db.Debug().Create(&models.Outlet{
		ID:          3892,
		Name:        "Outlet Pamungkas D",
		Description: "Outlet di jalan pamungkas 4",
		Merchant: &models.Merchant{
			ID:          8912,
			Name:        "Merchant Pamungkas",
			Description: "Merchant di jalan pamungkas",
		},
	})
	return &impl{
		db: db,
	}
}
