package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Dass33/gator/internal/database"
)

const limit int32 = 2

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := limit
	if len(cmd.arguments) > 0 {
		tmp_limit, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("Could not parse the first argument as an int: %v", err)
		}
		limit = int32(tmp_limit)
	}

	posts_par := database.UserPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.UserPosts(context.Background(), posts_par)
	if err != nil {
		return fmt.Errorf("Could not fetch feeds from the database: %v", err)
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description)
	}
	return nil
}
