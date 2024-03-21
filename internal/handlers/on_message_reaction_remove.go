package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
)

func OnMessageReactionRemove(s *discordgo.Session, i *discordgo.MessageReactionRemove) {
	log.WithFields(log.Fields{"GuildID": i.GuildID, "ChannelID": i.ChannelID, "EmojiName": i.Emoji.Name}).Debug("OnMessageReactionRemove Event")

	var reaction database.Reaction
	if err := reaction.Find("id = ?", i.MessageID); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database error")
		return
	} else if reaction.Role1 == "" {
		log.Error("not found role")
		return
	}

	var emoji = map[string]string{"1️⃣": reaction.Role1, "2️⃣": reaction.Role2, "3️⃣": reaction.Role3, "4️⃣": reaction.Role4, "5️⃣": reaction.Role5, "6️⃣": reaction.Role6, "7️⃣": reaction.Role7, "8️⃣": reaction.Role8, "9️⃣": reaction.Role9, "🔟": reaction.Role10}
	if id, ok := emoji[i.Emoji.Name]; !ok {
		log.Error("not found role")
		return
	} else {
		if err := s.GuildMemberRoleRemove(i.GuildID, i.UserID, id); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("role remove error")
		}
	}
}
