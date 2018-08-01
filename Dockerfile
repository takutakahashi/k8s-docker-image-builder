FROM golang:1.10 as builder

COPY . /go/src/github.com/takutakahashi/k8s-docker-image-builder
WORKDIR /go/src/github.com/takutakahashi/k8s-docker-image-builder
RUN go get -v
RUN go build

FROM google/cloud-sdk:slim
COPY --from=builder /go/src/github.com/takutakahashi/k8s-docker-image-builder/k8s-docker-image-builder /kdib
ADD start.sh /start.sh
CMD '/start.sh'
