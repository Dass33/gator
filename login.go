package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Login expects single arugment, the username")
	}
	username := cmd.arguments[0]
	u_exists, err := user_exists(username, s)
	if err != nil {
		return fmt.Errorf("User login failed: %v", err)
	}
	if !u_exists {
		return fmt.Errorf("User does not exits in the database")
	}
	err = s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("User login failed: %v", err)
	}

	fmt.Printf("User %v logged in\n", username)
	return nil
}
