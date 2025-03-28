package main

import (
    "log"
    "os"
    "database/sql"

    _ "github.com/lib/pq"

    "github.com/phucfix/gator/internal/database"
    "github.com/phucfix/gator/internal/config"
)

type state struct {
    db  *database.Queries
    cfg *config.Config
}

func main() {
    // Load in database URL to the config struct
    // Init app state, config and register cmds
    cfg, err := config.Read()
    if err != nil {
        log.Fatalf("Error reading config file: %v\n", err)
    }
    db, err := sql.Open("postgres", cfg.DBURL)
    if err != nil {
        log.Fatalf("Error opening database: %v\n", err)
    }

    // Use generated database package to create a new *database.Queries
    dbQueries := database.New(db)

    appState := state{ dbQueries, &cfg }

    appCommands := commands{ make(map[string]func(*state, command) error) }
    appCommands.register("login", handlerLogin)
    appCommands.register("register", handlerRegister)
    appCommands.register("reset", handlerReset)
    appCommands.register("users", handlerGetUsers)

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
