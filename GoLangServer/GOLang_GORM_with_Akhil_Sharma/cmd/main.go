package main

import (
	"github.com/apoorvkrishna22/golang_gorm_with_akhil_sharma/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.ResiterBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}
