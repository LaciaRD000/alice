package commands

import (
	"github.com/bwmarrin/discordgo"
	"normalBot/internal/database"
	"normalBot/internal/utils"
	"time"
)

func StatusPanelCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "status-panel",
		Description:              "対応状況パネルを設置します。",
		DefaultMemberPermissions: &permission,
	}
}

func StatusPanelHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	m, _ := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "対応状況",
			Description: "対応可能です。",
			Color:       utils.IntParse("00FF00"),
			Image: &discordgo.MessageEmbedImage{
				// when able to response image
				URL:    "https://cdn.discordapp.com/attachments/1222881867343204473/1222893474341191751/4.png?ex=6617df4c&is=66056a4c&hm=2660dcb8bb29ea87632a09a0290f47a436eba8efe5bd89c58a2376feb78a8f52&",
				Width:  500,
				Height: 500,
			},
			Timestamp: time.Now().Format(time.DateTime),
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "",
						Style:    discordgo.SuccessButton,
						Disabled: false,
						Emoji: discordgo.ComponentEmoji{
							Name: "❤️",
						},
						CustomID: "change_status",
					},
				},
			},
		},
	})

	utils.SendReport(s, i, utils.SendMessage{Content: "作成できました。", Ephemeral: true})
	statusPanel := database.StatusPanel{
		ID:     m.ID,
		Status: true,
	}
	if err := statusPanel.Create(); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
		return
	}
}
