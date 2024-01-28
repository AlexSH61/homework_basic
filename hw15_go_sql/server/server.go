package server

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/AlexSH61/homework_basic/hw15_go_sql/database"
)

type Handler struct {
	database *db.DataBase
}

func NewHandler(database *db.DataBase) *Handler {
	return &Handler{database: database}
}

func StartServer(address, port string) {
	h := NewHandler(nil)
	http.HandleFunc("/example", h.handlerRequest)
	serverAddress := fmt.Sprintf("%s:%s", address, port)
	fmt.Printf("Server listening on:  %s\n", serverAddress)
	server := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func (h *Handler) HandlerPost(w http.ResponseWriter, r *http.Request) {
	tx, err := h.database.BeginTransaction()
	if err != nil {
		http.Error(w, "error start transation", http.StatusInternalServerError)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
			http.Error(w, "error sql", http.StatusInternalServerError)
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Exec("INSERT INTO schema_store.users (name, email, password) VALUES ('sasha','ninja@gmail.com','1234')")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sql query : %v", err), http.StatusInternalServerError)
		return
	}
	_, err = tx.Exec("UPDATE schema_store.Users SET name = 'pasha' WHERE id = 1")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sql query : %v", err), http.StatusInternalServerError)
		return
	}
	_, err = tx.Exec("DELETE FROM schema_store.users WHERE id = 1")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sql query : %v", err), http.StatusInternalServerError)
		return
	}
	_, err = tx.Exec("INSERT INTO schema_store.products (name, price) VALUES ('book1', 19.99)")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sql query : %v", err), http.StatusInternalServerError)
		return
	}
	_, err = tx.Exec("INSERT INTO schema_store.products (name, price) VALUES ('book1', 19.99)")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sql query : %v", err), http.StatusInternalServerError)
		return
	}
}
func (h *Handler) HandlerGet(w http.ResponseWriter, r *http.Request) {
	tx, err := h.database.BeginTransaction()
	if err != nil {
		http.Error(w, "error start transation: %v", http.StatusInternalServerError)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
			http.Error(w, "error sql", http.StatusInternalServerError)
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Query("SELECT * FROM schema_store.users")
	if err != nil {
		http.Error(w, fmt.Sprintf("error sql query : %v", err), http.StatusInternalServerError)
	}
	_, err = tx.Query("SELECT * FROM schema_store.products")
	if err != nil {
		http.Error(w, fmt.Sprintf("error sql query : %v", err), http.StatusInternalServerError)
	}
	_, err = tx.Query("SELECT * FROM schema_store.orders WHERE user_id = 1")
	if err != nil {
		http.Error(w, fmt.Sprintf("error sql query : %v", err), http.StatusInternalServerError)
	}
}
func (h *Handler) handlerRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.HandlerGet(w, r)
	case http.MethodPost:
		h.HandlerPost(w, r)
	default:
		http.Error(w, "Method  not supported", http.StatusMethodNotAllowed)
	}
}
