package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not delete data from users table: %v", err)
	}

	fmt.Println("Data from users table successfuly deleted")

	err = s.db.ResetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Could not delete data from feeds table: %v", err)
	}

	fmt.Println("Data from feeds table successfuly deleted")

	err = s.db.ResetPosts(context.Background())
	if err != nil {
		return fmt.Errorf("Could not delete data from posts table: %v", err)
	}

	fmt.Println("Data from posts table successfuly deleted")
	return nil
}
