package main

import (
	"time"

	"github.com/xtaci/kcp-go"
)

// kcp fast模式
func setParam(conn *kcp.UDPSession) {
	conn.SetStreamMode(true)
	conn.SetWindowSize(4096, 4096)
	conn.SetDSCP(46)
	conn.SetMtu(1400)
	conn.SetReadDeadline(time.Now().Add(time.Hour))
	conn.SetWriteDeadline(time.Now().Add(time.Hour))
	conn.SetACKNoDelay(true)
	conn.SetNoDelay(1, 10, 2, 1)
}
