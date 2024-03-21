package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func OnGuildMemberUpdate(_ *discordgo.Session, _ *discordgo.GuildMemberUpdate) {
	// log.WithFields(log.Fields{"GuildID": i.GuildID, "Username": i.Member.User.Username}).Debug("GuildMemberUpdate Event")
	log.Debug("GuildMemberUpdate Event")
}
