package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func OnGuildMemberUpdate(s *discordgo.Session, i *discordgo.GuildMemberUpdate) {
	// log.WithFields(log.Fields{"GuildID": i.GuildID, "Username": i.Member.User.Username}).Debug("GuildMemberUpdate Event")
	fmt.Println("hi")
}
