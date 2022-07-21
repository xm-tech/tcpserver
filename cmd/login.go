package cmd

import "fmt"

// a implement of Command
type Login struct {
}

func (self *Login) Exec(data ...interface{}) interface{} {
	fmt.Println("Exec, data:")
	for _, v := range data {
		fmt.Println("v=", v)
	}
	return nil
}

func (self *Login) OnErr(err error) {
	fmt.Println(err.Error())
}
