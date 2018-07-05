#!/bin/sh

rm -rf model

swagger generate server -f openapi.yml
dep ensure
go build ./cmd/opus47-server

