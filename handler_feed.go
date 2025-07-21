package main

import (
	"context"
	"fmt"
	"errors"

	"github.com/frozendolphin/Gator/internal/database"
)

func handlerAddfeed(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("no arguments were given, please enter feedname and url string")
	}
	if len(cmd.arguments) == 1 {
		return errors.New("missing feedname or url string")
	}
	if len(cmd.arguments) > 2 {
		return errors.New("too many arguments")
	}

	current_username := s.the_state.UserName
	user, err := s.db.GetUser(context.Background(), current_username)
	if err != nil {
		return err
	}

	params := database.CreateFeedParams {
		Name: cmd.arguments[0],
		Url: cmd.arguments[1],
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil
}

func handlerFeeds(s *state, cmd command) error{

	if len(cmd.arguments) > 0 {
		return errors.New("too many arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for i, feed := range feeds {
		fmt.Printf("%v. %v - %v (%v)\n", i + 1, feed.Name, feed.Url, feed.Name_2)
	} 

	return nil
}