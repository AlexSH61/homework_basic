package main

import (
	"fmt"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/database"
	"github.com/AlexSH61/homework_basic/hw15_go_sql/server"
)

func main() {
	dbConfig := database.Config{
		ServerHost: "localhost",
		ServerPort: "8400",
		PgHost:     "127.0.0.1",
		PgPort:     "5432",
		PgUser:     "hw14_sql",
		PgPassword: "1234",
		PgNameBD:   "hw14_sql_db",
	}

	db, err := database.NewDB(dbConfig)
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	defer db.Close()

	server.SetDB(db)

	server.StartServer("127.0.0.1", "8081")
}
