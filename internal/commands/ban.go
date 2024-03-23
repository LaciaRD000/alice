package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/utils"
)

func BanCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionBanMembers
	return &discordgo.ApplicationCommand{
		Name:                     "ban",
		Description:              "指定されたユーザーをBanします。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "対象のユーザーを指定できます。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "理由を設定できます。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "delete-duration",
				Description: "削除をする過去のメッセージの日数を設定できます。",
				Required:    false,
			},
		},
	}
}

func BanHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		userID   string
		reason   = "なし"
		duration = -1
	)
	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "user":
			userID = option.Value.(string)
		case "reason":
			reason = option.Value.(string)
		case "delete-duration":
			duration = int(option.Value.(float64))
		}
	}

	user, err := s.GuildMember(i.GuildID, userID)
	if err != nil {
		log.Error(err)
		return
	}

	err = s.GuildBanCreateWithReason(i.GuildID, userID, reason, duration)
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。Banする権限がないか、またはその他のエラーです。", Ephemeral: true})
		log.Error(err)
		return
	}
	embeds := []*discordgo.MessageEmbed{
		{
			Title:       "サーバーからBanされました",
			Description: fmt.Sprintf("<@%s>さんがこのサーバーからBanされました。\n理由: %s", userID, reason),
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
