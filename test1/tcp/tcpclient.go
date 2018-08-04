package main

import (
	"flag"
	"fmt"
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

	echo := &Echo{}
	echo.Connect(addrs, echo)
	echo.Verify()
	echo.Send([]byte("hello"), 0)
	for {
		time.Sleep(100 * time.Second)
	}
}

type Echo struct {
	gotcp.Session
}

func (this *Echo) OnRecv(data []byte, flag byte) {
	onTcpRecv(data, time.Now().UnixNano())
}

func (this *Echo) OnClose() {
	fmt.Println("Echo.OnClose")
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
