package cmd

import "log"

type Echo struct {
	BaseCommand
}

func (self *Echo) Exec(data ...interface{}) interface{} {
	log.Println("Echo.Exec,data=", data)
	return data
}

func (self *Echo) Name() string {
	return "Echo"
}
