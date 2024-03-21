package components

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func ShopTicket(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			Title:    "購入",
			CustomID: "modals_buy",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							Style:       discordgo.TextInputShort,
							Label:       "買いたい商品の番号を入力してください。",
							Placeholder: "半角で入力してください。",
							MaxLength:   2,
							MinLength:   1,
							CustomID:    "goods_number",
							Required:    true,
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{discordgo.TextInput{
						Style:       discordgo.TextInputShort,
						Label:       "買いたい商品の個数を入力してください。",
						Placeholder: "半角で入力してください。",
						MaxLength:   2,
						MinLength:   1,
						CustomID:    "goods_quantity",
						Required:    true,
					},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{discordgo.TextInput{
						Style:       discordgo.TextInputShort,
						Label:       "PayPayの送金リンクを入力してください",
						Placeholder: "半角で入力してください。",
						MaxLength:   200,
						MinLength:   1,
						CustomID:    "pay_link",
						Required:    true,
					},
					},
				},
			},
		},
	})
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("modals error")
	}
}
