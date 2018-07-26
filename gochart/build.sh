#!/bin/bash

docker run --rm -e GOPATH=/go/:/temp/ -v /d/temp/:/temp/ -v "$PWD":/go/src/gochart -w /go/src/gochart golang go build ./...

docker build -t gochart .

docker rm -f gochart

docker run -d -p 3333:3333 --restart always --name gochart gochart
