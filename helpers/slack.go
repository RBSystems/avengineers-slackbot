package helpers

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"strconv"
)

func PostToSlack(message string) error {
	webhook := os.Getenv("SLACKBOT_WEBHOOK")

	request, err := http.NewRequest("POST", webhook, bytes.NewBuffer([]byte(`{"text": "`+message+`"}`)))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("Error: Slack returned a status code of: " + strconv.Itoa(response.StatusCode))
	}

	return nil
}
