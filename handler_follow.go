package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ankamason/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("the follow handler expects a single argument: the feed URL")
	}

	url := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	now := time.Now()
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Printf("Feed: %s\n", feedFollow.FeedName)
	fmt.Printf("User: %s\n", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
