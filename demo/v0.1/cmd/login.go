package cmd

import (
	"log"

	"github.com/xm-tech/tcpserver/cmd"
)

// a implement of Command
type Login struct {
	cmd.BaseCommand
}

func (self *Login) Exec(data ...interface{}) interface{} {
	log.Println("Login.Exec,data=", data)
	return nil
}

func (self *Login) OnErr(err error) {
	log.Println("login err:", err.Error())
}

func (self *Login) Name() string {
	return "Login"
}
