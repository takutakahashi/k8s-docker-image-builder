package routes

import (
	"github.com/labstack/echo"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/builder"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func setResponseBase(c echo.Context) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
}

func build(c echo.Context) error {
	image := c.FormValue("image")
	go builder.Build(builder.GetTarFile(c), image)
	return nil
}

func pull(c echo.Context) error {
	setResponseBase(c)
	image := c.FormValue("image")
	builder.Pull(c, image)
	return nil
}

func push(c echo.Context) error {
	setResponseBase(c)
	image := c.FormValue("image")
	builder.Push(image)
	return nil
}

func Route(e *echo.Echo) *echo.Echo {
	e.POST("/pull", pull)
	e.POST("/push", push)
	e.POST("/build", build)
	return e
}
