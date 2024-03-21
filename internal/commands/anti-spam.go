package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func AntiSpamCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "anti-spam",
		Description:              "spamからサーバーを守ります",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "mode",
				Description: "true/falseを切り替えます。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "anti-invite",
				Description: "招待リンクを送信することを禁止にします",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "maximum-mentions",
				Description: "一度にメンションする数を設定できます。設定しない場合は-1と入力してください。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "anti-duplicate",
				Description: "重複メッセージを削除します。(疑われるメッセージも含む) 設定しない場合は-1と入力してください。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "anti-raid",
				Description: "10秒のうちに参加してくるユーザーの上限を設定できます。その上限を超えたユーザーはBanされます。設定しない場合は-1と入力してください。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "maximum-lines",
				Description: "一度に送信できるメッセージの行数を設定できます。設定しない場合は-1と入力してください。",
				Required:    false,
			},
		},
	}
}

func AntiSpamHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	antiSpam := database.AntiSpam{ID: i.GuildID}

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "mode":
			antiSpam.Enabled = option.Value.(bool)
		case "anti-invite":
			antiSpam.AntiInvite = option.Value.(bool)
		case "maximum-mentions":
			antiSpam.MaximumMentions = int(option.Value.(float64))
		case "anti-duplicate":
			antiSpam.AntiDuplicate = int(option.Value.(float64))
		case "anti-raid":
			antiSpam.AntiRaid = int(option.Value.(float64))
		case "maximum-lines":
			antiSpam.MaximumLines = int(option.Value.(float64))
		default:
			log.Error("not found command option | check option!!")
		}
	}

	if err := antiSpam.Update(); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。\nReason: database error", Ephemeral: true})
		return
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "設定を変更できました。", Ephemeral: true})
}
