package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/melinaco4/company-manager/pkg/models"
	"github.com/melinaco4/company-manager/pkg/utils"
)

var NewCompany models.Company

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	newCompanies := models.GetAllCompanies()
	res, _ := json.Marshal(newCompanies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCompanyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyid"]
	ID, err := strconv.ParseInt(companyId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	companyDetails, _ := models.GetCompanyById(ID)
	res, _ := json.Marshal(companyDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	CreateCompany := &models.Company{}
	utils.ParseBody(r, CreateCompany)
	b := CreateCompany.CreateCompany()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	ID, err := strconv.ParseInt(companyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	company := models.DeleteCompany(ID)
	res, _ := json.Marshal(company)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var updateCompany = &models.Company{}
	utils.ParseBody(r, updateCompany)
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	ID, err := strconv.ParseInt(companyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	companyDetails, db := models.GetCompanyById(ID)
	if updateCompany.Name != "" {
		companyDetails.Name = updateCompany.Name
	}
	if updateCompany.Description != "" {
		companyDetails.Description = updateCompany.Description
	}
	if updateCompany.AmountofEmployees != "" {
		companyDetails.AmountofEmployees = updateCompany.AmountofEmployees
	}
	if !updateCompany.Registered {
		companyDetails.Registered = updateCompany.Registered
	}
	if updateCompany.Type != "" {
		companyDetails.Type = updateCompany.Type
	}
	db.Save(&companyDetails)
	w.Header().Set("Content-type", "pkglication")
}
