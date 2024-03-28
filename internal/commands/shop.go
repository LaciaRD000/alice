package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"normalBot/internal/database"
	"normalBot/internal/utils"
	"strconv"
	"strings"
)

func ShopCommand() *discordgo.ApplicationCommand {
	var permission int64 = discordgo.PermissionAdministrator
	return &discordgo.ApplicationCommand{
		Name:                     "shop",
		Description:              "チケットパネルを作成します。",
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-1",
				Description: "商品名を設定できます",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-1",
				Description: "値段を設定できます",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "welcome-mention",
				Description: "チケットを作成した時に作成されたチャンネルにて作成者にメンションをします。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "almost-ticket",
				Description: "同時に作成できる数を指定できます。",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "welcome-message",
				Description: "チケットを作成した時に作成されたチャンネルにてメッセージを送信します。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "support-member-role",
				Description: "サポートチームのロールを設定することができます。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "category",
				Description: "指定されたカテゴリーにチケットを作成します。",
				Required:    false,
				ChannelTypes: []discordgo.ChannelType{
					discordgo.ChannelTypeGuildCategory,
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "title",
				Description: "埋め込みのタイトルの文章を設定します",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "description",
				Description: "埋め込みの説明書きの文章を設定します",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "label",
				Description: "ボタンのラベルを指定できます。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "image-url",
				Description: "埋め込みの写真を指定できます。",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-2",
				Description: "商品名を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-2",
				Description: "値段を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-3",
				Description: "商品名を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-3",
				Description: "値段を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-4",
				Description: "商品名を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-4",
				Description: "値段を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-5",
				Description: "商品名を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-5",
				Description: "値段を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-6",
				Description: "商品名を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-6",
				Description: "値段を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-7",
				Description: "商品名を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-7",
				Description: "値段を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-8",
				Description: "商品名を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-8",
				Description: "値段を設定できます",
				Required:    false,
			},
		},
	}
}

func ShopHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	const (
		goodsIndex = 2
		priceIndex = 1
	)

	var (
		title       = "Shop Panel"
		description = "購入する場合は下のボタンを押してください。"
		imageURL    string
		image       = discordgo.MessageEmbedImage{}
		label       = "購入する"
		fields      = make([]*discordgo.MessageEmbedField, 10)
		shop        database.Shop
	)

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch {
		case option.Name == "title":
			title = option.Value.(string)
		case option.Name == "description":
			description = option.Value.(string)
		case option.Name == "label":
			label = option.Value.(string)
		case option.Name == "image-url":
			imageURL = option.Value.(string)
		case option.Name == "welcome-mention":
			shop.WelcomeMention = option.Value.(bool)
		case option.Name == "almost-ticket":
			shop.AlmostTicket = int(option.Value.(float64))
		case option.Name == "welcome-message":
			shop.WelcomeMessage = option.Value.(string)
		case option.Name == "support-member-role":
			shop.SupportMemberRole = option.Value.(string)
		case option.Name == "category":
			shop.Category = option.Value.(string)
		case strings.HasPrefix(option.Name, "goods-name"):
			index, _ := strconv.Atoi(strings.Split(option.Name, "-")[goodsIndex])
			if fields[index-1] == nil {
				fields[index-1] = &discordgo.MessageEmbedField{}
			}

			fields[index-1].Name = option.Value.(string)
		case strings.HasPrefix(option.Name, "price"):
			index, _ := strconv.Atoi(strings.Split(option.Name, "-")[priceIndex])
			if fields[index-1] == nil {
				fields[index-1] = &discordgo.MessageEmbedField{}
			}

			fields[index-1].Value = strconv.Itoa(int(option.Value.(float64)))
		default:
			log.Error("not found command option | check option!!")
		}
	}

	fields = utils.SliceParse(fields)

	if imageURL != "" {
		image = discordgo.MessageEmbedImage{
			URL:    imageURL,
			Width:  64,
			Height: 64,
		}
	}

	m, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    label,
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						Emoji: discordgo.ComponentEmoji{
							Name: "💵",
						},
						CustomID: "buy",
					},
				},
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title:       title,
			Description: description,
			Color:       255,
			Fields:      fields,
			Image:       &image,
		},
	})
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "Shop-Panelを作成できませんでした。", Ephemeral: true})
		return
	}

	shop.ID = m.ID

	if err = shop.Create(); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("database error")
		utils.SendReport(s, i, utils.SendMessage{Content: "Shop-Panelを作成できませんでした。\nReason: database error", Ephemeral: true})
		return
	}

	utils.SendReport(s, i, utils.SendMessage{Content: "Shop-Panelを作成できました。", Ephemeral: true})
}
