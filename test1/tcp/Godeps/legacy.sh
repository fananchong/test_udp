#!/bin/bash

set -ex

docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d github.com/golang/protobuf/proto
docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d github.com/fananchong/gotcp
docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d github.com/fananchong/gochart
docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d github.com/bitly/go-simplejson

