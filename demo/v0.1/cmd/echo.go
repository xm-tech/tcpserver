package cmd

import (
	"log"

	"github.com/xm-tech/tcpserver/cmd"
)

type Echo struct {
	cmd.BaseCommand
}

func (self *Echo) Exec(data ...interface{}) interface{} {
	log.Println("Echo.Exec,data=", data)
	return data
}

func (self *Echo) Name() string {
	return "Echo"
}
