#!/bin/bash

set -ex

docker run --rm -e GOPATH=/go/ -v "$PWD":/go/src/kcpserver -w /go/src/kcpserver golang go build ./kcpserver.go ./common.go
docker run --rm -e GOPATH=/go/ -v "$PWD":/go/src/kcpclient -w /go/src/kcpclient golang go build ./kcpclient.go ./common.go

cp -f kcpserver ../../bin/test1
cp -f kcpclient ../../bin/test1


#docker build -t kcpserver .
#docker build -t kcpclient .

#set +ex

#docker rm -f kcpserver

#docker run -d -p 5002:5002 --restart always --name kcpserver kcpserver

