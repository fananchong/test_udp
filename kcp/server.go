package main

import (
	"github.com/xtaci/kcp-go"
)

func main() {
	_, err := kcp.ListenWithOptions(":10000", nil, 10, 3)
	if err != nil {
		panic(err)
	}
}
