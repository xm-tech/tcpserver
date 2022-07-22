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

var Commands *cmd.Commands

func init() {
	log.Println("Server init...")
	Commands = cmd.NewCommands()
	Commands.AddCommand(100, &cmd.Login{})
	Commands.AddCommand(101, &cmd.Echo{})
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	login := Commands.GetCommand(100)
	login.Exec("login")

	echo := Commands.GetCommand(101)
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
