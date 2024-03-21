package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func OnReady(s *discordgo.Session, _ *discordgo.Ready) {
	log.Infof("bot is ready")
	err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Type: discordgo.ActivityTypeCompeting,
				Name: "Discord",
			},
		},
	})
	if err != nil {
		return
	}
	log.Debug("Set bot activity")
}
