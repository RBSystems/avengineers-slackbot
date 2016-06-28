package helpers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostToSlackSuccess(test *testing.T) {
	// Setup
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	os.Setenv("SLACKBOT_WEBHOOK", server.URL)

	assert.NoError(test, PostToSlack("This is a test message"))
}

func TestPostToSlackFail(test *testing.T) {
	// Setup
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	os.Setenv("SLACKBOT_WEBHOOK", server.URL)

	assert.Error(test, PostToSlack("This is a test message"))
}
