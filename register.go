package main

import (
    "fmt"
    "context"
    "time"

    "github.com/google/uuid"
    
    "github.com/phucfix/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
    if len(cmd.args) != 1 {
        return fmt.Errorf("usage: %s <name>", cmd.name)
    }
    name := cmd.args[0]

    // Create the new user in the database
    user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
        ID:        uuid.New(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        Name:      name,
    })
    if err != nil {
        return fmt.Errorf("Couldn't create user: %w", err)
    }

    err = s.cfg.SetUser(name)
    if err != nil {
        return fmt.Errorf("Couldn't set current user: %w", err)
    }

    fmt.Println("User created successfully:")
    printUser(user)
    return nil
}

func printUser(user database.User) {
    fmt.Printf(" * ID:        %v\n", user.ID)
    fmt.Printf(" * Name:      %v\n", user.Name) 
}

