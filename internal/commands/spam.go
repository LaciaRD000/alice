package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/modules/mention"
	"normalBot/internal/utils"
)

func Mention() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "spam",
		Description:              "ユーザーに対するメンションを開始します。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "送信内容を設定できます。",
				Required:    true,
			},
		},
	}
}

func MentionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var msg string

	if mention.ExistsData(i.GuildID) {
		utils.SendReport(s, i, utils.SendMessage{Content: "既にこのサーバーで実行されています", Ephemeral: true})
		return
	}

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "message":
			msg = option.Value.(string)
		default:
			log.Errorf("unknown option name: %s", option.Name)
		}
	}

	utils.SendReport(s, i, utils.SendMessage{Content: "送信を開始します。"})

	mention.InsertData(i.GuildID, i.ChannelID, msg)
	mention.SendMessage(s, i.GuildID)
}
