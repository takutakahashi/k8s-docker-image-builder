FROM golang:1.10

# install lsb-release
RUN apt-get update -y && apt-get install lsb-release -y
# install gcloud
RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
    echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    apt-get update -y && apt-get install google-cloud-sdk -y

WORKDIR /go/src/github.com/takutakahashi/k8s-docker-image-builder
COPY . /go/src/github.com/takutakahashi/k8s-docker-image-builder
RUN go get -v
