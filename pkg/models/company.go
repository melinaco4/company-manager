package models

import (
	"github.com/jinzhu/gorm"
	"github.com/melinaco4/company-manager/pkg/config"
)

//"github.com/melinaco4/book-manager/pkg/config"

var db *gorm.DB

type Company struct {
	gorm.Model
	Name              string `gorm""json:"name"`
	Description       string `json:"description"`
	AmountofEmployees string `json:"amount_of_employees"`
	Registered        bool   `json: "registered"`
	Type              string `json: "type"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Company{})
}

func (b *Company) CreateCompany() *Company {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllCompanies() []Company {
	var Companies []Company
	db.Find(&Companies)
	return Companies
}

func GetCompanyById(ID int64) (*Company, *gorm.DB) {
	var getCompany Company
	db := db.Where("ID=?", ID).Find(&getCompany)
	return &getCompany, db
}

func DeleteBook(ID int64) Company {
	var company Company
	db.Where("ID=?", ID).Delete(company)
	return company
}
