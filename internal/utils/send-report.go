package utils

import "github.com/bwmarrin/discordgo"

type SendMessage struct {
	Content   string
	Embeds    []*discordgo.MessageEmbed
	Ephemeral bool
	Type      discordgo.InteractionResponseType
}

func SendReport(s *discordgo.Session, i *discordgo.InteractionCreate, m SendMessage) {
	var msgFlag discordgo.MessageFlags
	if m.Ephemeral {
		msgFlag = discordgo.MessageFlagsEphemeral
	} else {
		msgFlag = discordgo.MessageFlagsCrossPosted
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   msgFlag,
			Content: m.Content,
			Embeds:  m.Embeds,
		},
	})
}
