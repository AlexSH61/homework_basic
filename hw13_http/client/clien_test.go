package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendGetRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/testPath", r.URL.Path)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test response"))
	}))
	defer server.Close()

	serverURL := server.URL[7:]
	response, err := SendGetRequest(serverURL, "testPath")
	assert.Nil(t, err)
	assert.Equal(t, "Test response", response)
}

func TestSendPostRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/testPath", r.URL.Path)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Test response for POST request"))
	}))
	defer server.Close()

	serverURL := server.URL[7:]
	payload := []byte(`{"key": "value"}`)
	response, err := SendPostRequest(serverURL, "testPath", payload)
	assert.Nil(t, err)
	assert.Equal(t, "Test response for POST request", response)
}
