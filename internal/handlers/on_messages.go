package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"normalBot/internal/database"
	"normalBot/internal/utils"
	"strconv"
	"strings"
)

func OnMessageCreate(s *discordgo.Session, i *discordgo.MessageCreate) {
	if s.State.User.ID == i.Author.ID {
		return
	}

	var levelConfig database.LevelConfig
	if err := levelConfig.Find("guild_id = ?", i.GuildID); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database error")
		return
	} else if levelConfig.Enabled {
		guildChannelID := fmt.Sprintf("%s_%s", i.GuildID, i.Author.ID)

		var userLevel database.UserLevel
		if err = userLevel.Find("guild_channel_id", guildChannelID); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("database error")
			return
		} else if userLevel.GuildChannelID == "" {
			userLevel = database.UserLevel{
				GuildChannelID: guildChannelID,
			}
			if err = userLevel.Create(); err != nil {
				log.WithFields(log.Fields{"error": err}).Error("database error")
				return
			}
		}

		if n := rand.Intn(2); n == 1 {
			userLevel.MessagesCount++
			log.Debug("カウント増加")
		} else {
		}

		if userLevel.MessagesCount >= 25 {
			userLevel.MessagesCount -= 25
			userLevel.Level++

			switch levelConfig.Option {
			case 1:
				_, _ = s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
					Embed: &discordgo.MessageEmbed{
						Title:       "レベルアップ通知",
						Description: levelMessageParse(levelConfig.Format, i.Author.ID, userLevel.Level),
						Color:       utils.IntParse("ffd700"),
					},
				})
			case 2:
				_, _ = s.ChannelMessageSendComplex(levelConfig.ChannelID, &discordgo.MessageSend{
					Embed: &discordgo.MessageEmbed{
						Title:       "レベルアップ通知",
						Description: levelMessageParse(levelConfig.Format, i.Author.ID, userLevel.Level),
						Color:       utils.IntParse("ffd700"),
					},
				})
			}
		}

		if err = userLevel.Update(); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("database error")
			return
		}
	}

	log.WithFields(log.Fields{"GuildID": i.GuildID, "Message": i.Content}).Debug("MessageCreate Event")
}

func levelMessageParse(s, ID string, lv int) string {
	s = strings.ReplaceAll(s, "<mention>", fmt.Sprintf("<@%s>", ID))
	s = strings.ReplaceAll(s, "<level>", strconv.Itoa(lv))
	return s
}
