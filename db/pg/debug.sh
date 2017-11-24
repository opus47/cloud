#!/bin/sh

docker kill opus47-db
docker run -it --name opus47-db --rm  opus47/db:latest
