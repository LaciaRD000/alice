package alice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/commands"
	"normalBot/internal/database"
	"normalBot/internal/handlers"
	"os"
	"os/signal"
)

type Bot struct {
	Environment Environment
	Session     *discordgo.Session
	Commands    []*discordgo.ApplicationCommand
}

func NewBot(env Environment) *Bot {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	})

	return &Bot{Environment: env}
}

func (bot *Bot) Startup() error {
	if err := database.Open("database.db"); err != nil {
		return fmt.Errorf("database error: %v", err)
	}

	if bot.Environment.DevMode {
		log.SetLevel(log.DebugLevel)
		log.Debug("start in developer mode")
	}

	err := bot.Environment.Parse()
	if err != nil {
		return err
	}

	bot.Session, err = discordgo.New("Bot " + bot.Environment.BotToken)
	if err != nil {
		log.Errorf("Invalid bot parameters: %v", err)
	}

	bot.Session.Identify.Intents |= discordgo.IntentsAll

	bot.RegisHandler()

	if err = bot.Session.Open(); err != nil {
		return err
	} // Websocket Connection To Discord
	defer func() {
		_ = bot.Session.Close() // Close
	}()

	if err = bot.CreateCommand(); err != nil {
		return err
	}

	log.Info("Press Ctrl+C to exit")

	bot.waitInterrupt()

	bot.DeleteCommand()

	return nil
}

func (bot *Bot) CreateCommand() error {
	PublicCmdList := []*discordgo.ApplicationCommand{
		commands.TicketCommand(),
		commands.VerifyCommand(),
		// commands.AntiSpamCommand(),
		commands.ShopCommand(),
		commands.HelpCommand(),
		// commands.PlayCommand():       false,
		// commands.DisconnectCommand(): false,
		commands.BanCommand(),
		commands.SpamCommand(),
		commands.UnMention(),
		commands.WelcomeCommand(),
		commands.ReactionPanelCommand(),
		commands.TimeoutCommand(),
		commands.ClearCommand(),
		commands.NukeCommand(),
		commands.StatusPanelCommand(),
		commands.LeaveCommand(),
	}

	if _, err := bot.Session.ApplicationCommandBulkOverwrite(bot.Session.State.User.ID, "", PublicCmdList); err != nil {
		return err
	} else {
		log.Debugf("Successfully override existing command")
	}

	PrivateCmdList := []*discordgo.ApplicationCommand{
		commands.LevelConfigCommand(),
	}

	for _, cmd := range PrivateCmdList {
		if c, err := bot.Session.ApplicationCommandCreate(bot.Session.State.User.ID, bot.Environment.DevGuildID, cmd); err != nil {
			log.Errorf("Failed create private command")
			return err
		} else {
			log.Debugf("Successfully create command: %v", c.Name)
			bot.Commands = append(bot.Commands, c)
		}
	}

	return nil
}

func (bot *Bot) DeleteCommand() {
	for _, v := range bot.Commands {
		err := bot.Session.ApplicationCommandDelete(bot.Session.State.User.ID, bot.Environment.DevGuildID, v.ID)
		if err != nil {
			log.Errorf("Cannot delete '%v' command: %v", v.Name, err)
		} else {
			log.Debugf("Successfully delete command: %v", v.Name)
		}
	}
}

func (bot *Bot) RegisHandler() {
	bot.Session.AddHandler(handlers.OnReady)
	bot.Session.AddHandler(handlers.OnInteraction)
	bot.Session.AddHandler(handlers.OnMessageCreate)
	bot.Session.AddHandler(handlers.OnAddGuildMember)
	bot.Session.AddHandler(handlers.OnAddMessageReaction)
	bot.Session.AddHandler(handlers.OnRemoveGuildMember)

	log.Debug("regis handler")
}

func (bot *Bot) waitInterrupt() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
