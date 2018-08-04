#!/bin/bash

set -ex

docker run --rm -e GOPATH=/go/ -v "$PWD":/go/src/gochart -w /go/src/gochart golang go build ./...

#docker build -t gochart .

#set +ex

#docker rm -f gochart

#docker run -d -p 3333:3333 -p 8000:8000 --restart always --name gochart gochart --showtext1=66


rm -rf gochart

