FROM google/cloud-sdk:slim
ADD k8s-docker-image-builder /kdib
ADD start.sh /start.sh
ENTRYPOINT '/start.sh'
