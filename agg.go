package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Dass33/gator/internal/database"
	"github.com/google/uuid"
)

func check_if_post_stored(s *state, ctx context.Context, post_time time.Time, url, title string) (bool, error) {
	posts_par := database.GetPostParams{
		PublishedAt: post_time,
		Url:         url,
		Title:       title,
	}

	_, err := s.db.GetPost(ctx, posts_par)
	if !errors.Is(err, sql.ErrNoRows) {
		if err != nil {
			return false, fmt.Errorf("There was error retrieving posts from database: %v", err)
		}
		return true, nil
	}
	return false, nil
}

func scrapeFeeds(s *state, ctx context.Context) error {
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("Could not fetch feeds from the database: %v", err)
	}

	for _, fd := range feeds {
		err := s.db.MarkFeedFetched(ctx, fd.ID)
		if err != nil {
			return fmt.Errorf("Could not the time stamp of %v: %v", fd.Name, err)
		}
		fetched_feed, err := fetchFeed(ctx, fd.Url)
		if err != nil {
			return fmt.Errorf("Feed %v, could not be fetched: %v", fd.Name, err)
		}

		for _, post := range fetched_feed.Channel.Item {
			post_time, err := time.Parse(time.RFC3339, post.PubDate)
			b, err := check_if_post_stored(s, ctx, post_time, fd.Url, post.Title)
			if b {
				if err != nil {
					fmt.Printf("Checking is post exists failed: %v", err)
				}
				continue
			}

			if err != nil {
				return fmt.Errorf("Time could not be prased: %v", err)
			}

			posts_par := database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				PublishedAt: post_time,
				Url:         fd.Url,
				Title:       post.Title,
				Description: post.Description,
				FeedID:      fd.ID,
			}
			s.db.CreatePost(ctx, posts_par)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Aggregate takes one argument, time between requests")
	}
	time_between_reqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Please enter the time in valid format, like 10m")
	}

	fmt.Printf("Collecting data every %v\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s, context.Background())
		if err != nil {
			fmt.Printf("Feeds could not be loaded from database: %v", err)
		}
	}
}
