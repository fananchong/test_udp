#!/bin/bash

rm -rf CMakeCache.txt
cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_TOOLCHAIN_FILE=~/vcpkg/scripts/buildsystems/vcpkg.cmake .
make
cp -f ./server/server ../../bin/test1/raknet_server
cp -f ./client/client ../../bin/test1/raknet_client

