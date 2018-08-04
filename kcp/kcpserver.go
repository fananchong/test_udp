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
	lis, err := kcp.ListenWithOptions(fmt.Sprintf("0.0.0.0:%d", port), nil, 0, 0)
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
			// 每 100ms 发送一次 400byte 消息
			t := time.NewTicker(interval)
			for {
				select {
				case <-t.C:
					_, err := conn.Write(msg)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}()
	}
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
