package main

import (
	"errors"
	"fmt"

)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("no arguments were given, please enter username")
	}
	if len(cmd.arguments) > 1 {
		return errors.New("too many arguments")
	}
	err := s.the_state.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Println("Username has been set as", cmd.arguments[0])
	return nil
}