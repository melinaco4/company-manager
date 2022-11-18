package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//"github.com/jinzhu/gorm/dialects/mysql"
var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:mysqlpw@tcp(127.0.0.1:3306)/companiesdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
