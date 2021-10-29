package repository

import (
	"github.com/hrz8/go-pos-mini/helpers"
	"github.com/hrz8/go-pos-mini/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		CountAll(trx *gorm.DB) (*int64, error)
		Create(trx *gorm.DB, product *models.Product) (*models.Product, error)
		GetBy(trx *gorm.DB, payload *models.Product) (*models.Product, error)
		Update(trx *gorm.DB, productInstance *models.Product, payload *models.ProductPayloadUpdate) (*models.Product, error)
		DeleteById(trx *gorm.DB, id uint64) error
		GetAll(trx *gorm.DB, payload *models.ProductPayloadGetAll) (*[]models.Product, error)
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
	if err := trx.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (i *impl) Create(trx *gorm.DB, product *models.Product) (*models.Product, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (i *impl) GetBy(trx *gorm.DB, payload *models.Product) (*models.Product, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Product{}
	if err := trx.Debug().Where(payload).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *impl) Update(trx *gorm.DB, productInstance *models.Product, payload *models.ProductPayloadUpdate) (*models.Product, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Model(productInstance).Updates(models.Product{
		Name:        helpers.Ternary(payload.Name, productInstance.Name).(string),
		Description: helpers.Ternary(payload.Description, productInstance.Description).(string),
	}).Error; err != nil {
		return nil, err
	}

	return productInstance, nil
}

func (i *impl) DeleteById(trx *gorm.DB, id uint64) error {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.Product{}
	if err := trx.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *impl) GetAll(trx *gorm.DB, payload *models.ProductPayloadGetAll) (*[]models.Product, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := []models.Product{}
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

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.Product{})
	return &impl{
		db: db,
	}
}
