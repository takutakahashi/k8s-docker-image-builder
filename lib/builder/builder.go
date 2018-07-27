package builder

import (
	"bufio"
	"github.com/labstack/echo"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/container"
	"io"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func GetTarFile(c echo.Context) io.Reader {
	recievedFile, err := c.FormFile("file")
	f, err := recievedFile.Open()
	check(err)
	return bufio.NewReader(f)
}

type BuildStatus struct {
	CreatedAt time.Time
	Status    string
}

func Status(id int) BuildStatus {
	return BuildStatus{CreatedAt: time.Now(), Status: "ready"}
}

func List() []string {
	return []string{"ready", "ready"}
}

func Build(tar io.Reader, imageName string) string {
	//github.Clone(repoName)
	response := container.Build(tar, imageName)
	return response
}

func Pull(c echo.Context, image string) {
	container.Pull(c, image)
}

func Push(image string) {
	container.Push(image)
}
