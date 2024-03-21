package components

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func DeleteTicket(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, err := s.ChannelDelete(i.ChannelID)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("delete_channel error")
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "エラーが発生しました。管理者にお問い合わせください。",
			},
		})
	}
}
