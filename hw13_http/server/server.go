package server

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

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
		http.Error(w, "Method  not supported", http.StatusMethodNotAllowed)
	}
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	fmt.Printf("GET request from %s\n", r.RemoteAddr)
	if msg == "" {
		fmt.Fprintln(w, "Empty")
	} else {
		response := fmt.Sprintf("Hey, i'm first server and you message :%s\n", msg)
		fmt.Fprintln(w, response)
	}
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POST request %s from: %s\n ", r.URL.Query().Get("msg"), r.RemoteAddr)
	msg := r.URL.Query().Get("msg")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}
	responseData := fmt.Sprintf("Answer from POST request: \n %s hi, I'm the first server: %s\n", msg, body)
	fmt.Fprint(w, responseData)
}
