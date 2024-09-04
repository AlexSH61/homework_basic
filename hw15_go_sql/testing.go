package store

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/store"
)

func TestStore(t *testing.T, databaseurl string) (*Store, func(...string)) {
	t.Helper()

	config := store.NewConfig()
	config.DatabaseURL = databaseurl
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}
	return s, func(tables ...string) {
		if len(tables) > 0 {
			query := fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
			if _, err := s.db.Exec(context.Background(), query); err != nil {
				t.Fatal(err)
			}
			s.Close()
		}
	}
}
