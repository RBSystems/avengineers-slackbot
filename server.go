package main

import (
	"fmt"

	"github.com/byuoitav/avengineers-slackbot/controllers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	// err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/ftp-microservice/master/swagger.yml")
	// if err != nil {
	// 	fmt.Printf("Could not load swagger.yaml file. Error: %s", err.Error())
	// 	panic(err)
	// }

	port := ":9000"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	// e.Get("/", hateoas.RootResponse)
	router.Get("/health", health.Check)

	router.Post("/message", controllers.Message)
	router.Post("/slack", controllers.Slack)

	fmt.Printf("The Avengineers Slackbot is listening on %s\n", port)
	router.Run(fasthttp.New(port))
}
