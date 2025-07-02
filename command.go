package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	data map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	fn, ok := c.data[cmd.name]
	if !ok {
		return fmt.Errorf("Given function had not been found")
	}
	err := fn(s, cmd)
	if err != nil {
		return fmt.Errorf("Command call errored: %v", err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.data[name] = f
}
