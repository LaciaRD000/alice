package commands

import (
	"github.com/bwmarrin/discordgo"
	"normalBot/internal/modules/mention"
	"normalBot/internal/utils"
)

func UnMention() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "un-mention",
		Description:              "ユーザーに対するメンションを止めます",
		DefaultMemberPermissions: &permission,
	}
}

func UnMentionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !mention.ExistsData(i.GuildID) {
		utils.SendReport(s, i, utils.SendMessage{Content: "このサーバーでは実行されていません。"})
		return
	}

	mention.DeleteData(i.GuildID)
	utils.SendReport(s, i, utils.SendMessage{Content: "送信を取りやめました。"})
}
