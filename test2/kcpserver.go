package main

import (
	"flag"
	"fmt"

	"github.com/xtaci/kcp-go"
)

func main() {
	mode := 0
	flag.IntVar(&mode, "mode", 0, "mode")
	flag.Parse()

	var datashards int = 0
	var parity int = 0
	var port int = 5004

	if mode == 1 {
		datashards = 2
		parity = 1
		port = 5005
	}

	KcpServer(port, datashards, parity)
}

func KcpServer(port, datashards, parity int) {
	lis, err := kcp.ListenWithOptions(fmt.Sprintf(":%d", port), nil, datashards, parity)
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
