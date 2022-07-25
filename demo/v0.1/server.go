package main

import (
	_cmd "github.com/xm-tech/tcpserver/cmd"
	"github.com/xm-tech/tcpserver/demo/v0.1/cmd"
	"github.com/xm-tech/tcpserver/server"
)

func main() {
	cmds := _cmd.NewCommands()
	cmds.AddCommand(100, &cmd.Echo{})
	cmds.AddCommand(101, &cmd.Login{})

	cmds.GetCommand(100).Exec()

	s := server.NewServer()

	s.Run()
}
