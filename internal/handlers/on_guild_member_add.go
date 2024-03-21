package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func OnGuildMemberUpdate(_ *discordgo.Session, _ *discordgo.GuildMemberUpdate) {
	// log.WithFields(log.Fields{"GuildID": i.GuildID, "Username": i.Member.User.Username}).Debug("GuildMemberUpdate Event")
}
