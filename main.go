package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	token := mustToken()
	fmt.Println("Here is the TG bot token", token)

	// tgClient = telegram.New(token)

	//fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fercher, processor)
}

func mustToken() string {
	token := flag.String("bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not sprcified")
	}

	return *token
}
