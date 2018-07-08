#!/bin/bash

set -e

if [[ $# -lt 1 ]]; then
  echo "usage recycle: <what>"
  exit 1
fi

docker-compose stop $@
docker-compose rm -vf $@
docker-compose up -d $@
