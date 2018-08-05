#!/bin/bash

./stop.sh

set -x

nohup ./tcpserver --interval=66 > /dev/null 2>&1 &
nohup ./kcpserver --interval=66 > /dev/null 2>&1 &
nohup ./raknet_server 66 > /dev/null 2>&1 &

sleep 5s

ps -aux | grep server

