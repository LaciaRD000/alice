package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/utils"
)

func ClearCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "clear",
		Description:              "指定された数だけメッセージをさかのぼって削除をします。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "almost",
				Description: "メッセージを削除する数を指定してください。",
				Required:    true,
			},
		},
	}
}

func ClearHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var almost int
	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "almost":
			almost = int(option.Value.(float64))
		default:
			log.Errorf("unknown option name: %s", option.Name)
		}
	}

	utils.SendReport(s, i, utils.SendMessage{Content: "メッセージを削除します。", Ephemeral: true, Type: discordgo.InteractionResponseChannelMessageWithSource})

	var st []*discordgo.Message

	for almost > 0 {
		if almost >= 100 {
			almost -= 100

			st, _ = s.ChannelMessages(i.ChannelID, 100, i.ID, "", "")
		} else {
			st, _ = s.ChannelMessages(i.ChannelID, almost, i.ID, "", "")
		}
		var msg []string
		for _, m := range st {
			msg = append(msg, m.ID)
		}

		if err := s.ChannelMessagesBulkDelete(i.ChannelID, msg); err != nil {
			log.WithFields(log.Fields{"error": err}).Debug("ChannelMessages Error")
			utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。", Ephemeral: true})
			return
		}
	}
}
