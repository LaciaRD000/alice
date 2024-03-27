package commands

import (
	"github.com/bwmarrin/discordgo"
)

func ClearCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "clear",
		Description:              "指定された数だけメッセージをさかのぼって削除をします。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "almost",
				Description: "メッセージを削除する数を指定してください。",
				Required:    true,
			},
		},
	}
}

func ClearHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	/*
		var almost int
		options := i.ApplicationCommandData().Options
		for _, option := range options {
			switch option.Name {
			case "almost":
				almost = int(option.Value.(float64))
			default:
				log.Errorf("unknown option name: %s", option.Name)
			}
		}
	*/
}
