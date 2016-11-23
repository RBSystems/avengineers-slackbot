package main

import (
	"log"
	"net/http"

	"github.com/byuoitav/avengineers-slackbot/handlers"
	"github.com/byuoitav/avengineers-slackbot/helpers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/robfig/cron"
)

var doctor = helpers.Hospital{}

func main() {
	err := helpers.LoadConfig(&doctor)
	if err != nil {
		log.Fatal(err)
	}

	timer := cron.New()
	timer.AddFunc("0 0/2 * * * *", func() { // Check health of services every two minutes
		helpers.CheckHealth(&doctor)
	})
	timer.Start()

	port := ":9000"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	router.POST("/message", handlers.Message)

	router.Start(port)
}
