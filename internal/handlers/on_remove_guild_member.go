package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func OnRemoveGuildMember(s *discordgo.Session, i *discordgo.GuildMemberRemove) {
	log.WithFields(log.Fields{"GuildID": i.GuildID}).Debug("OnRemoveGuildMember Event")

	var leave database.Leave
	if err := leave.Find("guild_id = ?", i.GuildID); err != nil {
		return
	} else if !leave.Enabled {
		return
	}
	_, _ = s.ChannelMessageSendComplex(leave.ChannelID, &discordgo.MessageSend{Embed: &discordgo.MessageEmbed{
		Title:       "サーバー脱退",
		Description: fmt.Sprintf("<@%s>さんがサーバーから脱退しました。", i.Member.User.ID),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:    i.Member.AvatarURL("64"),
			Width:  64,
			Height: 64,
		},
		Color: utils.IntParse("87cefa"),
	}})
}
