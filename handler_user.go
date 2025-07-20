package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/frozendolphin/Gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("no arguments were given, please enter username")
	}
	if len(cmd.arguments) > 1 {
		return errors.New("too many arguments")
	}

	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.the_state.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("Username has been set as", cmd.arguments[0])
	
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("no arguments were given, please enter username")
	}
	if len(cmd.arguments) > 1 {
		return errors.New("too many arguments")
	}
	new_user := database.CreateUserParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.arguments[0],
	}

	_, ok := s.db.GetUser(context.Background(), cmd.arguments[0])
	if ok == nil {
		return errors.New("user already exist")
	}
	
	user, err := s.db.CreateUser(context.Background(), new_user)
	if err != nil {
		return err
	}

	err = s.the_state.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("User: %v was created at %v\n", user.Name, user.CreatedAt)
	
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) >= 1 {
		return errors.New("too many arguments")
	}

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return errors.New("failed to delete all users")
	}

	return nil
}