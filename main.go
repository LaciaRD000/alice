package main

import (
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/alice"
	"os"
)

func main() {
	var environment alice.Environment
	parser := flags.NewParser(&environment, flags.Default)
	if _, err := parser.Parse(); err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
		return
	}

	bot := alice.NewBot(environment)
	if err := bot.Startup(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("bot error")
	}
}
