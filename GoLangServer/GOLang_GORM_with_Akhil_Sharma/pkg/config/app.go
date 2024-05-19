package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/bookStoreDb?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the connection
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
