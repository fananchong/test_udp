package main

import (
	"fmt"
	"sync"

	"github.com/fananchong/gochart"
)

type Chart struct {
	gochart.ChartTime
	raknet []int64
	kcp    []int64
	tcp    []int64
	m      sync.Mutex
}

func NewChart() *Chart {
	this := &Chart{raknet: make([]int64, 0), kcp: make([]int64, 0)}
	this.TickUnit = 100
	this.RefreshTime = DEFAULT_REFRESH_TIME
	this.SampleNum = DEFAULT_SAMPLE_NUM
	this.ChartType = "line"
	this.Title = "网络丢包测试"
	this.SubTitle = fmt.Sprintf("服务器每 %sms 发送 400byte 消息给客户端", *showtext1)
	this.YAxisText = "delay"
	this.YMax = "2000"
	this.ValueSuffix = "ms"
	this.TickLabelStep = "100"
	this.PlotLinesY = "{ color:'red', dashStyle:'longdashdot', value:100, width:1, label:{ text:'100ms', align:'left' } }"
	this.PlotLinesY += ",{ color:'red', dashStyle:'longdashdot', value:200, width:1, label:{ text:'200ms', align:'left' } }"
	return this
}

func (this *Chart) Update(now int64) map[string][]interface{} {
	datas := make(map[string][]interface{})
	this.m.Lock()
	datas["raknet"] = make([]interface{}, 0)
	for _, v := range this.raknet {
		datas["raknet"] = append(datas["raknet"], v)
	}
	datas["kcp"] = make([]interface{}, 0)
	for _, v := range this.kcp {
		datas["kcp"] = append(datas["kcp"], v)
	}
	datas["tcp"] = make([]interface{}, 0)
	for _, v := range this.tcp {
		datas["tcp"] = append(datas["tcp"], v)
	}
	this.raknet = this.raknet[:0]
	this.kcp = this.kcp[:0]
	this.tcp = this.tcp[:0]
	this.m.Unlock()
	return datas
}

func (this *Chart) AddRakNetData(v int64) {
	this.m.Lock()
	this.raknet = append(this.raknet, v)
	this.m.Unlock()
}

func (this *Chart) AddKcpData(v int64) {
	this.m.Lock()
	this.kcp = append(this.kcp, v)
	this.m.Unlock()
}

func (this *Chart) AddTcpData(v int64) {
	this.m.Lock()
	this.tcp = append(this.tcp, v)
	this.m.Unlock()
}
