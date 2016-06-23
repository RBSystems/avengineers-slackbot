package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

var messageJSON = `{"text":"This is a test message"}`

func TestMessage(test *testing.T) {
	// Setup
	router := echo.New()
	request := new(http.Request)
	recorder := httptest.NewRecorder()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(writer, "") // Slack responds silently
	}))
	defer server.Close()

	err := os.Setenv("SLACKBOT_WEBHOOK", server.URL)
	assert.NoError(test, err)

	context := router.NewContext(standard.NewRequest(request, router.Logger()), standard.NewResponse(recorder, router.Logger()))
	context.SetPath("/message")
	context.SetParamNames("body")
	context.SetParamValues(messageJSON)

	// Assertions
	if assert.NoError(test, Message(context)) {
		assert.Equal(test, http.StatusOK, recorder.Code)
		assert.Equal(test, recorder.Body.String(), "")
	}
}
