package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func OnMessageCreate(s *discordgo.Session, i *discordgo.MessageCreate) {
	if s.State.User.ID == i.Member.User.ID {
		return
	}
	log.WithFields(log.Fields{"GuildID": i.GuildID, "Message": i.Content}).Debug("MessageCreate Event")
}
