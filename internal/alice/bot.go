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
	// global: true, dev: false
	cmdList := map[*discordgo.ApplicationCommand]bool{
		commands.TicketCommand(): true,
		commands.VerifyCommand(): true,
		// commands.AntiSpamCommand(): false,
		commands.ShopCommand(): true,
		commands.HelpCommand(): true,
		// commands.PlayCommand():       false,
		// commands.DisconnectCommand(): false,
		commands.BanCommand():           true,
		commands.Mention():              true,
		commands.UnMention():            true,
		commands.WelcomeCommand():       true,
		commands.ReactionPanelCommand(): true,
		commands.TimeoutCommand():       false,
	}

	for cmd, value := range cmdList {
		var guildID string
		if !value {
			guildID = bot.Environment.DevGuildID
		}

		if c, err := bot.Session.ApplicationCommandCreate(bot.Session.State.User.ID, guildID, cmd); err != nil {
			return err
		} else {
			log.Debugf("Successfully create command: %v", cmd.Name)
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

	log.Debug("regis handler")
}

func (bot *Bot) waitInterrupt() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
