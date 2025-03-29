package main

import (
    "context"
    "fmt"
)

func handlerGetUsers(s *state, cmd command) error {
    if len(cmd.args) != 0 {
        return fmt.Errorf("usage: %s", cmd.name)
    }

    // Get all registered users
    users, err := s.db.GetUsers(context.Background())
    if err != nil {
        return fmt.Errorf("Couldn't get users: %w", err)
    }

    for _, user := range users {
        if user.Name == s.cfg.CurrentUsername {
            fmt.Printf("* %s (current)\n", user.Name)
            continue
        }
        fmt.Printf("* %s\n", user.Name)
    }
    return nil
}
