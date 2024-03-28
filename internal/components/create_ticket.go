package components

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func CreateTicket(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var ticket database.Ticket
	if err := ticket.Find("id = ?", i.Message.ID); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: database", Ephemeral: true})
		return
	}
	channels, _ := s.GuildChannels(i.GuildID)
	if !(ticket.AlmostTicket > lenChannelName(channels, i.Member.User.Username)) {
		utils.SendReport(s, i, utils.SendMessage{Content: "チケットの作成数が多すぎます。", Ephemeral: true})
		return
	}

	ch, err := s.GuildChannelCreateComplex(i.GuildID, discordgo.GuildChannelCreateData{
		Name:     i.Member.User.Username,
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: ticket.Category, // Category ID
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:    i.Member.User.ID,
				Type:  discordgo.PermissionOverwriteTypeMember,
				Allow: discordgo.PermissionAllText,
				Deny:  0,
			},
			{
				ID:    i.GuildID,
				Type:  discordgo.PermissionOverwriteTypeRole,
				Allow: 0,
				Deny:  discordgo.PermissionAllText,
			},
		},
	})
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: cannot create channel", Ephemeral: true})
		log.WithFields(log.Fields{"error": err}).Debug("create ticket error")
		return
	}

	// send delete panel embed
	if err = deletePanel(s, ch.ID); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("send embed error")
	}

	// slash-command option
	if ticket.WelcomeMention {
		_, _ = s.ChannelMessageSendComplex(ch.ID, &discordgo.MessageSend{
			Content: fmt.Sprintf("<@%s>", i.Member.User.ID),
		})
	}

	if ticket.WelcomeMessage != "" {
		_, _ = s.ChannelMessageSendComplex(ch.ID, &discordgo.MessageSend{
			Content: ticket.WelcomeMessage,
		})
	}

	if ticket.SupportMemberRole != "" {
		_ = s.ChannelPermissionSet(ch.ID, ticket.SupportMemberRole, discordgo.PermissionOverwriteTypeRole, discordgo.PermissionAllText, 0)
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: fmt.Sprintf("<#%s>にチケットを作成しました", ch.ID),
		},
	})
}

func deletePanel(s *discordgo.Session, chID string) (err error) {
	_, err = s.ChannelMessageSendComplex(chID, &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "チケットを削除する",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						Emoji: discordgo.ComponentEmoji{
							Name: "📩",
						},
						CustomID: "delete_ticket",
					},
				},
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title:       "チケットを削除",
			Description: "チケットを削除するには以下のボタンを押してください。",
			Color:       255,
		},
	})
	return err
}

func lenChannelName(channels []*discordgo.Channel, name string) (count int) {
	for _, channel := range channels {
		if channel.Name == name {
			count++
		}
	}
	return count
}
