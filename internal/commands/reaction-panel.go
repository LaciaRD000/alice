package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
)

func ReactionPanelCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "reaction-panel",
		Description:              "リアクションパネルを設置します。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "title",
				Description: "埋め込みのタイトルの文章を設定します",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "description",
				Description: "埋め込みの説明書きの文章を設定します",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-1",
				Description: "パネルに追加するロールを指定できます",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-2",
				Description: "パネルに追加するロールを指定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-3",
				Description: "パネルに追加するロールを指定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-4",
				Description: "パネルに追加するロールを指定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-5",
				Description: "パネルに追加するロールを指定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-6",
				Description: "パネルに追加するロールを指定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-7",
				Description: "パネルに追加するロールを指定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-8",
				Description: "パネルに追加するロールを指定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-9",
				Description: "パネルに追加するロールを指定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-10",
				Description: "パネルに追加するロールを指定できます",
				Required:    false,
			},
		},
	}
}

func ReactionPanelHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		title       string
		description string
		fields      = make([]*discordgo.MessageEmbedField, 10)
		reaction    database.Reaction
	)

	insertFields := func(i int, option *discordgo.ApplicationCommandInteractionDataOption) {
		if fields[i-1] == nil {
			fields[i-1] = &discordgo.MessageEmbedField{}
		}
		fields[i-1].Value = fmt.Sprintf("%d. <@&%s>", i, option.Value.(string))
	}

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch option.Name {
		case "title":
			title = option.Value.(string)
		case "description":
			description = option.Value.(string)
		case "role-1":
			insertFields(1, option)
			reaction.Role1 = option.Value.(string)
		case "role-2":
			insertFields(2, option)
			reaction.Role2 = option.Value.(string)
		case "role-3":
			insertFields(3, option)
			reaction.Role3 = option.Value.(string)
		case "role-4":
			insertFields(4, option)
			reaction.Role4 = option.Value.(string)
		case "role-5":
			insertFields(5, option)
			reaction.Role5 = option.Value.(string)
		case "role-6":
			insertFields(6, option)
			reaction.Role6 = option.Value.(string)
		case "role-7":
			insertFields(7, option)
			reaction.Role7 = option.Value.(string)
		case "role-8":
			insertFields(8, option)
			reaction.Role8 = option.Value.(string)
		case "role-9":
			insertFields(9, option)
			reaction.Role9 = option.Value.(string)
		case "role-10":
			insertFields(10, option)
			reaction.Role10 = option.Value.(string)
		default:
			log.Error("not found command option | check option!!")
		}
	}

	fields = utils.SliceParse(fields)

	m, _ := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       title,
			Description: description,
			Color:       255,
			Fields:      fields,
		},
	})

	reaction.ID = m.ID

	if err := reaction.Create(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error()
		utils.SendReport(s, i, utils.SendMessage{Content: "Reaction-Panelを作成できませんでした。\nReason: database error", Ephemeral: true})
		return
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "作成中です。", Ephemeral: true})

	if err := s.MessageReactionAdd(m.ChannelID, m.ID, ""); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("reaction error")
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。権限が足りないか、その他のエラーです"})
		return
	}

	utils.SendReport(s, i, utils.SendMessage{Content: "Reaction-Panelを作成できました。", Ephemeral: true})
}
