package builder

import (
	"bufio"
	"github.com/labstack/echo"
	"github.com/mholt/archiver"
	"github.com/rs/xid"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/container"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/github"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func dirwalk(path string) []string {
	files, err := ioutil.ReadDir(path)
	check(err)
	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(path, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(path, file.Name()))
	}

	return paths
}
func GetTarFile(c echo.Context) io.Reader {
	recievedFile, err := c.FormFile("file")
	f, err := recievedFile.Open()
	check(err)
	return bufio.NewReader(f)
}

func makeTar(repo string) io.Reader {
	tarPath := "/tmp/" + xid.New().String() + ".tar"
	fileList := dirwalk("/" + repo)
	archiver.Tar.Make(tarPath, fileList)
	f, err := os.Open(tarPath)
	check(err)
	return bufio.NewReader(f)
}

func BuildFromRepo(c echo.Context, repoName string, imageName string) string {
	github.Clone(repoName)
	tar := makeTar(repoName)
	response := container.Build(c, tar, imageName)
	return response
}

func Build(c echo.Context, tar io.Reader, imageName string) string {
	//github.Clone(repoName)
	response := container.Build(c, tar, imageName)
	return response
}

func Pull(c echo.Context, image string) {
	container.Pull(c, image)
}

func Push(c echo.Context, image string) {
	container.Push(c, image)
}
