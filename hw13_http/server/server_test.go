// server_test.go
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/example", nil)
	assert.NoError(t, err, "Не должно быть ошибок при создании запроса")
	recorder := httptest.NewRecorder()
	handlerRequest(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code, "Ответ не соответствует ожидаемому")
	expectedResponse := "Привет, я первый сервер"
	assert.Equal(t, expectedResponse, recorder.Body.String(), "Тело ответа не соответствует ожидаемому")
}
