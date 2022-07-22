package cmd

import "fmt"

type Echo struct {
	BaseCommand
}

func (self *Echo) Exec(data ...interface{}) interface{} {
	fmt.Println("Echo.Exec,data=", data)
	return data
}

func (self *Echo) Name() string {
	return "Echo"
}
