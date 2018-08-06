#!/bin/bash

set -ex

docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d github.com/golang/protobuf/proto
docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d github.com/fananchong/gotcp
docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang /bin/bash -c "git config --global http.proxy http://127.0.0.1:8123 && git config --global https.proxy https://127.0.0.1:8123 && go get -u -d -insecure github.com/xtaci/kcp-go"
docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d github.com/fananchong/gochart
docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d github.com/bitly/go-simplejson

