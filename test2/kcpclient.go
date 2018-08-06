package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/fananchong/gotcp"
	"github.com/xtaci/kcp-go"
)

var gChartSession *ChartSession

func main() {
	ip := "127.0.0.1"
	param2 := 100
	mode := 0
	flag.IntVar(&mode, "mode", 0, "mode")
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip")
	flag.IntVar(&param2, "interval", 100, "interval")
	flag.Parse()

	var datashards int = 0
	var parity int = 0
	var port int = 5004
	var prefix string = "4"

	if mode == 1 {
		datashards = 2
		parity = 1
		port = 5005
		prefix = "5"
	}

	addrs := fmt.Sprintf("%s:%d", ip, port)
	gChartSession = &ChartSession{}
	gChartSession.Connect("127.0.0.1:3333", gChartSession)
	gChartSession.Verify()
	KcpClient(addrs, time.Duration(param2)*time.Millisecond, datashards, parity, prefix)
}

func KcpClient(addrs string, interval time.Duration, datashards, parity int, prefix string) {
	conn, err := kcp.DialWithOptions(addrs, nil, datashards, parity)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to ", conn.RemoteAddr().String())

	setParam(conn)

	msg := getmsg()

	// recv
	go func() {
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
			onKcpRecv(data[beginIndex:endIndex], now, prefix)

			beginIndex = endIndex
			goto LABEL_AGAIN
		}
	}()

	// send
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

func onKcpRecv(data []byte, now int64, prefix string) {
	if data[0] != 97 || data[400-1] != 0 {
		panic("data error!!")
	}
	if preTCPRecvTime == 0 {
		preTCPRecvTime = now
	}

	detal := (now - preTCPRecvTime) / int64(time.Millisecond)
	preTCPRecvTime = now
	temp := []byte(fmt.Sprintf("%s_%d;", prefix, detal))
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

func getmsg() []byte {
	count := 400
	var msg []byte
	for i := 0; i < count; i++ {
		msg = append(msg, 97)
	}
	msg[count-1] = 0
	return msg
}
