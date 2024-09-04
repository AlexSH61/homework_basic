package store_test

import (
	"testing"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/model"
	"github.com/AlexSH61/homework_basic/hw15_go_sql/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseurl)
	defer teardown("users")
	u, err := s.User().InsertUser(&model.User)
	assert.Equal(t)

}
