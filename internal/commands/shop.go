package commands

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
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
				Name:        "title",
				Description: "埋め込みのタイトルの文章を設定します",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "description",
				Description: "埋め込みの説明書きの文章を設定します",
				Required:    true,
			},
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
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-9",
				Description: "商品名を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-9",
				Description: "値段を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "goods-name-10",
				Description: "商品名を設定できます",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "price-10",
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
		title       string
		description string
		fields      = make([]*discordgo.MessageEmbedField, 10)
	)

	options := i.ApplicationCommandData().Options
	for _, option := range options {
		switch {
		case option.Name == "title":
			title = option.Value.(string)
		case option.Name == "description":
			description = option.Value.(string)
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

	_, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "購入する",
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
		},
	})
	if err != nil {
		utils.SendReport(s, i, utils.SendMessage{Content: "Shop-Panelを作成できませんでした。", Ephemeral: true})
		return
	}
	utils.SendReport(s, i, utils.SendMessage{Content: "Shop-Panelを作成できました。", Ephemeral: true})
}
