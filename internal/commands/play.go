package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/modules/music"
	"normalBot/internal/utils"
)

func PlayCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "play",
		Description: "Voice Channelに参加し、音楽を再生してくれます。",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "url",
				Description: "再生する音楽のURL",
				Required:    true,
			},
		},
	}
}

func PlayHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userState, _ := s.State.VoiceState(i.GuildID, i.Member.User.ID)
	if userState == nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。VCに参加してからもう一度お試しください。", Ephemeral: true})
		return
	}

	var url string

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "url":
			url = option.Value.(string)
		}
	}

	vcData, ok := music.ExistsData(i.GuildID)
	if !ok { // もしサーバーに参加していなければ //
		vc, err := s.ChannelVoiceJoin(i.GuildID, userState.ChannelID, false, true)
		if err != nil {
			utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しましました。", Ephemeral: true})
			log.WithFields(log.Fields{"error": err}).Error("channel join error")
			return
		}
		vcData = music.InsertData(i.GuildID, vc, userState.ChannelID)
	}

	if vcData.ChannelID != userState.ChannelID {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。既にBotが違うチャンネルで接続中です。", Ephemeral: true})
		return
	}

	music.AppendQueue(&vcData, url)

	music.Play(&vcData)
}
