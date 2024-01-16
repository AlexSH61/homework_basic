package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AlexSH61/homework_basic/hw13_http/client"
	"github.com/AlexSH61/homework_basic/hw13_http/server"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Использование: main <адрес> <порт>")
		os.Exit(1)
	}

	address := os.Args[1]
	port := os.Args[2]
	go server.StartServer(address, port)
	time.Sleep(2 * time.Second)
	getResponse, err := client.SendGetRequest(fmt.Sprintf("%s:%s", address, port), "example")
	if err != nil {
		fmt.Println("Ошибка при отправке GET запроса:", err)
		return
	}
	fmt.Println("Ответ на GET запрос:", getResponse)

	postResponse, err := client.SendPostRequest(fmt.Sprintf("%s:%s", address, port), "example", []byte("Привет"))
	if err != nil {
		fmt.Println("Ошибка при отправке POST запроса:", err)
		return
	}
	fmt.Println("Ответ на POST запрос:", postResponse)
}
