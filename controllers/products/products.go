package products

import (
	"encoding/json"
	"golang-mvc-webapp/models"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    []models.ProductItem `json: data, omitempty"`
}

var ProductModel *models.ProductModel = models.GetProductModel()

func createAction(w http.ResponseWriter, r *http.Request) {
	var item models.ProductItem
	_ = json.NewDecoder(r.Body).Decode(&item)

	response := &Response{
		Success: true,
		Message: "Created Successfully!!",
		Data: []models.ProductItem{},
	}

	if valid, err := item.IsValid(); !valid {
		response.Success = false
		response.Message = err.Error()
	}

	if err := ProductModel.Create(item); err != nil {
		response.Success = false
		response.Message = err.Error()
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

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

	_ = json.NewEncoder(w).Encode(results)
}
