package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AlexSH61/homework_basic/hw13_http/client"
	"github.com/AlexSH61/homework_basic/hw13_http/server"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("Using: main <address> <port>")
		os.Exit(1)
	}
	address := os.Args[1]
	port := os.Args[2]
	go server.StartServer(address, port)
	time.Sleep(1 * time.Second)
	getResponse, err := client.SendGetRequest(fmt.Sprintf("%s:%s", address, port), "example")
	if err != nil {
		log.Println("Error sending GET request:", err)
		return
	}
	log.Println("Answer from Get request:\n", getResponse)

	postResponse, err := client.SendPostRequest(fmt.Sprintf("%s:%s", address, port), "example", []byte("test"))
	if err != nil {
		log.Println("Error sending request POST:", err)
		return
	}
	log.Println("Answer from POST on request:\n", postResponse)
}
