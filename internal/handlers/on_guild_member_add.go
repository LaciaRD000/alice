package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func OnGuildMemberUpdate(_ *discordgo.Session, i *discordgo.GuildMemberUpdate) {
	// GuildMemberUpdateはユーザーの場合追加された場合に限る。
	log.WithFields(log.Fields{"GuildID": i.GuildID, "Username": i.Member.User.Username}).Debug("OnAddGuildMember Event")
}
