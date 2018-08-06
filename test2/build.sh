#!/bin/bash

set -ex

docker run --rm -e GOPATH=/go/ -v "$PWD":/go/src/kcpserver -w /go/src/kcpserver golang go build ./kcpserver.go ./common.go
docker run --rm -e GOPATH=/go/ -v "$PWD":/go/src/kcpclient -w /go/src/kcpclient golang go build ./kcpclient.go ./common.go

cp -f kcpserver ../bin/test2
cp -f kcpclient ../bin/test2

rm -rf ./kcpserver
rm -rf ./kcpclient

