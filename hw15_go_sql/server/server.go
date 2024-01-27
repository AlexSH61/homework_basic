package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/database"
)

var db *database.DB

func SetDB(database *database.DB) {
	db = database
}

func StartServer(address, port string) {
	http.HandleFunc("/example", handlerRequest)
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

func handlerRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlerGet(w, r)
	case http.MethodPost:
		handlerPost(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	fmt.Printf("GET request from %s\n", r.RemoteAddr)
	if msg == "" {
		fmt.Fprintln(w, "Empty")
	} else {
		err := db.InsertUser("sasha", "ninja@gmail.com", "1234")
		if err != nil {
			fmt.Println("Error inserting user:", err)
			http.Error(w, "Error inserting user", http.StatusInternalServerError)
			return
		}
		err = db.UpdateUserName(1, "pasha")
		if err != nil {
			fmt.Println("Error inserting user:", err)
			http.Error(w, "Error inserting user", http.StatusInternalServerError)
			return
		}
		err = db.DeleteUser(1)
		if err != nil {
			fmt.Println("Error inserting user:", err)
			http.Error(w, "Error inserting user", http.StatusInternalServerError)
			return
		}

		response := fmt.Sprintf("Hey, I'm the first server and your message is: %s\n", msg)
		fmt.Fprintln(w, response)
	}
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POST request %s from: %s\n", r.URL.Query().Get("msg"), r.RemoteAddr)
	msg := r.URL.Query().Get("msg")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}
	err = db.UpdateUserName(1, "pasha")
	if err != nil {
		fmt.Println("Error updating user name:", err)
		http.Error(w, "Error updating user name", http.StatusInternalServerError)
		return
	}
	responseData := fmt.Sprintf("Answer from POST request: \n %s hi, : %s\n", msg, body)
	fmt.Fprint(w, responseData)
}
