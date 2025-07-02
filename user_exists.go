package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func user_exists(username string, s *state) (bool, error) {
	_, err := s.db.GetUser(context.Background(), username)
	if !errors.Is(err, sql.ErrNoRows) {
		if err != nil {
			return false, fmt.Errorf("There was error retrieving username from database: %v", err)
		}
		return true, nil
	}
	return false, nil
}
