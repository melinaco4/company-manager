package models

import (
	"github/jinzhu/gorm"

	"github.com/jinzhu/gorm"
	"github.com/melinaco4/book-manager/config"
	"github.com/melinaco4/book-manager/pkg/config"
)

//"github.com/melinaco4/book-manager/pkg/config"

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
