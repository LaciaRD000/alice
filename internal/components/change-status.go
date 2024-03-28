package components

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func ChangeStatus(s *discordgo.Session, i *discordgo.InteractionCreate) {
	perms, _ := s.State.UserChannelPermissions(i.Member.User.ID, i.ChannelID)

	if perms&discordgo.PermissionAdministrator == 0 {
		utils.SendReport(s, i, utils.SendMessage{Content: "管理者のみが変更可能です。", Ephemeral: true})
		return
	}
	var statusPanel database.StatusPanel
	if err := statusPanel.Find("id = ?", i.Message.ID); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
		return
	}

	log.Debugf("Before: %v", statusPanel)
	var embed discordgo.MessageEmbed
	if !statusPanel.Status {
		embed = discordgo.MessageEmbed{
			Title:       "対応状況",
			Description: "対応可能です。",
			Color:       utils.IntParse("00FF00"),
			Image: &discordgo.MessageEmbedImage{
				URL:    "https://cdn.discordapp.com/attachments/1222881867343204473/1222893474341191751/4.png?ex=6617df4c&is=66056a4c&hm=2660dcb8bb29ea87632a09a0290f47a436eba8efe5bd89c58a2376feb78a8f52&",
				Width:  500,
				Height: 500,
			},
		}
		statusPanel.Status = true
	} else {
		embed = discordgo.MessageEmbed{
			Title:       "対応状況",
			Description: "対応不可能です。",
			Color:       utils.IntParse("FF0000"),
			Image: &discordgo.MessageEmbedImage{
				URL:    "https://cdn.discordapp.com/attachments/1222881867343204473/1222893474567950436/3.png?ex=6617df4c&is=66056a4c&hm=2e6f0aa40655411c7e04b0da13ac332ccfd129572295209b4102168dc52ccdeb&",
				Width:  500,
				Height: 500,
			},
		}
		statusPanel.Status = false
	}
	if _, err := s.ChannelMessageEditComplex(&discordgo.MessageEdit{
		ID:      i.Message.ID,
		Channel: i.ChannelID,
		Embed:   &embed,
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
	}); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。メッセージの編集の権限が不足しています。", Ephemeral: true})
		return
	}

	log.Debugf("After:  %v", statusPanel)
	if err := statusPanel.Update(); err != nil {
		log.WithFields(log.Fields{"error": err}).Debug("database error")
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しましした。\nReason: database error"})
		return
	}

	utils.SendReport(s, i, utils.SendMessage{Content: "変更しました。", Ephemeral: true})
}
