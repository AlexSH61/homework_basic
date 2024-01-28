// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/config"
	db "github.com/AlexSH61/homework_basic/hw15_go_sql/database"
	"github.com/AlexSH61/homework_basic/hw15_go_sql/server"
)

func main() {
	cfg, err := config.ReadConfig("setting.cfg")
	if err != nil {
		fmt.Println("error read configurate")
	}
	dbConn, err := db.NewDataBase(cfg)
	if err != nil {
		fmt.Println("Error initializing the database:", err)
		return
	}
	defer dbConn.Close()
	handlers := server.NewHandler(dbConn)

	http.HandleFunc("/post", handlers.HandlerPost)
	http.HandleFunc("/get", handlers.HandlerGet)

	server.StartServer(cfg.ServerHost, cfg.ServerPort)
}
