package commands

import (
	"github.com/bwmarrin/discordgo"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func WelcomeCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "welcome",
		Description:              "サーバーに参加してきたユーザーに対してメッセージを送ります。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type: discordgo.ApplicationCommandOptionChannel,
				ChannelTypes: []discordgo.ChannelType{
					discordgo.ChannelTypeGuildText,
				},
				Name:        "channel",
				Description: "メッセージが送信されるチャンネルを選択してください。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "enabled",
				Description: "有効化するか選択してください。",
				Required:    false,
			},
		},
	}
}

func WelcomeHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		channel string
		enabled = false
	)

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "channel":
			channel = option.Value.(string)
		case "enabled":
			enabled = option.Value.(bool)
		}
	}

	var welcome database.Welcome
	if err := welcome.Find("guild_id = ?", i.GuildID); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
		return
	}
	if welcome.GuildID == "" {
		// 作成されていなければ作成しよう //
		welcome = database.Welcome{
			GuildID:   i.GuildID,
			Enabled:   enabled,
			ChannelID: channel,
		}
		if err := welcome.Create(); err != nil {
			utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
			return
		}
	} else {
		// 作成されていれば上書きしよう //
		welcome = database.Welcome{
			GuildID:   i.GuildID,
			Enabled:   enabled,
			ChannelID: channel,
		}

		if err := welcome.Update(); err != nil {
			utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
			return
		}
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "設定を変更できました。", Ephemeral: true})
}
