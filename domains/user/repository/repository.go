package repository

import (
	"github.com/hrz8/go-pos-mini/models"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(trx *gorm.DB, user *models.User) (*models.User, error)
		GetBy(trx *gorm.DB, payload *models.User) (*models.User, error)
	}

	impl struct {
		db *gorm.DB
	}
)

func (i *impl) Create(trx *gorm.DB, user *models.User) (*models.User, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Create(&user).Error; err != nil {
		return nil, err
	}

	user.Password = nil
	return user, nil
}

func (i *impl) GetBy(trx *gorm.DB, payload *models.User) (*models.User, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.User{}
	if err := trx.Debug().Where(payload).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	db.AutoMigrate(&models.User{})
	return &impl{
		db: db,
	}
}
