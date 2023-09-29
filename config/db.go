package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {

	dsn := "root@tcp(127.0.0.1:3306)/go_res_api?charset=utf8mb4&parseTime=True&loc=Local"
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}
