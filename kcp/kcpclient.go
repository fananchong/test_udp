package main

import (
	"flag"
	"fmt"
	"io"
	"time"

	"github.com/fananchong/gotcp"
	"github.com/xtaci/kcp-go"
)

var gChartSession *ChartSession

func main() {
	param1 := "127.0.0.1"
	flag.StringVar(&param1, "ip", "127.0.0.1", "ip")
	flag.Parse()

	addrs := fmt.Sprintf("%s:5002", param1)
	gChartSession = &ChartSession{}
	gChartSession.Connect("127.0.0.1:3333", gChartSession)
	gChartSession.Verify()
	KcpClient(addrs)
}

func KcpClient(addrs string) {
	conn, err := kcp.DialWithOptions(addrs, nil, 0, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to ", conn.RemoteAddr().String())

	setParam(conn)

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
	delaySlice     []int64
)

func onKcpRecv(data []byte, now int64) {
	if preTCPRecvTime == 0 {
		preTCPRecvTime = now
	}

	detal := (now - preTCPRecvTime) / int64(time.Millisecond)
	preTCPRecvTime = now
	temp := []byte(fmt.Sprintf("2_%d;", detal))
	gChartSession.SendRaw(temp)
	delaySlice = append(delaySlice, detal)
	if len(delaySlice)%200 == 0 {
		fmt.Println(delaySlice)
	}
}

type ChartSession struct {
	gotcp.Session
}

func (this *ChartSession) OnRecv(data []byte, flag byte) {

}

func (this *ChartSession) OnClose() {

}
