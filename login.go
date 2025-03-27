package main

import (
    "fmt"
)

func handlerLogin(s *state, cmd command) error {
    if len(cmd.args) != 1 {
        return fmt.Errorf("usage: %s <name>", cmd.name)
    }
    name := cmd.args[0]

    if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
    }

    fmt.Println("User switched successfully!")
    return nil
}
