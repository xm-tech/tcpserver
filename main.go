package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/xm-tech/tcpserver/cmd"
)

func init() {
	log.Println("Server init...")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	login := cmd.Cmds.GetCommand(100)
	login.Exec("login")

	echo := cmd.Cmds.GetCommand(101)
	echo.Exec("hi")

	go signalHandler()
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
