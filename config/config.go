package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type (
	SERVICE struct {
		RESTPORT      int    `mapstructure:"RESTPORT"`
		JWTSECRET     string `mapstructure:"JWTSECRET"`
		ADMINPASSWORD string `mapstructure:"ADMINPASSWORD"`
	}
	DATABASE struct {
		HOST     string `mapstructure:"HOST"`
		PORT     int    `mapstructure:"PORT"`
		USER     string `mapstructure:"USER"`
		PASSWORD string `mapstructure:"PASSWORD"`
		NAME     string `mapstructure:"NAME"`
	}
	AppConfig struct {
		SERVICE  SERVICE  `mapstructure:"SERVICE"`
		DATABASE DATABASE `mapstructure:"DATABASE"`
	}
)

var (
	once      sync.Once
	appConfig AppConfig
)

// NewConfig return configurations implementation
func NewConfig() *AppConfig {
	once.Do(func() {
		v := viper.New()

		v.AddConfigPath(".")
		v.SetConfigName("config")
		v.SetConfigType("yml")

		if err := v.ReadInConfig(); err != nil {
			log.Fatal("Failed to read config file")
		}

		v.Unmarshal(&appConfig)
	})
	return &appConfig
}
