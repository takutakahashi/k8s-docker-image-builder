package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/takutakahashi/k8s-docker-image-builder/routes"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	log.Fatal(routes.Route(e).Start(":8080"))
}

// check https://echo.labstack.com/cookbook/streaming-response
