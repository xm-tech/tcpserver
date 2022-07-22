package cmd

import "sync"

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

func (self *Commands) GetCommand(cmd int) Command {
	self.Lock()
	defer self.Unlock()
	return self.cmds[cmd]
}
