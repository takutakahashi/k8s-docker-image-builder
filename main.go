package main

import (
	"github.com/labstack/echo"
	"github.com/takutakahashi/k8s-docker-image-builder/routes"
	"log"
)

func main() {
	log.Fatal(routes.Route(echo.New()).Start(":8080"))
}

// check https://echo.labstack.com/cookbook/streaming-response
