package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/modules/music"
	"normalBot/internal/utils"
)

func DisconnectCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "disconnect",
		Description: "BotをVCから切断します。",
	}
}

func DisconnectHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	vc, ok := music.ExistsData(i.GuildID)
	if !ok {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。まだVCに参加していません。", Ephemeral: true})
		return
	}

	if err := vc.Connection.Disconnect(); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。", Ephemeral: true})
		log.WithFields(log.Fields{"error": err}).Error("music disconnect error")
		return
	}
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "切断できました",
		},
	})
	music.DeleteData(i.GuildID)
}
