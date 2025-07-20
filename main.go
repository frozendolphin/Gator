package main

import (
	"fmt"
	"log"

	"github.com/frozendolphin/Gator/internal/config"
)

func main() {
	cred, err := config.Read()
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}
	cred.SetUser("frozendolphin")
	cred, err = config.Read()
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}
	fmt.Printf("url: %v\n", cred.UrlDb)
	fmt.Printf("username: %v\n", cred.UserName)
}