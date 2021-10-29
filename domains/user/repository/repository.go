package repository

import (
	Config "github.com/hrz8/go-pos-mini/config"
	"github.com/hrz8/go-pos-mini/helpers"
	"github.com/hrz8/go-pos-mini/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	RepositoryInterface interface {
		Create(trx *gorm.DB, user *models.User) (*models.User, error)
		GetBy(trx *gorm.DB, payload *models.User) (*models.User, error)
		Update(trx *gorm.DB, userInstance *models.User, payload *models.UserPayloadUpdate) (*models.User, error)
		DeleteById(trx *gorm.DB, id uint64) error
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

func (i *impl) Update(trx *gorm.DB, userInstance *models.User, payload *models.UserPayloadUpdate) (*models.User, error) {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	if err := trx.Debug().Model(userInstance).Updates(models.User{
		Password:  helpers.Ternary(payload.Password, userInstance.Password).(*string),
		FirstName: helpers.Ternary(payload.FirstName, userInstance.FirstName).(string),
		LastName:  helpers.Ternary(payload.LastName, userInstance.LastName).(*string),
	}).Error; err != nil {
		return nil, err
	}

	userInstance.Password = nil
	return userInstance, nil
}

func (i *impl) DeleteById(trx *gorm.DB, id uint64) error {
	// transaction check
	if trx == nil {
		trx = i.db
	}

	// execution
	result := models.User{}
	if err := trx.Debug().Delete(&result, id).Error; err != nil {
		return err
	}
	return nil
}

func NewRepository(db *gorm.DB, appConfig *Config.AppConfig) RepositoryInterface {
	db.AutoMigrate(&models.User{})
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(appConfig.SERVICE.ADMINPASSWORD), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)
	db.Debug().Create(&models.User{
		ID:        999,
		Email:     "admin@posmini.com",
		Password:  &hashedPasswordStr,
		FirstName: "Admin",
	})
	return &impl{
		db: db,
	}
}
