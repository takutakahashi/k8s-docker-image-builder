package container

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/cli/cli/config"
	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/labstack/echo"
	"io"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func streamResponseToClient(c echo.Context, r io.ReadCloser) {
	for true {
		n, _ := io.Copy(c.Response(), r)
		fmt.Printf("%d", n)
		if n == 0 {
			break
		}
		c.Response().Flush()
		time.Sleep(1)
	}
}

//https://kuroeveryday.blogspot.com/2017/09/golang-build-image-with-dockerfile.html
func Build(c echo.Context, tar io.Reader, image string) string {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	buildOpt := types.ImageBuildOptions{
		PullParent: true,
		Tags:       []string{image},
	}
	buildResponse, err := cli.ImageBuild(ctx, tar, buildOpt)
	check(err)
	streamResponseToClient(c, buildResponse.Body)
	return "ok"
}

func Pull(c echo.Context, image string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	auth := getEncodedAuthJSON(image)
	pullOpt := types.ImagePullOptions{RegistryAuth: auth}
	response, err := cli.ImagePull(ctx, image, pullOpt)
	check(err)
	streamResponseToClient(c, response)
}

func Push(c echo.Context, image string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	auth := getEncodedAuthJSON(image)
	pushOpt := types.ImagePushOptions{RegistryAuth: auth}
	response, err := cli.ImagePush(ctx, image, pushOpt)
	check(err)
	streamResponseToClient(c, response)
}

func getEncodedAuthJSON(image string) string {
	named, err := reference.ParseNamed(image)
	check(err)
	hostname, _ := reference.SplitHostname(named)
	configFile := config.LoadDefaultConfigFile(os.Stderr)
	authConfig, err := configFile.GetAuthConfig(hostname)
	check(err)
	encodedJSON, err := json.Marshal(authConfig)
	check(err)
	return base64.URLEncoding.EncodeToString(encodedJSON)
}
