package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/commands"
	"normalBot/internal/components"
	"normalBot/internal/modals"
	"strings"
)

var (
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ticket": commands.TicketHandler,
		"verify": commands.VerifyHandler,
		// "anti-spam": commands.AntiSpamHandler,
		"shop": commands.ShopHandler,
		"help": commands.HelpHandler,
		// "play":       commands.PlayHandler,
		// "disconnect": commands.DisconnectHandler,
		"ban":            commands.BanHandler,
		"mention":        commands.MentionHandler,
		"un-mention":     commands.UnMentionHandler,
		"welcome":        commands.WelcomeHandler,
		"reaction-panel": commands.ReactionPanelHandler,
	}
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"create_ticket": components.CreateTicket,
		"delete_ticket": components.DeleteTicket,
		"verify":        components.Verify,
		"buy":           components.ShopTicket,
	}
)

func OnInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.WithFields(log.Fields{"Type": i.Type}).Debug("InteractionCreate Event")

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		} else {
			log.Errorf("cannot find the command: %s", i.ApplicationCommandData().Name)
		}
	case discordgo.InteractionMessageComponent:
		if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
			h(s, i)
		} else {
			log.Errorf("cannot find the component: %s", i.ApplicationCommandData().Name)
		}
	case discordgo.InteractionModalSubmit:
		switch {
		case strings.HasPrefix(i.ModalSubmitData().CustomID, "modals_verify"):
			modals.VerifyHandler(s, i)
		case i.ModalSubmitData().CustomID == "modals_buy":
			modals.BuyHandler(s, i)
		default:
			log.Errorf("cannot find the modals: %s", i.ModalSubmitData().CustomID)
		}
	}
}
