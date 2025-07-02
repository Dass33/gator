package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Dass33/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Login expects two arugment, name and url")
	}
	ctx := context.Background()
	name := cmd.arguments[0]
	url := cmd.arguments[1]

	feed_pr := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
	}
	feed, err := s.db.AddFeed(context.Background(), feed_pr)
	if err != nil {
		return fmt.Errorf("Feed could not be added: %v", err)
	}

	_, err = follow(ctx, s, url, user)
	if err != nil {
		return fmt.Errorf("Feed could not be followed: %v", err)
	}
	fmt.Printf("Feed: %v\n", feed)
	return nil
}
