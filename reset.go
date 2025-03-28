package main

import (
    "context"
    "fmt"
)

// Reset the database to a blank state
func handlerReset(s *state, cmd command) error {
    if len(cmd.args) != 0 {
        return fmt.Errorf("usage: %s", cmd.name)
    }

    if err := s.db.DeleteUsers(context.Background()); err != nil {
        return fmt.Errorf("Couldn't reset the database: %w", err)
    }

    fmt.Println("Database reset successfully!")
    return nil
}
