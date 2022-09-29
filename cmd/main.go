package main

import (
	"github.com/gorilla/mux"
	"github.com/melinaco4/book-manager/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
}

/*

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/melinaco4/book-manager/pkg/routes"
)
*/
