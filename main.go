package main

import (
	"flag"
	"log"
)

func main() {
	log.Print("service started")
	log.Printf("token is %s", mustToken())
}

func mustToken() string {
	token := flag.String("bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		panic("token is empty")
	}

	return *token
}
