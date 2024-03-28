package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/utils"
)

func NukeCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "nuke",
		Description:              "チャンネルを再作成します。",
		DefaultMemberPermissions: &permission,
	}
}

func NukeHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ch, err := s.ChannelDelete(i.ChannelID)
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。チャンネルの権限がないようです。"})
		log.WithFields(log.Fields{"error": err}).Debug("ChannelDelete Error")
		return
	}

	if _, err = s.GuildChannelCreateComplex(i.GuildID, discordgo.GuildChannelCreateData{
		Name:                 ch.Name,
		Type:                 ch.Type,
		Topic:                ch.Topic,
		Bitrate:              ch.Bitrate,
		UserLimit:            ch.UserLimit,
		RateLimitPerUser:     ch.RateLimitPerUser,
		Position:             ch.Position,
		PermissionOverwrites: ch.PermissionOverwrites,
		ParentID:             ch.ParentID,
		NSFW:                 ch.NSFW,
	}); err != nil {
		return
	}
}
