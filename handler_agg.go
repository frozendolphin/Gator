package main

import (
	"context"
	"encoding/json"
	"fmt"
	"errors"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	jsonData, err := json.MarshalIndent(feed, "", "  ")
	if err != nil {
		return errors.New("marshal indent failed")
	}
	fmt.Println(string(jsonData))

	return nil
}