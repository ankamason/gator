package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ankamason/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		parsed, err := strconv.Atoi(cmd.args[0])
		if err == nil {
			limit = parsed
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		if post.Description.Valid {
			fmt.Printf("Description: %s\n", post.Description.String)
		}
		if post.PublishedAt.Valid {
			fmt.Printf("Published: %s\n", post.PublishedAt.Time.Format("Jan 2, 2006"))
		}
		fmt.Println()
	}

	return nil
}
