package main

// usage: gochart --log_dir=./log -stderrthreshold 0

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
	showtext1        = flag.String("showtext1", "100", "showtext1")
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
			for {
				var b [1024]byte
				ln, err := conn.Read(b[:])
				if err != nil {
					fmt.Println(err)
					break
				}
				if ln == 0 {
					continue
				}
				nn := strings.Split(string(b[:ln]), "_")
				n, err1 := strconv.Atoi(nn[1])
				if err1 == nil {
					if nn[0] == "1" {
						g_chart.AddRakNetData(int64(n))
					} else if nn[0] == "2" {
						g_chart.AddKcpData(int64(n))
					}
				} else {
					panic("error data!")
				}
			}
		}()
	}

}
