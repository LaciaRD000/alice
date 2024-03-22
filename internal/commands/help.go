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
					Name:  "/ban",
					Value: "対象のユーザーをサーバーからBanします。",
				},
				{
					Name:  "/disconnect",
					Value: "参加しているVCから切断します",
				},
				{
					Name:  "/play",
					Value: "voice channelに参加し、音楽を流してくれます。",
				},
				{
					Name:  "/reaction-panel",
					Value: "リアクションパネルを設置します。",
				},
				{
					Name:  "/shop",
					Value: "半自動販売機を設置します。",
				},
				{
					Name:  "/spam",
					Value: "指定されたメッセージをSpamします。",
				},
				{
					Name:  "/stop",
					Value: "spamを止めます。",
				},
				{
					Name:  "/ticket",
					Value: "チケットパネルを設置します。",
				},
				{
					Name:  "/verify",
					Value: "認証パネルを設置します。",
				},
				{
					Name:  "/welcome",
					Value: "サーバー参加時にメッセージを送信する設定を変更できます。",
				},
			},
		},
	}
	utils.SendReport(s, i, utils.SendMessage{Embeds: embeds})
}
