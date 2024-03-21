package mention

import "github.com/bwmarrin/discordgo"

type data struct {
	channelID string
	message   string
}

var dict = make(map[string]data)

func InsertData(guildID, channelID, message string) {
	dict[guildID] = data{
		channelID: channelID,
		message:   message,
	}
}

func ExistsData(guildID string) (ok bool) {
	_, ok = dict[guildID]
	return ok
}

func DeleteData(guildID string) {
	delete(dict, guildID)
}

func SendMessage(s *discordgo.Session, guildID string) {
	if !ExistsData(guildID) {
		return
	}

	_, _ = s.ChannelMessageSendComplex(dict[guildID].channelID, &discordgo.MessageSend{
		Content: dict[guildID].message,
	})

	SendMessage(s, guildID)
}
