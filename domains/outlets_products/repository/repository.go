package repository

import (
	"github.com/hrz8/go-pos-mini/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		GetPricesProductId(trx *gorm.DB, payload *[]uint64) (*[]models.OutletsProducts, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) GetPricesProductId(trx *gorm.DB, payload *[]uint64) (*[]models.OutletsProducts, error) {
	if trx == nil {
		trx = i.db
	}

	result := []models.OutletsProducts{}

	if err := trx.Debug().Where("product_id IN ?", *payload).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &impl{
		db: db,
	}
}
