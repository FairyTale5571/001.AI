package bot

import (
	"001.AI/database"
	"001.AI/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"sort"
)

func sendPrivateMessage(user string, m string) {
	channel, err := s.UserChannelCreate(user)
	if err != nil {
		logger.PrintLog("Cant open private channel\n"+
			"%s\n", err.Error())
		return
	}
	_, err = s.ChannelMessageSend(channel.ID, m)
	if err != nil {
		logger.PrintLog("Cant send message in private channel\n"+
			"%s\n", err.Error())
		return
	}
}

func sendMessage(channelId string, t string) {
	_, err := s.ChannelMessageSend(channelId, t)
	if err != nil {
		logger.PrintLog("cant send message? %s", err.Error())
		return
	}
}

func sendConnectMessage(guildId string, t string) {
	channels, err := database.GetWelcomeChannelId(guildId)
	if err != nil {
		logger.PrintLog("get rule channel error: %s\n", err.Error())
	}
	if len(channels) == 0 {
		logger.PrintLog("connect messages not configured\n")
		return
	}
	for _, elem := range channels {
		sendMessage(elem, t)
	}
}

func sendPrivateEmbedMessage(user string, embed *discordgo.MessageEmbed) {
	channel, err := s.UserChannelCreate(user)
	if err != nil {
		logger.PrintLog("Cant open private channel\n"+
			"%s\n", err.Error())
		return
	}
	msg, err := s.ChannelMessageSendEmbed(channel.ID, embed)
	if err != nil {
		logger.PrintLog("Cant send message in private channel\n"+
			"%s\n", err.Error())
		return
	}

	toEdit := discordgo.NewMessageEdit(channel.ID, msg.ID)
	toEdit.Components = []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label: "ПЛАТФОРМА",
					URL:   "https://platform.001k.trade/",
					Style: discordgo.LinkButton,
					Emoji: discordgo.ComponentEmoji{
						ID: "932727415409635348",
					},
				},
				discordgo.Button{
					Label: "ЗАПОЛНИ ФОРМУ",
					URL:   "https://forms.gle/as1vGdFkANdqKFUe8",
					Style: discordgo.LinkButton,
					Emoji: discordgo.ComponentEmoji{
						Name: "ℹ",
					},
				},
				discordgo.Button{
					Label: "МЕТОДИЧКА",
					URL:   "https://paper.dropbox.com/doc/D3BhdvMwZOMiBKQwl2bVI",
					Style: discordgo.LinkButton,
					Emoji: discordgo.ComponentEmoji{
						Name: "📑",
					},
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label: "20% на Spot",
					URL:   "https://www.binance.com/ru/register?ref=EZRRJ46M",
					Style: discordgo.LinkButton,
					Emoji: discordgo.ComponentEmoji{
						ID: "904592477137293382",
					},
				},
				discordgo.Button{
					Label: "10% на Futures",
					URL:   "https://www.binance.com/ru/futures/ref/37763047",
					Style: discordgo.LinkButton,
					Emoji: discordgo.ComponentEmoji{
						ID: "904592477137293382",
					},
				},
				discordgo.Button{
					Label: "30$ для TradingView",
					URL:   "https://ru.tradingview.com/gopro/?offer_id=10&aff_id=28995",
					Style: discordgo.LinkButton,
					Emoji: discordgo.ComponentEmoji{
						ID: "904592477430886450",
					},
				},
			},
		},
	}
	if _, err := s.ChannelMessageEditComplex(toEdit); err != nil {
		logger.PrintLog("error: cant edit private message: %s\n", err.Error())
		return
	}
}

func printSimpleMessage(c string, m string) (string, error) {
	msg, err := s.ChannelMessageSend(c, m)
	if err != nil {
		logger.PrintLog("error print message")
		return "", err
	}
	return msg.ID, nil
}

func generateWelcomeEmbed(m *discordgo.User) *discordgo.MessageEmbed {

	embed := &discordgo.MessageEmbed{
		URL:   "https://platform.001k.trade/",
		Type:  discordgo.EmbedTypeImage,
		Title: "Добро пожаловать!",
		Description: "Привет, **" + m.Username + "**!\n" +
			"Рады тебя приветствовать в нашей большой команде трейдеров!\n" +
			"Даем тебе план действий⤵️\n" +
			"Первым делом тебе нужно ознакомиться с правилами, которые ты обязан соблюдать в нашем комьюнити (смотри канал **" + pingChannel("536314966391914511") + "**).\n" +
			"Также настоятельно просим не менять свой ник в discord во избежание недоразумений в процессе обучения.\n" +
			"Дальше приступай к изучению канала **" + pingChannel("846679056463429652") + "**\n" +
			"Чтобы получить доступ к нашей платформе и материалам — заполни форму по ссылке внизу. Мы сделаем тебе аккаунт и скинем данные для входа в течении 24 часов\n" +
			"Сейчас рекомендуем полностью изучить наш discord, пройдись по всем каналам, посмотри как ведется работа на сервере. \n\n" +
			"Также обрати внимание на канал " + pingChannel("846684805915869184") + ", где мы собрали самые популярные вопросы, которые у тебя могут возникнуть в процессе обучения. \n" +
			"\nКогда ты получишь доступ к материалам, приступай к просмотру видео на платформе. Рекомендуем сначала посмотреть все видео, для того чтобы хорошо ориентироваться по материалам, а уже на втором круге делать все дз. Обязательно пиши заметки и веди рабочую тетрадь. \n" +
			"\nТвоими менторами будет наша команда!\n" +
			"По поводу любых вопросов и проверки домашнего задания всегда можешь обращаться в ЛС " + "**@team**" + " (Finetiq, tema_ycpex, OS, LuckyTick, ABIL, kovalyov).\n" +
			"\nТакже, при регистрации по ссылкам ниже, ты получишь скидку на комиссию Binance, 10% на Futures и 20% на Spot\n" +
			"И 30$ бонус от TradingView при регистрации по нашей ссылке" +
			"\n\n" +
			"Дай обратную связь, чтобы я удостоверился, что тебе всё понятно!)",
		Timestamp: "",
		Color:     0x9300FF,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "У тебя все получится!",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.imgur.com/LJmTLap.png",
		},
	}
	return embed
}

func pingUser(id string) string {
	return fmt.Sprintf("<@%v>", id)
}

func pingChannel(id string) string {
	return fmt.Sprintf("<#%s>", id)
}

func pingRole(id string) string {
	return fmt.Sprintf("<@&%s>", id)
}

func sortMap(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func getAllUsers(chanel, guild string) {

	var lastId = ""
	var ret = "Список должников: \n"
	for {
		fmt.Printf("collect all users, waiting...\n")
		members, err := s.GuildMembers(guild, lastId, 1000)
		if err != nil {
			fmt.Printf("cant get all users %s\n", err.Error())
			break
		}
		lastId = members[len(members)-1].User.ID
		for _, member := range members {
			if member.JoinedAt.Unix() > member.JoinedAt.Add(Year*1).Unix() {
				continue
			}
			ret += fmt.Sprintf("Пользователь: %s#%s | Присоединился: %v",
				member.User.Username,
				member.User.Discriminator,
				member.JoinedAt.Format("2006-01-02 15:01:05"),
			)
			ret += fmt.Sprintf("\n")
		}
		if len(members) < 1000 {
			fmt.Printf("users < 1000 | finish \n")
			break
		}
	}
	fmt.Printf("%s\n", ret)
	printSimpleMessage(chanel, ret)
}

func help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	emb := &discordgo.MessageEmbed{
		Type:  discordgo.EmbedTypeImage,
		Title: "Как пользваться ботом?",
		Description: "Доступные команды:\n" +
			"**/help-001** - это меню\n" +
			"**/set-channel-forms** - установить канал для приема регистраций формы\n" +
			"**/set-verified-role** - установить роль, которая будет выдаваться при приеме правил\n" +
			"**/remove-verified-role** - уберет роль которая выдается при приеме правил\n" +
			"**/set-welcome-channel** - установить канал для логов новых подключений и/или отключений\n" +
			"**!print-rule** - распечатать правила с двумя кнопками\n" +
			"Пока что это все )",
		Color: 0x00FFFF,
		Image: &discordgo.MessageEmbedImage{URL: "https://i.imgur.com/LJmTLap.png"},
		Author: &discordgo.MessageEmbedAuthor{
			Name: "001.AI",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Developed by FairyTale#5571",
		},
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				emb,
			},
			Flags: 1 << 6,
		},
	}); err != nil {
		fmt.Printf("cant create help embed\n")
		return
	}
}

func sendImage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := i.ApplicationCommandData().Options[0].StringValue()
	fmt.Printf("ready to send image\n")
	image := img.CreateImage(text)
	if image == nil {
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Files: []*discordgo.File{
				{
					Name:        image.Name(),
					ContentType: "image",
					Reader:      image,
				},
			},
		},
	})
}
