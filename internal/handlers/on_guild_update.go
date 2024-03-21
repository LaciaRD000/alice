package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
)

func OnGuildUpdate(_ *discordgo.Session, i *discordgo.GuildUpdate) {
	log.WithFields(log.Fields{"GuildName": i.Name, "GuildID": i.ID}).Debug("GuildUpdate Event")
	// ToDo メンバーが参加してきたらなんかイベントが発火している。
	/*
		_, _ = s.ChannelMessageSendComplex(i.SystemChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "導入ありがとうございます。",
				Description: "``/help``からスラッシュコマンドのリストをご確認ください。",
				Color:       utils.IntParse("ff7f"),
			},
		})

		antiSpam := database.AntiSpam{ID: i.ID}
		_ = antiSpam.Update()
		welcome := database.Welcome{GuildID: i.ID}
		_ = welcome.Create()
	*/
	welcome := database.Welcome{}
	if err := welcome.Find("id = ?", i.ID); err != nil {
		return
	} else if !welcome.Enabled {
		log.Debugf("welcome false | id: %s", i.ID)
		return
	}

	/*
		_, _ = s.ChannelMessageSendComplex(welcome.ChannelID, &discordgo.MessageSend{Embed: &discordgo.MessageEmbed{
			Title:       "サーバー参加",
			Description: fmt.Sprintf("<@%s>さんがサーバーに参加しました。", ),
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    i.Member.AvatarURL("64"),
				Width:  64,
				Height: 64,
			},
		}})
	*/
}
