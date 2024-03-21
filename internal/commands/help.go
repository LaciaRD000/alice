package commands

import (
	"github.com/bwmarrin/discordgo"
	"normalBot/internal/utils"
)

func HelpCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "help",
		Description: "Botについての情報を送信します。",
	}
}

func HelpHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embeds := []*discordgo.MessageEmbed{
		{
			Title:       "BotのSlash-Commandの一覧",
			Description: "Page 1",
			Color:       utils.IntParse("ffff"),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "/anti-spam",
					Value: "Spam対策をします。",
				},
				{
					Name:  "/play",
					Value: "voice channelに参加し、音楽を流してくれます。",
				},
				{
					Name:  "/shop",
					Value: "半自動販売機を設置してくれます",
				},
				{
					Name:  "/ticket",
					Value: "チケットパネルを設置してくれます。",
				},
				{
					Name:  "/verify",
					Value: "認証パネルを設置してくれます。",
				},
			},
		},
	}
	utils.SendReport(s, i, utils.SendMessage{Embeds: embeds})
}
