#!/bin/bash

rm -rf CMakeCache.txt
cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_TOOLCHAIN_FILE=/home/fananchong1/vcpkg/scripts/buildsystems/vcpkg.cmake .
make
cp -f ./server/server ../bin/raknet_server
cp -f ./client/client ../bin/raknet_client

