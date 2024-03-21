package utils

import "github.com/bwmarrin/discordgo"

type SendMessage struct {
	Content    string
	Embeds     []*discordgo.MessageEmbed
	Components []discordgo.MessageComponent
	Ephemeral  bool
	Type       discordgo.InteractionResponseType
}

func SendReport(s *discordgo.Session, i *discordgo.InteractionCreate, m SendMessage) {
	var msgFlag discordgo.MessageFlags
	if m.Ephemeral {
		msgFlag = discordgo.MessageFlagsEphemeral
	} else {
		msgFlag = discordgo.MessageFlagsCrossPosted
	}

	if m.Type == 0 { // No Type
		m.Type = discordgo.InteractionResponseChannelMessageWithSource
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: m.Type,
		Data: &discordgo.InteractionResponseData{
			Flags:      msgFlag,
			Content:    m.Content,
			Embeds:     m.Embeds,
			Components: m.Components,
		},
	})
}
