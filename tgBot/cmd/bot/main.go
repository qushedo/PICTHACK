package main

import (
	"log"
	"tgBot/bot"
)

func main() {
	b, err := bot.New("bot.yml")
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
}
