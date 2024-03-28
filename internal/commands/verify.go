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
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "title",
				Description: "埋め込みのタイトルを指定できます。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "description",
				Description: "埋め込みのタイトルを指定できます。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "label",
				Description: "ボタンのラベルを指定できます。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "image-url",
				Description: "埋め込みの写真を指定できます。",
				Required:    false,
			},
		},
	}
}

func VerifyHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		title       = "認証"
		description = "ロールを受け取るには認証が必要です。"
		imageURL    string
		image       = discordgo.MessageEmbedImage{}
		label       = "認証"
		verify      database.Verify
	)

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		// log.Debugf("name: %s(%T) | value: %v(%T)", option.Name, option.Name, option.Value, option.Value)
		switch option.Name {
		case "title":
			title = option.Value.(string)
		case "description":
			description = option.Value.(string)
		case "label":
			label = option.Value.(string)
		case "image-url":
			imageURL = option.Value.(string)
		case "role":
			verify.Role = option.Value.(string)
		case "type":
			verify.Type = int(option.Value.(float64))
		}
	}

	if imageURL != "" {
		image = discordgo.MessageEmbedImage{
			URL:    imageURL,
			Width:  64,
			Height: 64,
		}
	}

	m, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    label,
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
			Title:       title,
			Description: description,
			Color:       64154,
			Image:       &image,
		},
	})

	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "Verify-Panelを作成できませんでした。", Ephemeral: true})
		return
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "Verify-Panelを作成できました。", Ephemeral: true})

	verify.ID = m.ID
	if err = verify.Create(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database error")
	}
}
