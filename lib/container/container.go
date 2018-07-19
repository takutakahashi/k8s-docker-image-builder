package container

import(
  "os"
  "fmt"
  "bufio"
  "io/ioutil"
  "context"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/client"
)

const BUFSIZE = 1024
//https://kuroeveryday.blogspot.com/2017/09/golang-build-image-with-dockerfile.html
func Build(path string, image string){
  ctx := context.Background()
  cli, err := client.NewEnvClient()
  tar, _ := os.Open("/tmp/Dockerfile.tar.gz")
  if err != nil {
		panic(err)
	}
  //dockerfile := readDockerfile(path)
  buildOpt := types.ImageBuildOptions{
        PullParent:     true,
        Tags:           []string{"lat:1"},
    }
  buildResponse, err := cli.ImageBuild(ctx, bufio.NewReader(tar), buildOpt)
  //containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
  b, err := ioutil.ReadAll(buildResponse.Body)
  fmt.Println(string(b))
}

func readDockerfile(repoPath string) string{
  file, err := os.Open(repoPath+"/Dockerfile")
  if err != nil {
        // Openエラー処理
  }
  defer file.Close()
  buf := make([]byte, BUFSIZE)
    for {
        n, err := file.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            // Readエラー処理
            break
        }
    }
    return string(buf)
}
