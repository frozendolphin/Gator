package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/frozendolphin/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.arguments) == 0 {
		return errors.New("no arguments were given, please enter url")
	}
	if len(cmd.arguments) > 1 {
		return errors.New("too many arguments")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("feed: %v has been followed by %v\n", feed_follow.FeedName, feed_follow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {

	if len(cmd.arguments) >= 1 {
		return errors.New("too many arguments")
	}

	following_feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for i, feed := range following_feeds {
		fmt.Printf("%v. %v - %v (%v)\n", i + 1, feed.FeedName, feed.FeedUrl, feed.UserName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	
	if len(cmd.arguments) == 0 {
		return errors.New("no arguments were given, please enter feed_url")
	}
	if len(cmd.arguments) > 1 {
		return errors.New("too many arguments")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	param := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), param)
	if err != nil {
		return err
	}

	fmt.Printf("%v unfollowed by %v", feed.Name, user.Name)

	return nil
}