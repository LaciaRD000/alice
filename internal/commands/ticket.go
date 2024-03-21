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
		},
	}
}

func TicketHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	m, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "チケットを作成する",
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
			Title:       "チケットを作成",
			Description: "チケットを作成するには以下のボタンを押してください。",
			Color:       255,
		},
	})
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "Ticket-Panelを作成できませんでした。", Ephemeral: true})
		return
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "Ticket-Panelを作成できました。", Ephemeral: true})

	ticket := database.Ticket{
		ID:     m.ID,
		UserID: i.Member.User.ID,
	}
	options := i.ApplicationCommandData().Options
	for _, option := range options {
		// log.Debugf("name: %s(%T) | value: %v(%T)", option.Name, option.Name, option.Value, option.Value)
		switch option.Name {
		case "welcome-mention":
			ticket.WelcomeMention = option.Value.(bool)
		case "almost-ticket":
			ticket.AlmostTicket = int(option.Value.(float64))
		case "welcome-message":
			ticket.WelcomeMessage = option.Value.(string)
		case "support-member-role":
			ticket.SupportMemberRole = option.Value.(string)
		default:
			log.Error("not found command option | check option!!")
		}
	}
	if err = ticket.Create(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database error")
	}
}
