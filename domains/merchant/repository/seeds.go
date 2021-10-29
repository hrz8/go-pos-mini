package repository

import (
	Config "github.com/hrz8/go-pos-mini/config"
	"github.com/hrz8/go-pos-mini/models"
	"gorm.io/gorm"
)

func RunSeed(db *gorm.DB, appConfig *Config.AppConfig) {
	db.Debug().Create(&models.Merchant{
		ID:          8912,
		Name:        "Merchant Pamungkas",
		Description: "Merchant di jalan pamungkas",
	})
}
