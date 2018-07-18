package builder

import (
  "github.com/takutakahashi/k8s-docker-image-builder/lib/container"
  "github.com/takutakahashi/k8s-docker-image-builder/lib/github"
  "time"
)

type BuildStatus struct {
  CreatedAt time.Time
  Status string
}

func Status(id int) BuildStatus {
  return BuildStatus {CreatedAt: time.Now(), Status: "ready"}
}

func List() []string {
  return []string{"ready", "ready"}
}

func Build(repoName string, imageName string) string {
  github.Clone(repoName)
  container.Build(repoName, imageName)
  return "ok"
}
