package main

import (
	"log"
	"os"

	"github.com/frozendolphin/Gator/internal/config"
)

type state struct {
	the_state *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	cfg_state := state {
		the_state: &cfg,
	}

	handler := GetCommands()

	handler.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("command name required")
	}

	cmd := command{
		name: os.Args[1],
		arguments: os.Args[2:],
	}

	err = handler.run(&cfg_state, cmd)
	if err != nil {
		log.Fatalf("error occured in running command : %v", err)
	}
}