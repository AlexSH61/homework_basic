// server_test.go
package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/example?msg=Hello", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerGet(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Hey, i'm first server and you message :Hello\n\n", rr.Body.String())
}

func TestHandlerPost(t *testing.T) {
	req, err := http.NewRequest("POST", "/example?msg=Greetings", strings.NewReader(""))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerPost(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Answer from POST request: \n Greetings hi, I'm the first server: \n", rr.Body.String())
}
