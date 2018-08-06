package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/fananchong/gochart"
)

var (
	DEFAULT_REFRESH_TIME = 1
	DEFAULT_SAMPLE_NUM   = 2 * 60 / DEFAULT_REFRESH_TIME * 10
)

var (
	g_chart   *Chart = nil
	showtext1        = flag.String("showtext1", "服务器每 66ms 发送 400byte 消息给客户端", "showtext1")
)

func main() {
	flag.Parse()

	g_chart = NewChart()
	s := &gochart.ChartServer{}
	s.AddChart("chart", g_chart, false)
	go func() { fmt.Println(s.ListenAndServe(":8000").Error()) }()

	//
	port := 3333
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("start listen:", port)
	for {
		conn, err := lis.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("on connect. addr =", conn.RemoteAddr())

		go func() {
			var left []byte
			for {
				var tempbuff [1024]byte
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

				nn := strings.Split(string(data[beginIndex:endIndex]), "_")
				n, err1 := strconv.Atoi(nn[1])
				if err1 == nil {
					if nn[0] == "1" {
						g_chart.AddRakNetData(int64(n))
					} else if nn[0] == "2" {
						g_chart.AddKcpData(int64(n))
					} else if nn[0] == "3" {
						g_chart.AddTcpData(int64(n))
					} else if nn[0] == "4" {
						g_chart.AddK1Data(int64(n))
					} else if nn[0] == "5" {
						g_chart.AddK2Data(int64(n))
					}
				} else {
					panic("error data!")
				}
				beginIndex = endIndex + 1
				goto LABEL_AGAIN
			}
		}()
	}

}

func findData(beginIndex int, data []byte, datalen int) int {
	for i := beginIndex; i < datalen; i++ {
		if string(data[i]) == ";" {
			return i
		}
	}
	return -1
}
