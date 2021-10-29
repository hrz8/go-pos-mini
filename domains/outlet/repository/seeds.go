package repository

import (
	Config "github.com/hrz8/go-pos-mini/config"
	"github.com/hrz8/go-pos-mini/models"
	"gorm.io/gorm"
)

func RunSeed(db *gorm.DB, appConfig *Config.AppConfig) {
	merchantInstance := &models.Merchant{
		ID:          8912,
		Name:        "Merchant Pamungkas",
		Description: "Merchant di jalan pamungkas",
	}

	outlets := []models.Outlet{
		{
			ID:          3892,
			Name:        "Outlet Pamungkas A",
			Description: "Outlet di jalan pamungkas 1",
			Merchant:    merchantInstance,
		},
		{
			ID:          3893,
			Name:        "Outlet Pamungkas B",
			Description: "Outlet di jalan pamungkas 2",
			Merchant:    merchantInstance,
		},
		{
			ID:          3894,
			Name:        "Outlet Pamungkas C",
			Description: "Outlet di jalan pamungkas 3",
			Merchant:    merchantInstance,
		},
		{
			ID:          3895,
			Name:        "Outlet Pamungkas D",
			Description: "Outlet di jalan pamungkas 4",
			Merchant:    merchantInstance,
		},
	}
	db.Debug().Create(&outlets)

	outletsIDs := [4]uint64{3892, 3893, 3894, 3895}
	productIDs := [5]uint64{123, 124, 125, 126, 127}
	for _, outletID := range outletsIDs {
		for index, productID := range productIDs {
			m2mLink := models.OutletsProducts{
				OutletID:  outletID,
				ProductID: productID,
				Price:     uint64(((uint64(index) + uint64(1)) * uint64(100)) + outletID),
			}
			db.Debug().Create(&m2mLink)
		}
	}
}
