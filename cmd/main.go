package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/melinaco4/company-manager/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterCompaniesRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
