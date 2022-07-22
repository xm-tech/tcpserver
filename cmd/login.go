package cmd

import "fmt"

// a implement of Command
type Login struct {
	BaseCommand
}

func (self *Login) Exec(data ...interface{}) interface{} {
	fmt.Println("Login.Exec,data=", data)
	return nil
}

func (self *Login) OnErr(err error) {
	fmt.Println("login err:", err.Error())
}

func (self *Login) Name() string {
	return "Login"
}
