#!/bin/bash

set -ex

docker run --rm -e GOPATH=/go/ -v "$PWD":/go/src/kcpserver -w /go/src/kcpserver golang go build ./kcpserver.go

docker build -t kcpserver .

set +ex

docker rm -f kcpserver

docker run -d -p 5002:5002 --restart always --name kcpserver kcpserver

