package builder

import (
	"bufio"
	"fmt"
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

func resln(c echo.Context, line string) {
	fmt.Fprintf(c.Response(), "%q\n", line)
	c.Response().Flush()
}

func getFileDirList(path string) []string {
	files, err := ioutil.ReadDir(path)
	check(err)
	var paths []string
	for _, file := range files {
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
	tarPath := xid.New().String() + ".tar"
	archiver.Tar.Make(tarPath, getFileDirList(repo))
	f, err := os.Open(tarPath)
	defer os.RemoveAll(tarPath)
	check(err)
	return bufio.NewReader(f)
}

func BuildFromRepo(c echo.Context, repoName string, branchName string, imageName string) string {
	resln(c, "clone repo")
	repoPath := github.Clone(repoName, branchName)
	defer os.RemoveAll(repoPath)
	resln(c, "making tar")
	tar := makeTar(repoPath)
	resln(c, "build start")
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
