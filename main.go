package main

import (
	"bufio"
	"github.com/labstack/echo"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/builder"
	"io"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getTarFile(c echo.Context) io.Reader {
	recievedFile, err := c.FormFile("file")
	f, err := recievedFile.Open()
	check(err)
	return bufio.NewReader(f)
}

func setResponseBase(c echo.Context) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
}

func build(c echo.Context) error {
	imageName := c.FormValue("image")
	go builder.Build(getTarFile(c), imageName)
	return nil
}

//func Pull(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func pull(c echo.Context) error {
	setResponseBase(c)
	image := c.FormValue("image")
	builder.Pull(image)
	return nil
}

func route(e *echo.Echo) *echo.Echo {
	e.POST("/pull", pull)
	e.POST("/build", build)
	return e
}

func main() {
	log.Fatal(route(echo.New()).Start(":8080"))
}

// check https://echo.labstack.com/cookbook/streaming-response
