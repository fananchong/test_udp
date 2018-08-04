package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	param1 := 100
	flag.IntVar(&param1, "interval", 100, "interval")
	flag.Parse()

	port := 5003
	interval := time.Duration(param1) * time.Millisecond
	msg := getmsg()
	TcpServer(port, interval, msg)
}

func TcpServer(port int, interval time.Duration, msg []byte) {
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("on connect. addr =", conn.RemoteAddr())

		go func() {
			conn.SetNoDelay(true)
			conn.SetWriteBuffer(128 * 1024)
			conn.SetReadBuffer(128 * 1024)

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
