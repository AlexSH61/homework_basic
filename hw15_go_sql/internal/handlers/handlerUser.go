package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/model"
	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/repository"
)

type UserHandler struct {
	UserRepo repository.User
}

func NewUserHandler(repo repository.User) *UserHandler {
	return &UserHandler{UserRepo: repo}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect method, you need POST", http.StatusBadRequest)
		return
	}
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "incorrect input", http.StatusBadRequest)
		return
	}
	id, err := uh.UserRepo.InsertUser(r.Context(), &user)
	if err != nil {
		http.Error(w, "creation failed", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "incorrect method, you need GET", http.StatusBadRequest)
		return
	}
	idUser := r.URL.Query().Get("id")
	if idUser == "" {
		http.Error(w, "missing user ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idUser)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := uh.UserRepo.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode user", http.StatusInternalServerError)
	}
}
func (uh *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Incorrect method, you need PUT", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.URL.Query().Get("id")
	if userIDStr == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Incorrect user ID", http.StatusBadRequest)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user.ID = userID
	if err := uh.UserRepo.UpdateUserbyID(r.Context(), &user); err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (uh *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "incorrect method, you need DELETE", http.StatusBadRequest)
		return
	}
	idUser := r.URL.Query().Get("id")
	if idUser == "" {
		http.Error(w, "missing user ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idUser)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	err = uh.UserRepo.DeleteUsersById(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
