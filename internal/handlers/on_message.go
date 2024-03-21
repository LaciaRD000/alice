package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func OnMessageCreate(_ *discordgo.Session, i *discordgo.MessageCreate) {
	log.WithFields(log.Fields{"GuildID": i.GuildID, "Message": i.Content}).Debug("MessageCreate Event")
}
