package main

import (
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/xm-tech/tcpserver/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	go signalHandler()

	var s server.Server
	s.Exec()
}

func signalHandler() {
	c := make(chan os.Signal, 1)
	// 监听系统中断事件: 当中断事件发生时，中断信号写到 os.Signal 类型的 chan 里面
	signal.Notify(c, os.Interrupt)
	for {
		<-c
		os.Exit(0)
	}
}
