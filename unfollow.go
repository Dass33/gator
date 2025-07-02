package main

import (
	"context"
	"fmt"

	"github.com/Dass33/gator/internal/database"
)

func unfollow(ctx context.Context, s *state, url string, user database.User) error {
	feed, err := s.db.GetFeedFromUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("Could not get feed from given url: %v", err)
	}

	feed_follow_pr := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	return s.db.DeleteFeedFollow(ctx, feed_follow_pr)
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Unfollow expects one argument, url")
	}
	ctx := context.Background()
	url := cmd.arguments[0]

	err := unfollow(ctx, s, url, user)
	if err != nil {
		return fmt.Errorf("Feed follow could not be removed: %v", err)
	}

	fmt.Printf("Feed successfuly deleted")
	return nil
}
