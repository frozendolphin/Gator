package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/frozendolphin/Gator/internal/database"
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

	return nil
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

	fmt.Printf("%v:\n", rssfeed.Channel.Title)

	for i, item := range rssfeed.Channel.Item {
		fmt.Printf("%v. %v\n", i + 1, item.Title)
	}

	fmt.Printf("\n\n")

	return nil
}
