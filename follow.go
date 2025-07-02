package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Dass33/gator/internal/database"
	"github.com/google/uuid"
)

func follow(ctx context.Context, s *state, url string, user database.User) ([]database.CreateFeedFollowRow, error) {
	feed, err := s.db.GetFeedFromUrl(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("Could not get feed from given url: %v", err)
	}

	feed_follow_pr := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	return s.db.CreateFeedFollow(ctx, feed_follow_pr)
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Follow expects one argument, url")
	}
	ctx := context.Background()
	url := cmd.arguments[0]

	feed_follow, err := follow(ctx, s, url, user)
	if err != nil {
		return fmt.Errorf("Feed follow could not be added: %v", err)
	}

	fmt.Printf("Feed: %v\n", feed_follow)
	return nil
}
