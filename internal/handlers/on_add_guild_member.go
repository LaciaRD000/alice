package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func OnAddGuildMember(s *discordgo.Session, i *discordgo.GuildMemberAdd) {
	log.WithFields(log.Fields{"GuildID": i.GuildID}).Debug("OnAddGuildMember Event")

	var welcome database.Welcome
	if err := welcome.Find("guild_id = ?", i.GuildID); err != nil {
		return
	} else if !welcome.Enabled {
		return
	}
	_, _ = s.ChannelMessageSendComplex(welcome.ChannelID, &discordgo.MessageSend{Embed: &discordgo.MessageEmbed{
		Title:       "サーバー参加",
		Description: fmt.Sprintf("<@%s>さんがサーバーに参加しました。", i.Member.User.ID),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:    i.Member.AvatarURL("64"),
			Width:  64,
			Height: 64,
		},
		Color: utils.IntParse("87cefa"),
	}})
}
