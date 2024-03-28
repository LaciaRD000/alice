package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func TicketCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "ticket",
		Description:              "チケットパネルを作成します。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "welcome-mention",
				Description: "チケットを作成した時に作成されたチャンネルにて作成者にメンションをします。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "almost-ticket",
				Description: "同時に作成できる数を指定できます。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "welcome-message",
				Description: "チケットを作成した時に作成されたチャンネルにてメッセージを送信します。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "support-member-role",
				Description: "サポートチームのロールを設定することができます。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "category",
				Description: "指定されたカテゴリーにチケットを作成します。",
				Required:    false,
				ChannelTypes: []discordgo.ChannelType{
					discordgo.ChannelTypeGuildCategory,
				},
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

func TicketHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		title       = "チケットを作成"
		description = "チケットを作成するには以下のボタンを押してください。"
		imageURL    string
		image       = discordgo.MessageEmbedImage{}
		label       = "チケットを作成"
		ticket      database.Ticket
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
		case "welcome-mention":
			ticket.WelcomeMention = option.Value.(bool)
		case "almost-ticket":
			ticket.AlmostTicket = int(option.Value.(float64))
		case "welcome-message":
			ticket.WelcomeMessage = option.Value.(string)
		case "support-member-role":
			ticket.SupportMemberRole = option.Value.(string)
		case "category":
			ticket.Category = option.Value.(string)
		default:
			log.Error("not found command option | check option!!")
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
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						Emoji: discordgo.ComponentEmoji{
							Name: "📩",
						},
						CustomID: "create_ticket",
					},
				},
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title:       title,
			Description: description,
			Color:       255,
			Image:       &image,
		},
	})
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "Ticket-Panelを作成できませんでした。", Ephemeral: true})
		return
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "Ticket-Panelを作成できました。", Ephemeral: true})

	ticket.ID = m.ID
	ticket.UserID = i.Member.User.ID

	if err = ticket.Create(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database error")
	}
}
