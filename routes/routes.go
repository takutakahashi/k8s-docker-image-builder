package routes

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/auth"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/builder"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func setResponseBase(c echo.Context) error {
	if !auth.Check(c.FormValue("token")) {
		c.Response().WriteHeader(http.StatusForbidden)
		return fmt.Errorf("not authorized token")
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return nil
}

func build(c echo.Context) error {
	err := setResponseBase(c)
	check(err)
	image := c.FormValue("image")
	builder.Build(c, builder.GetTarFile(c), image)
	return nil
}

func publish(c echo.Context) error {
	err := setResponseBase(c)
	check(err)
	image, repo, branch := c.FormValue("image"), c.FormValue("repo"), c.FormValue("branch")
	builder.BuildFromRepo(c, repo, branch, image)
	builder.Push(c, image)
	return nil
}

func pull(c echo.Context) error {
	err := setResponseBase(c)
	check(err)
	image := c.FormValue("image")
	builder.Pull(c, image)
	return nil
}

func push(c echo.Context) error {
	err := setResponseBase(c)
	check(err)
	image := c.FormValue("image")
	builder.Push(c, image)
	return nil
}

func Route(e *echo.Echo) *echo.Echo {
	e.POST("/pull", pull)
	e.POST("/push", push)
	e.POST("/build", build)
	e.POST("/publish", publish)
	return e
}
