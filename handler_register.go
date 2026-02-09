package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ankamason/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the username")
	}

	name := cmd.args[0]
	now := time.Now()

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User was created!")
	fmt.Printf("  ID: %s\n", user.ID)
	fmt.Printf("  Name: %s\n", user.Name)
	fmt.Printf("  CreatedAt: %s\n", user.CreatedAt)
	fmt.Printf("  UpdatedAt: %s\n", user.UpdatedAt)

	return nil
}
