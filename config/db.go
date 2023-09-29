package config

import (
	"prakerja/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {

	dsn := "root@tcp(127.0.0.1:3306)/perpustakaan?charset=utf8mb4&parseTime=True&loc=Local"
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
	initMigrate()
}

func initMigrate() {
	database.AutoMigrate(&model.Book{})
}

func DB() *gorm.DB {
	return database
}
