package commands

import (
	"github.com/bwmarrin/discordgo"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func LeaveCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "leave",
		Description:              "サーバーから脱退したユーザーに対してメッセージを送ります。",
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

func LeaveHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	var leave database.Leave
	if err := leave.Find("guild_id = ?", i.GuildID); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
		return
	}
	if leave.GuildID == "" {
		// 作成されていなければ作成しよう //
		leave = database.Leave{
			GuildID:   i.GuildID,
			Enabled:   enabled,
			ChannelID: channel,
		}
		if err := leave.Create(); err != nil {
			utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
			return
		}
	} else {
		// 作成されていれば上書きしよう //
		leave = database.Leave{
			GuildID:   i.GuildID,
			Enabled:   enabled,
			ChannelID: channel,
		}

		if err := leave.Update(); err != nil {
			utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
			return
		}
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "設定を変更できました。", Ephemeral: true})
}
