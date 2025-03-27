package main

import (
    "fmt"
)



type command struct {
    name string
    args []string
}

type commands struct {
    commandName map[string]func(*state, command) error
}

//  This method registers a new handler function for a command name.
func (c *commands) register(name string, f func(*state,command) error) error {
    if c.commandName == nil {
        return fmt.Errorf("Command Map is not init")
    }

    c.commandName[name] = f
    return nil
}

// This method runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
    if c == nil {
        return fmt.Errorf("Command not found")
    }

    if s == nil {
        return fmt.Errorf("Application state not found")
    }

    handler, exist := c.commandName[cmd.name]
    if !exist {
        return fmt.Errorf("%s not exist", cmd.name)
    }

    return handler(s, cmd)
}
