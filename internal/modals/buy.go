package modals

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
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

	var shop database.Shop
	if err := shop.Find("id = ?", i.Message.ID); err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: database", Ephemeral: true})
		return
	}
	channels, _ := s.GuildChannels(i.GuildID)
	if !(shop.AlmostTicket > lenChannelName(channels, i.Member.User.Username)) {
		utils.SendReport(s, i, utils.SendMessage{Content: "チケットの作成数が多すぎます。", Ephemeral: true})
		return
	}

	ch, err := s.GuildChannelCreateComplex(i.GuildID, discordgo.GuildChannelCreateData{
		Name:     i.Member.User.Username,
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: shop.Category, // Category ID
		PermissionOverwrites: []*discordgo.PermissionOverwrite{
			{
				ID:    i.Member.User.ID,
				Type:  discordgo.PermissionOverwriteTypeMember,
				Allow: discordgo.PermissionAllText,
				Deny:  0,
			},
			{
				ID:    i.GuildID,
				Type:  discordgo.PermissionOverwriteTypeRole,
				Allow: 0,
				Deny:  discordgo.PermissionAllText,
			},
		},
	})
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "エラーが発生しました。管理者にお問い合わせください。\nReason: cannot create channel", Ephemeral: true})
		log.WithFields(log.Fields{"error": err}).Debug("create shop error")
		return
	}

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

	// slash-command option
	if shop.WelcomeMention {
		_, _ = s.ChannelMessageSendComplex(ch.ID, &discordgo.MessageSend{
			Content: fmt.Sprintf("<@%s>", i.Member.User.ID),
		})
	}

	if shop.WelcomeMessage != "" {
		_, _ = s.ChannelMessageSendComplex(ch.ID, &discordgo.MessageSend{
			Content: shop.WelcomeMessage,
		})
	}

	if shop.SupportMemberRole != "" {
		_ = s.ChannelPermissionSet(ch.ID, shop.SupportMemberRole, discordgo.PermissionOverwriteTypeRole, discordgo.PermissionAllText, 0)
	}
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

func lenChannelName(channels []*discordgo.Channel, name string) (count int) {
	for _, channel := range channels {
		if channel.Name == name {
			count++
		}
	}
	return count
}
