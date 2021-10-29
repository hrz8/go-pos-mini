package database

import (
	"fmt"
	"log"

	"github.com/hrz8/go-pos-mini/config"
	MysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	MysqlInterface interface {
		Connect() *gorm.DB
	}

	mysql struct {
		appConfig config.AppConfig
	}
)

func (m *mysql) Connect() *gorm.DB {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		m.appConfig.DATABASE.USER,
		m.appConfig.DATABASE.PASSWORD,
		m.appConfig.DATABASE.HOST,
		m.appConfig.DATABASE.PORT,
		m.appConfig.DATABASE.NAME,
	)
	db, err := gorm.Open(MysqlDriver.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open connection for mysql")
	}
	return db
}

func NewMysql(appConfig *config.AppConfig) MysqlInterface {
	return &mysql{
		appConfig: *appConfig,
	}
}
