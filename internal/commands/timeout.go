package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/utils"
	"time"
)

func TimeoutCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionModerateMembers
	return &discordgo.ApplicationCommand{
		Name:                     "timeout",
		Description:              "指定されたユーザーをTimeoutします。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "対象のユーザーを指定します。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "until",
				Description: "タイムアウトの解除するまでの期間を指定してください。（28日まで）例）1h2m3s",
				Required:    true,
			},
		},
	}
}

func TimeoutHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		userID string
		until  string
	)

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "user":
			userID = option.Value.(string)
		case "until":
			until = option.Value.(string)
		}
	}

	user, err := s.GuildMember(i.GuildID, userID)
	if err != nil {
		log.Error(err)
		return
	}

	now := time.Now()
	duration, _ := time.ParseDuration(until)
	now = now.Add(duration)

	err = s.GuildMemberTimeout(i.GuildID, userID, &now)
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。Timeoutする権限がないか、またはその他のエラーです。", Ephemeral: true})
		log.Error(err)
		return
	}
	embeds := []*discordgo.MessageEmbed{
		{
			Title:       "サーバーからTimeoutされました",
			Description: fmt.Sprintf("<@%s>さんがこのサーバーからTimeoutされました。", userID),
			Color:       utils.IntParse("0000FF"),
			Author: &discordgo.MessageEmbedAuthor{
				Name:    fmt.Sprintf("実行者: %s", i.Member.User.Username),
				IconURL: i.Member.AvatarURL("128"),
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    user.AvatarURL("64"),
				Width:  64,
				Height: 64,
			},
		},
	}
	utils.SendReport(s, i, utils.SendMessage{Embeds: embeds})
}
