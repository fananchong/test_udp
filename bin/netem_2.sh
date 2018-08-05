#!/bin/bash


NETCARD=eth0


if [ $1 -eq 1 ]; then
	set -x
	modprobe ifb
	ip link set dev ifb0 up
	tc qdisc add dev $NETCARD ingress
	tc filter add dev $NETCARD parent ffff: protocol ip u32 match u32 0 0 flowid 1:1 action mirred egress redirect dev ifb0
	tc qdisc add dev $NETCARD root netem delay 100ms 20ms distribution normal loss $2%
	tc qdisc add dev ifb0 root netem delay 100ms 20ms distribution normal loss $2%

	tc qdisc show dev $NETCARD
	tc qdisc show dev ifb0
	tc filter list dev $NETCARD parent ffff:fff1
fi

if [ $1 -eq 0 ]; then
	set -x
	tc qdisc del dev ifb0 root netem
	tc qdisc del dev $NETCARD root netem
	tc qdisc del dev $NETCARD ingress
	ip link set dev ifb0 down
	
	tc qdisc show dev $NETCARD
        tc qdisc show dev ifb0
        tc filter list dev $NETCARD parent ffff:fff1
fi

