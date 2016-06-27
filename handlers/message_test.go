package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

var messageJSON = `{"text":"This is a test message"}`

func TestMessageSuccess(test *testing.T) {
	// Setup
	router := echo.New()
	request := new(http.Request)
	recorder := httptest.NewRecorder()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	os.Setenv("SLACKBOT_WEBHOOK", server.URL)

	context := router.NewContext(standard.NewRequest(request, router.Logger()), standard.NewResponse(recorder, router.Logger()))
	context.SetPath("/message")
	context.SetParamNames("body")
	context.SetParamValues(messageJSON)

	// Assertions
	if assert.NoError(test, Message(context)) {
		assert.Equal(test, http.StatusOK, recorder.Code)
	}
}

func TestMessageFail(test *testing.T) {
	// Setup
	router := echo.New()
	request := new(http.Request)
	recorder := httptest.NewRecorder()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	os.Setenv("SLACKBOT_WEBHOOK", server.URL)

	context := router.NewContext(standard.NewRequest(request, router.Logger()), standard.NewResponse(recorder, router.Logger()))
	context.SetPath("/message")
	context.SetParamNames("body")
	context.SetParamValues(messageJSON)

	// Assertions
	assert.Error(test, Message(context))
}
