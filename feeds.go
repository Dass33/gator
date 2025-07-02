package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Could not fetch feeds from the database: %v", err)
	}

	for _, fd := range feeds {
		fmt.Printf("* %v\n", fd.Name)
		fmt.Printf("\turl: %v\n", fd.Url)
	}
	return nil
}
