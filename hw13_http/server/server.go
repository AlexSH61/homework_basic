package server

import (
	"fmt"
	"net/http"
)

func StartServer(address, port string) {
	http.HandleFunc("/example", handlerRequest)
	serverAddress := fmt.Sprintf("%s:%s", address, port)
	fmt.Printf("Сервер слушает на %s\n", serverAddress)
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}

func handlerRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Получен запрос типа %s от %s\n", r.Method, r.RemoteAddr)
	responseData := "Привет, я первый сервер"
	fmt.Fprint(w, responseData)
}
