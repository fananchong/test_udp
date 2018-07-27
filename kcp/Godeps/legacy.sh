#!/bin/bash

set -ex

git config http.proxy http://127.0.0.1:8123
git config https.proxy https://127.0.0.1:8123

docker run --rm --net=host -v /d/temp/:/temp/ -e GOPATH=/temp golang go get -u -d -insecure github.com/xtaci/kcp-go

git config --unset http.proxy
git config --unset https.proxy

