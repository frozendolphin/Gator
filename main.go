package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/frozendolphin/Gator/internal/config"
	"github.com/frozendolphin/Gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db *database.Queries
	the_state *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	db, err := sql.Open("postgres", cfg.UrlDb)
	if err != nil {
		log.Fatalf("error occured opening database: %v", err)
	}

	dbQueries := database.New(db)

	cfg_state := state {
		db: dbQueries,
		the_state: &cfg,
	}

	handler := GetCommands()

	handler.register("login", handlerLogin)
	handler.register("register", handlerRegister)
	handler.register("reset", handlerReset)
	handler.register("users", handlerUsers)
	handler.register("agg", handlerAgg)
	handler.register("addfeed", handlerAddfeed)
	handler.register("feeds", handlerFeeds)

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