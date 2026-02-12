package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ankamason/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("the agg handler expects a single argument: time_between_reqs (e.g., 1s, 1m, 1h)")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("couldn't get next feed to fetch: %v\n", err)
		return
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Printf("couldn't mark feed as fetched: %v\n", err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("couldn't fetch feed %s: %v\n", feed.Name, err)
		return
	}

	fmt.Printf("Found %d posts from %s\n", len(rssFeed.Channel.Item), feed.Name)

	for _, item := range rssFeed.Channel.Item {
		publishedAt := sql.NullTime{}
		if item.PubDate != "" {
			t, err := parseTime(item.PubDate)
			if err == nil {
				publishedAt = sql.NullTime{Time: t, Valid: true}
			}
		}

		description := sql.NullString{}
		if item.Description != "" {
			description = sql.NullString{String: item.Description, Valid: true}
		}

		now := time.Now()
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			fmt.Printf("couldn't create post: %v\n", err)
		}
	}
}

func parseTime(timeStr string) (time.Time, error) {
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		"2006-01-02T15:04:05Z",
		"2006-01-02",
	}

	for _, format := range formats {
		t, err := time.Parse(format, timeStr)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("couldn't parse time: %s", timeStr)
}
