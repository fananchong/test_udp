package main

import (
	"flag"
	"fmt"
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

	var left []byte
	for {
		var tempbuff [102400]byte
		templen, err := conn.Read(tempbuff[:])
		if err != nil {
			fmt.Println(err)
			break
		}
		if templen == 0 {
			continue
		}

		data := append(left, tempbuff[0:templen]...)
		left = left[:0]
		datalen := len(data)

		beginIndex := 0
	LABEL_AGAIN:
		endIndex := findData(beginIndex, data, datalen)
		if endIndex < 0 {
			if beginIndex < datalen {
				left = append(left, data[beginIndex:datalen]...)
			}
			continue
		}

		now := time.Now().UnixNano()
		onTcpRecv(data[beginIndex:endIndex], now)

		beginIndex = endIndex
		goto LABEL_AGAIN
	}
}

func findData(beginIndex int, data []byte, datalen int) int {
	if beginIndex+400 <= datalen {
		return beginIndex + 400
	}
	return -1
}

var (
	preTCPRecvTime int64 = 0
	delaySlice     []int64
)

func onTcpRecv(data []byte, now int64) {
	if data[0] != 97 || data[400-1] != 0 {
		panic("data error!!")
	}

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
