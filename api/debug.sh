#!/bin/sh

docker kill opus47-api
docker run \
  -v `pwd`:/go/src/github.com/opus47/cloud/api \
  -it --name opus47-api --rm opus47/api:latest \
  bash
