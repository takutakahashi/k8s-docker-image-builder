package routes

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/auth"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/builder"
	"net/http"
	"reflect"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var chMap map[string](chan int)

func makeChannel(key string) chan int {
	if chMap == nil {
		chMap = map[string](chan int){}
	}
	chMap[key] = make(chan int)
	return chMap[key]
}

func closeChannel(key string) {
	close(chMap[key])
	delete(chMap, key)
}

func list() []string {
	keys := reflect.ValueOf(chMap).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
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
	ch := makeChannel(image)
	go func(c echo.Context, ch chan int) {
		builder.BuildFromRepo(repo, branch, image)
		builder.Push(image)
		closeChannel(image)
	}(c, ch)
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
	builder.Push(image)
	return nil
}

func statusCheck(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return nil
}

func buildList(c echo.Context) error {
	err := setResponseBase(c)
	check(err)
	return c.JSON(http.StatusOK, list())
}

func Route(e *echo.Echo) *echo.Echo {
	e.POST("/pull", pull)
	e.POST("/push", push)
	e.POST("/build", build)
	e.POST("/publish", publish)
	e.GET("/status", statusCheck)
	e.GET("/build/list", buildList)
	return e
}
