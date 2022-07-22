package cmd

import "fmt"

// the defualt implements of the Command
type BaseCommand struct {
}

func (self *BaseCommand) Exec(data ...interface{}) interface{} {
	fmt.Println(self.Name(), ".Exec()")
	return nil
}

func (self *BaseCommand) OnErr(err error) {
	fmt.Println("BaseCommand err:", err.Error())
}

func (self *BaseCommand) Name() string {
	return "BaseCommand"
}
