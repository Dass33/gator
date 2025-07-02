package main

import (
	"context"
	"fmt"

	"github.com/Dass33/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	feeds, err := s.db.GetUsersFeeds(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("Could not fetch feeds from the database: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %v\n", feed.Name)
	}
	return nil
}
