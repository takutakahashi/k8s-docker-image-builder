package container

import(
  "io"
  "fmt"
  "os"
  "encoding/json"
  "encoding/base64"
  "io/ioutil"
  "context"
  "github.com/docker/distribution/reference"
  "github.com/docker/cli/cli/config"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/client"
)
func check(err error){
  if err!= nil {
    panic(err)
  }
}
//https://kuroeveryday.blogspot.com/2017/09/golang-build-image-with-dockerfile.html
func Build(tar io.Reader, image string) string{
  ctx := context.Background()
  cli, err := client.NewEnvClient()
  buildOpt := types.ImageBuildOptions{
        PullParent:     true,
        Tags:           []string{image},
    }
  buildResponse, err := cli.ImageBuild(ctx, tar, buildOpt)
  check(err)
  b, err := ioutil.ReadAll(buildResponse.Body)
  fmt.Printf("%q", string(b))
  return string(b)
}

func Pull(image string) {
  ctx := context.Background()
  cli, err := client.NewEnvClient()
  auth := getEncodedAuthJSON(image)
  pullOpt := types.ImagePullOptions{RegistryAuth: auth}
  response, err := cli.ImagePull(ctx, image, pullOpt)
  check(err)
  b, err := ioutil.ReadAll(response)
  fmt.Printf("%q", string(b))
}

func Push(image string) {
  ctx := context.Background()
  cli, err := client.NewEnvClient()
  auth := getEncodedAuthJSON(image)
  pushOpt := types.ImagePushOptions{RegistryAuth: auth}
  response, err := cli.ImagePush(ctx, image, pushOpt)
  check(err)
  b, err := ioutil.ReadAll(response)
  fmt.Printf("%q", string(b))
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
