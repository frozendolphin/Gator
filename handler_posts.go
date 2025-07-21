package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/frozendolphin/Gator/internal/database"
)



func handlerBrowse(s *state, cmd command, user database.User) error {

	the_limit := "2"
	if len(cmd.arguments) == 1 {
		the_limit = cmd.arguments[0]
	}
	if len(cmd.arguments) > 2 {
		return errors.New("too many arguments")
	}

	limit, err := strconv.Atoi(the_limit)
	if err != nil {
		return err
	}
	
	params := database.GetPostsForUserParams {
		UserID: user.ID,
		Limit: int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}