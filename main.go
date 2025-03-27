package main

import (
    "log"
    "os"

    "github.com/phucfix/gator/internal/config"
)

type state struct {
    cfg *config.Config
}

func main() {
    // Init app state, config and register cmds
    cfg, err := config.Read()
    if err != nil {
        log.Fatalf("Error reading config file: %v\n", err)
    }

    appState := state{ &cfg }
    appCommands := commands{ make(map[string]func(*state, command) error) }
    appCommands.register("login", handlerLogin)

    // Check for user command line argument
    if len(os.Args) < 2 {
        log.Fatal("Usage: cli <command> [args...]")
    }

    cmdName := os.Args[1]
    cmdArgs := os.Args[2:]

    err = appCommands.run(&appState, command{ cmdName, cmdArgs })
    if err != nil {
        log.Fatal(err)
    }
}
