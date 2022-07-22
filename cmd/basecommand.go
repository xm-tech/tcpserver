package cmd

import (
	"log"
)

// the defualt implements of the Command
type BaseCommand struct {
}

func (self *BaseCommand) Exec(data ...interface{}) interface{} {
	log.Println(self.Name(), ".Exec()")
	return nil
}

func (self *BaseCommand) OnErr(err error) {
	log.Println("BaseCommand err:", err.Error())
}

func (self *BaseCommand) Name() string {
	return "BaseCommand"
}
