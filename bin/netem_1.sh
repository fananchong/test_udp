#!/bin/bash


NETCARD=eth0

if [ $1 -eq 1 ]; then
	set -x
	tc qdisc add dev $NETCARD root netem delay 100ms 20ms distribution normal loss $2%
	tc qdisc show dev $NETCARD
fi

if [ $1 -eq 0 ]; then
	set -x
	tc qdisc del dev $NETCARD root netem
	tc qdisc show dev $NETCARD
fi

