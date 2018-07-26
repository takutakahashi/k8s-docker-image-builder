package main

import (
	"bufio"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/takutakahashi/k8s-docker-image-builder/lib/builder"
	"io"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
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
	imageName := r.FormValue("image")
	go builder.Build(getTarFile(r), imageName)
	fmt.Fprintf(w, "%q", "reserved")
}

func Pull(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseMultipartForm(5 * 1024 * 1024)
	image := r.FormValue("image")
	check(err)
	builder.Pull(image)
}

func main() {
	router := httprouter.New()
	router.POST("/build", Build)
	router.POST("/pull", Pull)
	log.Fatal(http.ListenAndServe(":8080", router))
}
