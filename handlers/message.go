package handlers

import (
	"github.com/byuoitav/avengineers-slackbot/helpers"
	"github.com/labstack/echo"
)

type message struct {
	Text string `json:"text"`
}

func Message(context echo.Context) error {
	message := message{}
	context.Bind(&message)

	err := helpers.PostToSlack(message.Text)
	if err != nil {
		return err
	}

	return nil
}
