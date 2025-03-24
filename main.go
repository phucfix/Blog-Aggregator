package main

import (
    "fmt"
    "log"

    "github.com/phucfix/gator/internal/config"
)

func main() {
    cfg, err := config.Read()
    if err != nil {
        log.Fatalf("Error reading config: %v\n", err)
    }
    fmt.Println(cfg)

    err = cfg.SetUser("phucfix")
    if err != nil {
        log.Fatalf("Error reading config: %v\n", err)
    }
    fmt.Println(cfg)
}
