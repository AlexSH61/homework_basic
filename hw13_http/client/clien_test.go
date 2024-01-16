package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendGetRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Ожидался GET-запрос")
		assert.Equal(t, "/testPath", r.URL.Path, "Ожидался URL-путь /testPath")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Тестовый ответ"))
	}))
	defer server.Close()

	serverURL := server.URL[7:]
	response, err := SendGetRequest(serverURL, "testPath")
	assert.Nil(t, err, "Ошибка при отправке GET-запроса")
	assert.Equal(t, "Тестовый ответ", response, "Ожидался ответ Тестовый ответ")
}

func TestSendPostRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Ожидался POST-запрос")
		assert.Equal(t, "/testPath", r.URL.Path, "Ожидался URL-путь /testPath")

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Тестовый ответ на POST-запрос"))
	}))
	defer server.Close()

	serverURL := server.URL[7:]
	payload := []byte(`{"key": "value"}`)
	response, err := SendPostRequest(serverURL, "testPath", payload)
	assert.Nil(t, err, "Ошибка при отправке POST-запроса")
	assert.Equal(t, "Тестовый ответ на POST-запрос", response, "Ожидался ответ Тестовый ответ на POST-запрос")
}
