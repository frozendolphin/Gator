package main

import (
	"errors"
)

type command struct {
	name string 
	arguments []string 
}

type commands struct {
	all_commands map[string]func(*state, command) error
}

func GetCommands() *commands {
	return &commands{
		all_commands: make(map[string]func(*state, command) error),
	}
}

func (c *commands) run(s *state, cmd command) error {
	handler, exist := c.all_commands[cmd.name]
	if !exist {
		return errors.New("command doesn't exist")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.all_commands[name] = f
}