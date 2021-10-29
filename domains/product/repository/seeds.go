package repository

import (
	Config "github.com/hrz8/go-pos-mini/config"
	"github.com/hrz8/go-pos-mini/models"
	"gorm.io/gorm"
)

func RunSeed(db *gorm.DB, appConfig *Config.AppConfig) {
	products := []models.Product{
		{ID: 123, Name: "Product A"},
		{ID: 124, Name: "Product B"},
		{ID: 125, Name: "Product C"},
		{ID: 126, Name: "Product D"},
		{ID: 127, Name: "Product E"},
	}
	db.Debug().Create(&products)
}
