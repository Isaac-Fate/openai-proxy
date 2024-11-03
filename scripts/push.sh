#!/bin/sh

DOCKER_IMAGE_NAME=openai-proxy
PROJECT_ID=openai-proxy-440516
REPOSITORY_NAME=openai-proxy
TAG=latest
TARGET_IMAGE_NAME=us-east1-docker.pkg.dev/$PROJECT_ID/$REPOSITORY_NAME/$DOCKER_IMAGE_NAME:$TAG

# Tag the image
docker tag $DOCKER_IMAGE_NAME $TARGET_IMAGE_NAME

# Push the image to the Google artifact registry
docker push $TARGET_IMAGE_NAME
