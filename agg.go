package main

import (
    "context"
    "fmt"
)

func handlerFetchFeed(s *state, cmd command) error {
    if len(cmd.args) != 0 {
        return fmt.Errorf("usage: %s", cmd.name)
    }
    feedURL := "https://www.wagslane.dev/index.xml"

    rssFeed, err := fetchFeed(context.Background(), feedURL)
    if err != nil {
        return fmt.Errorf("Couldn't fetch feed at %s: %w", feedURL, err)
    }

	fmt.Printf("Feed: %+v\n", rssFeed)
    return nil
}
