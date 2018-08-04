package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/fananchong/gotcp"
)

var (
	port     int
	interval time.Duration
	msg      []byte
)

func main() {
	param1 := 100
	flag.IntVar(&param1, "interval", 100, "interval")
	flag.Parse()

	port = 5003
	interval = time.Duration(param1) * time.Millisecond
	msg = getmsg()

	s := &gotcp.Server{}
	s.RegisterSessType(Echo{})
	s.Start(fmt.Sprintf(":%d", port))
	for {
		time.Sleep(100 * time.Second)
	}
}

type Echo struct {
	gotcp.Session
	t        *time.Ticker
	chanStop chan int
}

func (this *Echo) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		this.Verify()
		this.t = time.NewTicker(interval)
		go func() {
			for {
				select {
				case <-this.t.C:
					this.Send(msg, 0)
				case <-this.chanStop:
					return
				}
			}
		}()
	}
}

func (this *Echo) OnClose() {
	fmt.Println("Echo.OnClose")
	if this.t != nil {
		this.t.Stop()
	}
	this.chanStop <- 1
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
