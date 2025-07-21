package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/frozendolphin/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("no arguments were given, please enter the duration")
	}
	if len(cmd.arguments) > 1 {
		return errors.New("too many arguments")
	}

	time_between_reqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return errors.New("time duration parsing failed")
	}

	fmt.Printf("Collecting feeds every %v\n\n", cmd.arguments[0])

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		ScrapeFeeds(s)
	}
}

func ScrapeFeeds(s *state) error {

	next_feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	params := database.MarkFeedFetchedParams {
		LastFetchedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		ID: next_feed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}

	rssfeed, err := fetchFeed(context.Background(), next_feed.Url)
	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	for _, item := range rssfeed.Channel.Item {

		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		params2 := database.CreatePostParams {
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid: true,
			},
			PublishedAt: publishedAt,
			FeedID: next_feed.ID,
		}

		_, err := s.db.CreatePost(context.Background(), params2)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint\n") {
				continue
			}
			log.Printf("Couldn't create post : %v\n", err)
			continue
		}
	}

	log.Printf("Feed %s collected, %v posts found\n\n", next_feed.Name, len(rssfeed.Channel.Item))
	return nil
}
