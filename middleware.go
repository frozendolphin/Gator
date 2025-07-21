package main

import (
	"context"

	"github.com/frozendolphin/Gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		
		user, err := s.db.GetUser(context.Background(), s.the_state.UserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}