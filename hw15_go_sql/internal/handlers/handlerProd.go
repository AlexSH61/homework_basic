package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/model"
	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/repository"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	RepoProduct repository.Product
}

func NewProductHandler(repoProd repository.Product) *ProductHandler {
	return &ProductHandler{
		RepoProduct: repoProd,
	}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed, use POST", http.StatusMethodNotAllowed)
		return
	}

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := ph.RepoProduct.InsertProduct(r.Context(), &product)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed, use PUT", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	productIDStr := vars["id"]
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := ph.RepoProduct.UpdateProductbyID(r.Context(), productID); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed, use DELETE", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	productIDStr := vars["id"]
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = ph.RepoProduct.DeleteProductbyID(r.Context(), productID)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
