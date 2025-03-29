package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/phucfix/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.name)
	}
	feedName := cmd.args[0]
	feedURL := cmd.args[1]

	// Get current user
	user, err := s.db.GetUserFromName(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("Couldn't get current user: %w", err)
	}

	// Connect the feed to the user
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:	uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feedName,
		Url: feedURL,
		UserID: user.ID,
	})

	fmt.Println("Feed added successfully:")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:        %s\n", user.Name)
}

