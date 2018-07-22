package container

import(
  "io"
  "io/ioutil"
  "context"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/client"
)
//https://kuroeveryday.blogspot.com/2017/09/golang-build-image-with-dockerfile.html
func Build(tar io.Reader, image string) string{
  ctx := context.Background()
  cli, err := client.NewEnvClient()

  buildOpt := types.ImageBuildOptions{
        PullParent:     true,
        Tags:           []string{image},
    }
  buildResponse, err := cli.ImageBuild(ctx, tar, buildOpt)
	if err != nil {
		panic(err)
	}
  b, err := ioutil.ReadAll(buildResponse.Body)
  return string(b)
}
