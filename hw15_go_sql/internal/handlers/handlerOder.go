package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/model"
	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/repository"
	"github.com/gorilla/mux"
)

type OrderHandler struct {
	Repo repository.Order
}

func NewOrderHandler(repo repository.Order) *OrderHandler {
	return &OrderHandler{
		Repo: repo,
	}
}

func (oh *OrderHandler) InsertOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid input order", http.StatusBadRequest)
		return
	}

	id, err := oh.Repo.InsertOrder(r.Context(), &order)
	if err != nil {
		http.Error(w, "Failed to insert order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func (oh *OrderHandler) GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	userIDStr := vars["user_ID"]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	orders, err := oh.Repo.GetOrdersByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to get orders", http.StatusInternalServerError)
		return
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, "Failed to encode orders", http.StatusInternalServerError)
	}
}

func (oh *OrderHandler) UpdateOrderByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	orderIDStr := vars["id"]

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid input order", http.StatusBadRequest)
		return
	}

	order.ID = orderID
	if err := oh.Repo.UpdateOrderByID(r.Context(), &order); err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
