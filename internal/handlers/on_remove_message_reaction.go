package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"time"
)

func OnRemoveMessageReaction(s *discordgo.Session, i *discordgo.MessageReactionRemove) {
	log.WithFields(log.Fields{"GuildID": i.GuildID, "ChannelID": i.ChannelID, "EmojiName": i.Emoji.Name}).Debug("OnRemoveMessageReaction Event")

	if s.State.User.ID == i.UserID {
		return
	}

	var reaction database.Reaction
	if err := reaction.Find("id = ?", i.MessageID); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database error")
		return
	} else if reaction.Role1 == "" {
		log.Error("not found role")
		return
	}

	var emoji = map[string]string{"1️⃣": reaction.Role1, "2️⃣": reaction.Role2, "3️⃣": reaction.Role3, "4️⃣": reaction.Role4, "5️⃣": reaction.Role5, "6️⃣": reaction.Role6, "7️⃣": reaction.Role7, "8️⃣": reaction.Role8, "9️⃣": reaction.Role9, "🔟": reaction.Role10}
	id, ok := emoji[i.Emoji.Name]
	if !ok {
		log.Error("not found role")
		return
	} else {
		if err := s.GuildMemberRoleRemove(i.GuildID, i.UserID, id); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("role add error")
			return
		}
		if err := s.MessageReactionRemove(i.ChannelID, i.MessageID, id, i.UserID); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("reaction remove error")
		}
	}

	m, _ := s.ChannelMessageSendReply(i.ChannelID, fmt.Sprintf("<@%s>さんの<@&%s>を削除しました。", i.UserID, id), &discordgo.MessageReference{
		MessageID: i.MessageID,
		ChannelID: i.ChannelID,
		GuildID:   i.GuildID,
	})

	time.AfterFunc(time.Second*3, func() {
		_ = s.ChannelMessageDelete(i.ChannelID, m.ID)
	})
}
