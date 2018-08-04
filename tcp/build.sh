#!/bin/bash

set -ex

docker run --rm -e GOPATH=/go/ -v "$PWD":/go/src/tcpserver -w /go/src/tcpserver golang go build ./tcpserver.go
docker run --rm -e GOPATH=/go/ -v "$PWD":/go/src/tcpclient -w /go/src/tcpclient golang go build ./tcpclient.go

cp -f tcpserver ../bin
cp -f tcpclient ../bin
