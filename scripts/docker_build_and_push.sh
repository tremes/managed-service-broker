#!/bin/sh

APP_NAME=managed-service-broker

IMAGE_MASTER_TAG=quay.io/integreatly/$APP_NAME:latest

docker login --username $REGISTRY_USERNAME --password $REGISTRY_PASSWORD $REGISTRY_HOST
docker build -t $IMAGE_MASTER_TAF .
docker push $IMAGE_MASTER_TAG
