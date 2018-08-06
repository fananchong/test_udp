#!/bin/bash

./stop.sh

set -x

nohup ./kcpserver --mode=0 > /dev/null 2>&1 &
nohup ./kcpserver --mode=1 > /dev/null 2>&1 &

sleep 3s

ps -aux | grep server
