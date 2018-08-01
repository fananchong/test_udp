package main

import (
	"fmt"
	"io"
	"time"

	"github.com/fananchong/gotcp"
	"github.com/xtaci/kcp-go"
)

var gChartSession *ChartSession

func main() {
	addrs := "127.0.0.1:5002"
	gChartSession = &ChartSession{}
	gChartSession.Connect("127.0.0.1:3333", gChartSession)
	gChartSession.Verify()
	KcpClient(addrs)
}

func KcpClient(addrs string) {
	conn, err := kcp.DialWithOptions(addrs, nil, 10, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to ", conn.RemoteAddr().String())

	conn.Write([]byte("hello!"))

	var tempbuf [1024]byte
	for {
		readnum, err := io.ReadAtLeast(conn, tempbuf[0:], 400)
		if err != nil {
			fmt.Println(err)
			return
		}
		if readnum != 400 {
			fmt.Println("readnum != 400")
			panic("data error!")
		}
		now := time.Now().UnixNano()
		onKcpRecv(tempbuf[:400], now)
	}
}

var (
	preTCPRecvTime int64 = 0
)

func onKcpRecv(data []byte, now int64) {
	if preTCPRecvTime == 0 {
		preTCPRecvTime = now
	}

	detal := (now - preTCPRecvTime) / int64(time.Millisecond)
	preTCPRecvTime = now
	temp := []byte(fmt.Sprintf("2_%d", detal))
	gChartSession.SendRaw(temp)
}

type ChartSession struct {
	gotcp.Session
}

func (this *ChartSession) OnRecv(data []byte, flag byte) {

}

func (this *ChartSession) OnClose() {

}

