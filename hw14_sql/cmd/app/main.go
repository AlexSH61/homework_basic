package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape("aleksandr"),
		url.QueryEscape("pwd1234"),
		"localhost",
		"54320",
		"hw14_sql",
		5)

	ctx, _ := context.WithCancel(context.Background())
	poolConfig, _ := pgxpool.ParseConfig(connStr)
	for i := 1; 1 < 5; i++ {
		go func(i int) {
			conn, err := pgxpool.ConnectConfig(ctx, poolConfig)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Connection failled %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Connection OK")
			err = conn.Ping(ctx)
			if err != nil {
				fmt.Fprintf(os.Stderr, "PING connected failled %v\n", err)
				os.Exit(1)
			}
			fmt.Println(i, "Ping connection, status OK")
			defer conn.Close()
		}(i)
	}
}
