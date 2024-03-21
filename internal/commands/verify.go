package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func VerifyCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "verify",
		Description:              "認証パネルを作成します",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role",
				Description: "認証した時に付与されるロールを指定できます。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "type",
				Description: "認証の方法を指定できます",
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Modalによる計算",
						Value: 1,
					},
					{
						Name:  "ボタンを押す",
						Value: 2,
					},
				},
				Required: true,
			},
		},
	}
}

func VerifyHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	m, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "認証",
						Style:    discordgo.SuccessButton,
						Disabled: false,
						Emoji: discordgo.ComponentEmoji{
							Name: "✅",
						},
						CustomID: "verify",
					},
				},
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title:       "認証",
			Description: "ロールを受け取るには認証が必要です。",
			Color:       64154,
		},
	})

	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "Verify-Panelを作成できませんでした。", Ephemeral: true})
		return
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "Verify-Panelを作成できました。", Ephemeral: true})
	verify := database.Verify{
		ID: m.ID,
	}
	options := i.ApplicationCommandData().Options
	for _, option := range options {
		// log.Debugf("name: %s(%T) | value: %v(%T)", option.Name, option.Name, option.Value, option.Value)
		switch option.Name {
		case "role":
			verify.Role = option.Value.(string)
		case "type":
			verify.Type = int(option.Value.(float64))
		}
	}
	if err = verify.Create(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database error")
	}
}
