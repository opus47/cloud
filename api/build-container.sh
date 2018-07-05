#!/bin/sh

./build.sh

docker build -f api.dock -t opus47.net:latest .
