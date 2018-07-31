package main

import (
	"fmt"
	"io"
	"time"

	"github.com/xtaci/kcp-go"
)

func main() {
	addrs := "127.0.0.1:5002"
	KcpClient(addrs)
}

func KcpClient(addrs string) {
	conn, err := kcp.DialWithOptions(addrs, nil, 10, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to ", conn.RemoteAddr().String())

	//		buf := gotcp.NewByteBuffer()
	//		msglen := len(MSG)
	var tempbuf [1024]byte
	for {
		//			leastlen := msglen - buf.RdSize()
		readnum, err := io.ReadAtLeast(conn, tempbuf[0:], 400)
		if err != nil {
			fmt.Println(err)
			return
		}
		if readnum != 400 {
			fmt.Println("readnum != 400")
			panic("data error!")
		}
		fmt.Println("======================")
		//			buf.Append(tempbuf[:readnum])
		now := time.Now().UnixNano()
		//			for buf.RdSize() >= msglen {
		//				msgbuff := buf.RdBuf()
		onKcpRecv(tempbuf[:400], now)
		//				buf.RdFlip(msglen)
		//			}
	}
}

var (
	preTCPRecvTime int64 = 0
)

func onKcpRecv(data []byte, now int64) {
	if preTCPRecvTime == 0 {
		preTCPRecvTime = now
	}

	//	detal := (now - preTCPRecvTime) / int64(time.Millisecond)
	//	preTCPRecvTime = now

	//	g_chart.AddKcpData(detal)
}
