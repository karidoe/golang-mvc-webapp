package products

import (
	"github.com/gorilla/mux"
)

func BindRoutes(r *mux.Router)  {
	r.HandleFunc("/products", indexAction).Methods("GET")
	
	sr := r.PathPrefix("/products").Subrouter()

	//Index page
	sr.HandleFunc("", createAction).Methods("POST")
	sr.HandleFunc("/{id}", getOneAction).Methods("GET")
	sr.HandleFunc("/{id}", createAction).Methods("POST")
	sr.HandleFunc("/{id}", updateAction).Methods("PATCH")
	sr.HandleFunc("/{id}", deleteAction).Methods("DELETE")
}