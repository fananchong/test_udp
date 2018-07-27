#!/bin/bash

docker run --rm -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d github.com/xtaci/kcp-go