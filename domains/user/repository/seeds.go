package repository

import (
	Config "github.com/hrz8/go-pos-mini/config"
	"github.com/hrz8/go-pos-mini/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunSeed(db *gorm.DB, appConfig *Config.AppConfig) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(appConfig.SERVICE.ADMINPASSWORD), bcrypt.DefaultCost)
	hashedPasswordStr := string(hashedPassword)
	db.Debug().Create(&models.User{
		ID:        999,
		Email:     "admin@posmini.com",
		Password:  &hashedPasswordStr,
		FirstName: "Admin",
	})
}
