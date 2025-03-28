package main

import (
    "context"
    "fmt"
)

func handlerLogin(s *state, cmd command) error {
    if len(cmd.args) != 1 {
        return fmt.Errorf("usage: %s <name>", cmd.name)
    }
    name := cmd.args[0]

    // Check if user in database or not
    _, err := s.db.GetUser(context.Background(), name)
    if err != nil {
        return fmt.Errorf("couldn't find user: %w", err)
    }

    if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
    }

    fmt.Println("User switched successfully!")
    return nil
}
