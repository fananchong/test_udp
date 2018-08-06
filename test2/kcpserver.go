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

	port := 5004
	interval := time.Duration(param1) * time.Millisecond
	KcpServer(port, interval)
}

func KcpServer(port int, interval time.Duration) {
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
			for {
				// echo
				var buff [102400]byte
				var n int
				var err error
				if n, err = conn.Read(buff[:]); err != nil {
					fmt.Println(err)
					return
				}
				if _, err = conn.Write(buff[:n]); err != nil {
					fmt.Println(err)
					return
				}
			}
		}()
	}
}
