package store

import (
	"os"
	"testing"
)

var databaseurl string

func TestMain(m *testing.M) {
	databaseurl = os.Getenv("DATABASE_URL")
	if databaseurl == "" {
		databaseurl = "host=localhost dbname=hw15_go_sql sslmode=disable"
	}
	os.Exit(m.Run())
}
