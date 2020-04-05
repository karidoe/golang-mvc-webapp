package products

import (
	"encoding/json"
	"fmt"
	"golang-mvc-webapp/models"
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	Success bool                 `json:"success"`
	Message string               `json:"message,omitempty"`
	Data    []models.ProductItem `json: data, omitempty"`
}

type ErrorResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
	Errors  map[string]string `json: errors"`
}

var ProductModel *models.ProductModel = models.GetProductModel()

// To insert a new product into the products collection
func createAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	var item models.ProductItem
	var err error

	_ = json.NewDecoder(r.Body).Decode(&item)

	if errors := ProductModel.Validate(item); len(errors) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  errors,
		})
		return
	}

	if err = ProductModel.Create(&item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorResponse{
			Success: false,
			Message: "Could not insert data",
			Errors:  map[string]string{"info": "Hello world!"},
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Response{
		Success: true,
		Message: "Created Successfully",
		Data:    []models.ProductItem{item},
	})
}

//To show all existing products in the products collection
func indexAction(w http.ResponseWriter, r *http.Request) {
	var results []models.ProductItem
	results, err := ProductModel.All()
	response := &Response{}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	if err != nil {
		response.Message = err.Error()
	}

	response.Success = true
	response.Data = results

	json.NewEncoder(w).Encode(results)
}

func getOneAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	fmt.Fprintf(w, "Get Info %#v", mux.Vars(r)["id"])
}

func updateAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	fmt.Fprintf(w, "Update Product %#v", mux.Vars(r)["id"])
}

func deleteAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	fmt.Fprintf(w, "Delete Product ID: %#v", mux.Vars(r)["id"])
}
