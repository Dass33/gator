package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Dass33/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Register expects single arugment, the username")
	}
	username := cmd.arguments[0]

	u_exists, err := user_exists(username, s)
	if err != nil {
		return fmt.Errorf("User login failed: %v", err)
	}
	if u_exists {
		return fmt.Errorf("User already exists")
	}

	usr_par := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}
	_, err = s.db.CreateUser(context.Background(), usr_par)
	if err != nil {
		return fmt.Errorf("User creation failed: %v", err)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("User login failed: %v", err)
	}

	fmt.Printf("Registered: %v\n", usr_par)
	return nil
}
