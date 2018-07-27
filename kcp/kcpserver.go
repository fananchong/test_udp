package main

import (
	"fmt"
	"strconv"
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
			counter := 0
			for {
				select {
				case <-t.C:
					prefix := []byte(strconv.Itoa(counter) + "_")
					conn.Write(append(prefix, msg...))
				}
			}
		}()
	}
}

func getmsg() []byte {
	var msg []byte
	for i := 0; i < 400; i++ {
		msg = append(msg, 97)
	}
	return msg
}
