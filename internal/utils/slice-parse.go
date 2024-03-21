package utils

import (
	"github.com/bwmarrin/discordgo"
)

func SliceParse(s []*discordgo.MessageEmbedField) []*discordgo.MessageEmbedField {
	for i, v := range s {
		if v == nil {
			s = s[:i]
			break
		}
		// v.Inline = true
	}
	return s
}
