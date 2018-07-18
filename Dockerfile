FROM golang:1.10

WORKDIR /go/src/github.com/takutakahashi/k8s-docker-image-builder
COPY . /go/src/github.com/takutakahashi/k8s-docker-image-builder
