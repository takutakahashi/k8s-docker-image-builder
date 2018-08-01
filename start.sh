#!/bin/sh -e
ls /secrets/credentials.json
ls /root/.ssh/id_rsa
gcloud auth activate-service-account --key-file=/secrets/credentials.json
gcloud --quiet auth configure-docker
ssh git@github.com
/kdib
