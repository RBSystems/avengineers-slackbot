package main

import (
	"log"

	"github.com/byuoitav/avengineers-slackbot/handlers"
	"github.com/byuoitav/avengineers-slackbot/helpers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := helpers.CheckHealth()
	if err != nil {
		log.Fatal(err)
	}

	port := ":9000"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	router.Get("/health", health.Check)

	router.Post("/message", handlers.Message)
	router.Post("/slack", handlers.Slack)

	log.Println("The Avengineers Slackbot is listening on " + port)
	router.Run(fasthttp.New(port))
}
