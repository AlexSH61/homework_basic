package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexSH61/homework_basic/hw15_go_sql/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestAPiserver_HandlerHello(t *testing.T) {
	s, _ := New(config.NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello")
}
