package modals

import (
	"github.com/bwmarrin/discordgo"
	"normalBot/internal/database"
	"normalBot/internal/utils"
	"strings"
)

func VerifyHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	const answerIndex = 2

	modalAnswer := strings.Split(i.ModalSubmitData().CustomID, "_")[answerIndex]
	userAnswer := i.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	if modalAnswer != userAnswer {
		utils.SendReport(s, i, utils.SendMessage{Content: "答えが違います。", Ephemeral: true})
		return
	}

	var verify database.Verify

	if err := verify.Find("id = ?", i.Message.ID); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: database", Ephemeral: true})
		return
	}

	err := s.GuildMemberRoleAdd(i.GuildID, i.Interaction.Member.User.ID, verify.Role)
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: cannot add role", Ephemeral: true})
		return
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "認証が完了しました。", Ephemeral: true})
}
