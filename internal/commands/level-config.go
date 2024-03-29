package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func LevelConfigCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "level-config",
		Description:              "levelの設定を変更できます。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "enabled",
				Description: "有効化するか選択してください。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "level-up-notice",
				Description: "レベルアップした時の通知の仕方について選択してください。",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "レベルアップしたチャンネルで通知する。",
						Value: 1,
					},
					{
						Name:  "指定されたチャンネルで通知をする。",
						Value: 2,
					},
					{
						Name:  "通知をしない。",
						Value: 3,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel",
				Description: "レベルアップした時の通知の仕方についてで指定されたチャンネルで通知をするを選択した方は指定してください。",
				Required:    false,
				ChannelTypes: []discordgo.ChannelType{
					discordgo.ChannelTypeGuildText,
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "format",
				Description: "通知を行う際のFormatを指定してください。ユーザーをメンションするには<mention>と書いてください。",
				Required:    false,
			},
		},
	}
}

func LevelConfigHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var levelConfig = database.LevelConfig{GuildID: i.GuildID}

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "enabled":
			levelConfig.Enabled = option.Value.(bool)
		case "level-up-notice":
			levelConfig.Option = int(option.Value.(float64))
		case "channel":
			levelConfig.ChannelID = option.Value.(string)
		case "format":
			levelConfig.Format = option.Value.(string)
		default:
			log.Error("not found command option | check option!!")
		}
	}

	if err := levelConfig.Update(); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
		return
	}

	utils.SendReport(s, i, utils.SendMessage{Content: "設定を変更できました。", Ephemeral: true})
}
