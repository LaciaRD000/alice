package components

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func Verify(s *discordgo.Session, i *discordgo.InteractionCreate) {
	const (
		verifyTypeModal  = 1
		verifyTypeButton = 2
	)
	var verify database.Verify
	if err := verify.Find("id = ?", i.Message.ID); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: database", Ephemeral: true})
		return
	}
	if verify.Type == verifyTypeModal {
		for _, role := range i.Member.Roles {
			if role == verify.Role {
				utils.SendReport(s, i, utils.SendMessage{Content: "既に認証済みです。", Ephemeral: true})
				return
			}
		}

		a := rand.Intn(10)
		b := rand.Intn(10)

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				Title:    "認証",
				CustomID: fmt.Sprintf("modals_verify_%d", a+b),
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								Style:       discordgo.TextInputShort,
								Label:       fmt.Sprintf("下の欄に%d+%dの答えを入力してください。", a, b),
								Placeholder: "半角で入力してください",
								MaxLength:   2,
								MinLength:   1,
								CustomID:    "answer",
								Required:    true,
							},
						},
					},
				},
			},
		})
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("modals error")
		}
	} else if verify.Type == verifyTypeButton {
		for _, role := range i.Member.Roles {
			if role == verify.Role {
				utils.SendReport(s, i, utils.SendMessage{Content: "既に認証済みです。", Ephemeral: true})
				return
			}
		}

		err := s.GuildMemberRoleAdd(i.GuildID, i.Interaction.Member.User.ID, verify.Role)
		if err != nil {
			utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: cannot add role", Ephemeral: true})
			return
		}
		utils.SendReport(s, i, utils.SendMessage{Content: "認証が完了しました。", Ephemeral: true})
	} else {
		log.Error("not found verify type")
	}
}
