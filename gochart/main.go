package main

// usage: gochart --log_dir=./log -stderrthreshold 0

import (
	"github.com/fananchong/gochart"
)

var (
	DEFAULT_REFRESH_TIME = 1
	DEFAULT_SAMPLE_NUM   = 2 * 60 / DEFAULT_REFRESH_TIME * 10
)

var (
	g_chart *Chart = nil
)

func main() {
	g_chart = NewChart()
	s := &gochart.ChartServer{}
	s.AddChart("chart", g_chart, false)
	go func() { println(s.ListenAndServe(":8000").Error()) }()
}
