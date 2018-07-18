package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "github.com/takutakahashi/k8s-docker-image-builder/lib/builder"
)

func Build(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  repoName, imageName := p.ByName("repository"), p.ByName("image")
  builder.Build(repoName, imageName)
  fmt.Fprintf(w, "%q", "ok")
}

func main() {
  router := httprouter.New()
  router.POST("/build/*repository", Build)
  log.Fatal(http.ListenAndServe(":8080", router))
}
