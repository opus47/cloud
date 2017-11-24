#!/bin/sh

docker kill opus47-api
docker run -d --name opus47-api --rm opus47/api:latest
