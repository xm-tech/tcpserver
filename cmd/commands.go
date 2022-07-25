package cmd

import (
	"log"
	"sync"
)

var Cmds *Commands

func init() {
	Cmds = NewCommands()
	log.Println("Cmds init succ")
}

type Commands struct {
	cmds map[int]Command
	sync.Mutex
}

func NewCommands() *Commands {
	commands := &Commands{
		cmds: make(map[int]Command),
	}
	return commands
}

func (self *Commands) AddCommand(cmd int, command Command) {
	self.Lock()
	defer self.Unlock()
	self.cmds[cmd] = command
}

// 其实这个地方不用加锁，因为我们的场景不会出现并发操作
func (self *Commands) GetCommand(cmd int) Command {
	self.Lock()
	defer self.Unlock()
	return self.cmds[cmd]
}
