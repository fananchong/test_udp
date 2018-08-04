package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/fananchong/gotcp"
)

var gChartSession *ChartSession

func main() {
	param1 := "127.0.0.1"
	flag.StringVar(&param1, "ip", "127.0.0.1", "ip")
	flag.Parse()

	addrs := fmt.Sprintf("%s:5003", param1)
	gChartSession = &ChartSession{}
	gChartSession.Connect("127.0.0.1:3333", gChartSession)
	gChartSession.Verify()
	TcpClient(addrs)
}

func TcpClient(addrs string) {
	addr, _ := net.ResolveTCPAddr("tcp4", addrs)
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to ", conn.RemoteAddr().String())

	conn.SetNoDelay(true)
	conn.SetWriteBuffer(128 * 1024)
	conn.SetReadBuffer(128 * 1024)

	buf := gotcp.NewByteBuffer()
	msglen := 400
	var tempbuf [1024]byte
	for {
		leastlen := msglen - buf.RdSize()
		readnum, err := io.ReadAtLeast(conn, tempbuf[0:], leastlen)
		if err != nil {
			fmt.Println(err)
			return
		}
		buf.Append(tempbuf[:readnum])
		now := time.Now().UnixNano()
		for buf.RdSize() >= msglen {
			msgbuff := buf.RdBuf()
			onTcpRecv(msgbuff[:msglen], now)
			buf.RdFlip(msglen)
		}
	}
}

var (
	preTCPRecvTime int64 = 0
	delaySlice     []int64
)

func onTcpRecv(data []byte, now int64) {
	if preTCPRecvTime == 0 {
		preTCPRecvTime = now
	}

	detal := (now - preTCPRecvTime) / int64(time.Millisecond)
	preTCPRecvTime = now
	temp := []byte(fmt.Sprintf("3_%d;", detal))
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
