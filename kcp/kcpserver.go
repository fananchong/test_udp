package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/xtaci/kcp-go"
)

func main() {
	param1 := 100
	flag.IntVar(&param1, "interval", 100, "interval")
	flag.Parse()

	port := 5002
	interval := time.Duration(param1) * time.Millisecond
	msg := getmsg()
	KcpServer(port, interval, msg)
}

func KcpServer(port int, interval time.Duration, msg []byte) {
	lis, err := kcp.ListenWithOptions(fmt.Sprintf("0.0.0.0:%d", port), nil, 10, 3)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		setParam(conn.(*kcp.UDPSession))

		fmt.Println("on connect. addr =", conn.RemoteAddr())

		go func() {
			// 每100ms发送一次 hello消息
			t := time.NewTicker(interval)
			for {
				select {
				case <-t.C:
					conn.Write(msg)
				}
			}
		}()
	}
}

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

func getmsg() []byte {
	count := 400
	var msg []byte
	for i := 0; i < count; i++ {
		msg = append(msg, 97)
	}
	msg[count-1] = 0
	return msg
}
