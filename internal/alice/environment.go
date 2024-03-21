package alice

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

type Environment struct {
	Path    string `short:"e" long:"env" description:"Environment file path (default: workspace/environment.json)" default:"workspace/environment.json"`
	DevMode bool   `short:"d" long:"dev" description:"Start in developer mode"`

	BotToken   string `json:"token"`
	DevGuildID string `json:"guild_id"`
}

func (e *Environment) Parse() error {
	data, err := os.ReadFile(e.Path)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("environment file error")
	}
	return json.Unmarshal(data, &e)
}
