package main

import (
	"fmt"
	"time"

	"github.com/xtaci/kcp-go"
)

func main() {
	port := 5002
	interval := 100 * time.Millisecond
	msg := getmsg()
	KcpServer(port, interval, msg)
}

func KcpServer(port int, interval time.Duration, msg []byte) {
	lis, err := kcp.ListenWithOptions(fmt.Sprintf(":%d", port), nil, 10, 3)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("on connect. addr =", conn.RemoteAddr())

		go func() {
			// 每100ms发送一次 hello消息
			t := time.NewTicker(interval)
			for {
				select {
				case <-t.C:
					conn.Write(msg)
					fmt.Println("aaaaaaaaaaaaaaaaaaaaaa")
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
