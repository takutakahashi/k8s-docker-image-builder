package main

import (
  "fmt"
  "log"
  "io"
  "bufio"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "github.com/takutakahashi/k8s-docker-image-builder/lib/builder"
)

func check(err error){
  if err!= nil {
    panic(err)
  }
}

func getTarFile(r *http.Request) io.Reader {
  err := r.ParseMultipartForm(5 * 1024 * 1024)
  check(err)
  recievedFile, _, err := r.FormFile("file")
  check(err)
  return bufio.NewReader(recievedFile)
}

func Build(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  response := builder.Build(getTarFile(r), p.ByName("image"))
  fmt.Fprintf(w, "%q", response)
}

func main() {
  router := httprouter.New()
  router.POST("/build", Build)
  log.Fatal(http.ListenAndServe(":8080", router))
}
