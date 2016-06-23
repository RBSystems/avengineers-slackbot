package helpers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostToSlack(test *testing.T) {
	// Setup
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(writer, "") // Slack responds silently
	}))
	defer server.Close()

	err := os.Setenv("SLACKBOT_WEBHOOK", server.URL)
	assert.NoError(test, err)

	assert.NoError(test, PostToSlack("This is a test message"))
}
