package modals

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/utils"
	"strings"
)

func BuyHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	goodsNumber := i.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	goodsQuantity := i.ModalSubmitData().Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	payLink := i.ModalSubmitData().Components[2].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

	if !strings.HasPrefix(payLink, "https://pay.paypay.ne.jp/") {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。入力されたPayPayのリンクに誤りがあります。", Ephemeral: true})
		return
	}

	ch, err := s.GuildChannelCreate(i.GuildID, i.Member.User.Username, discordgo.ChannelTypeGuildText)
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: cannot create channel", Ephemeral: true})
		return
	}

	err = s.ChannelPermissionSet(ch.ID, i.GuildID, discordgo.PermissionOverwriteTypeRole, 0, discordgo.PermissionAllText)
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: cannot override channel permissions", Ephemeral: true})
		return
	}
	_ = s.ChannelPermissionSet(ch.ID, i.Member.User.ID, discordgo.PermissionOverwriteTypeMember, discordgo.PermissionAllText, 0)

	// send info
	if err = infoPanel(s, ch.ID, goodsNumber, goodsQuantity, payLink); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("send embed error")
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: fmt.Sprintf("<#%s>にチケットを作成しました", ch.ID),
		},
	})
}

func infoPanel(s *discordgo.Session, chID, goodsNumber, goodsQuantity, payLink string) (err error) {
	_, err = s.ChannelMessageSendComplex(chID, &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "チケットを削除する",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						Emoji: discordgo.ComponentEmoji{
							Name: "📩",
						},
						CustomID: "delete_ticket",
					},
				},
			},
		},
		Embeds: []*discordgo.MessageEmbed{
			{
				Color: 64154,
				Title: "お問い合わせ情報",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "商品番号",
						Value: goodsNumber,
					},
					{
						Name:  "個数",
						Value: goodsQuantity,
					},
					{
						Name:  "送金リンク",
						Value: payLink,
					},
				},
			},
			{
				Title:       "チケットを削除",
				Description: "チケットを削除するには以下のボタンを押してください。",
				Color:       255,
			},
		},
	})
	return err
}
