package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not fetch users from the database: %v", err)
	}

	for _, usr := range users {
		fmt.Printf("* %v", usr.Name)
		if s.config.CurrentUserName == usr.Name {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
	return nil
}
