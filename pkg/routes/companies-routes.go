package routes

import (
	"github.com/gorilla/mux"
	"github.com/melinaco4/company-manager/pkg/controllers"
)

var RegisterCompaniesRoutes = func(router *mux.Router) {
	router.HandleFunc("/company/", controllers.CreateCompany).Methods("POST")
	router.HandleFunc("/company/", controllers.GetCompanies).Methods("GET")
	router.HandleFunc("/company/{companyId}", controllers.GetCompanyById).Methods("GET")
	router.HandleFunc("/company/{companyId}", controllers.UpdateCompany).Methods("PUT")
	router.HandleFunc("/company/{companyId}", controllers.DeleteCompany).Methods("DELETE")
}
