package bot

import (
	"001.AI/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func sendPrivateMessage(user string, m string) {
	channel, err := s.UserChannelCreate(user)
	if err != nil {
		logger.PrintLog("Cant open private channel\n" +
			"%s\n",err.Error())
		return
	}
	_, err = s.ChannelMessageSend(channel.ID, m)
	if err != nil {
		logger.PrintLog("Cant send message in private channel\n" +
			"%s\n",err.Error())
		return
	}
}

func sendMessage(channelId string, t string) {
	_, err := s.ChannelMessageSend(channelId, t)
	if err != nil {
		logger.PrintLog("cant send message? %s",err.Error())
		return
	}
}

func sendPrivateEmbedMessage(user string, embed *discordgo.MessageEmbed) {
	channel, err := s.UserChannelCreate(user)
	if err != nil {
		logger.PrintLog("Cant open private channel\n" +
			"%s\n",err.Error())
		return
	}
	_, err = s.ChannelMessageSendEmbed(channel.ID, embed)
	if err != nil {
		logger.PrintLog("Cant send message in private channel\n" +
			"%s\n",err.Error())
		return
	}
}

func printSimpleMessage(c string,m string) (string,error) {
	msg, err := s.ChannelMessageSend(c, m)
	if err != nil {
		logger.PrintLog("error print message")
		return "", err
	}
	return msg.ID, nil
}

func generateWelcomeEmbed(m *discordgo.User) *discordgo.MessageEmbed {
	
	embed := &discordgo.MessageEmbed{
		URL:         "https://platform.001k.trade/",
		Type:        discordgo.EmbedTypeImage,
		Title:       "Добро пожаловать!",
		Description: "Привет, **"+m.Username+"**!\n" +
			"Рады тебя приветствовать в нашей большой команде трейдеров!\n" +
			"Даем тебе план действий⤵️\n" +
			"Первым делом тебе нужно ознакомиться с правилами, которые ты обязан соблюдать в нашем комьюнити (смотри канал **#правила**).\n" +
			"Также настоятельно просим не менять свой ник в discord во избежание недоразумений в процессе обучения.\n" +
			"Дальше приступай к изучению канала **#как-начать**\n" +
			"Чтобы получить доступ к нашей платформе и материалам — заполни эту таблицу. Мы сделаем тебе аккаунт и скинем данные для входа.",
		Timestamp:   "",
		Color:       0x9300FF,
		Footer:      &discordgo.MessageEmbedFooter{
			Text:         "У тебя все получится!",
		},
		Thumbnail:   &discordgo.MessageEmbedThumbnail{
			URL:      "https://i.imgur.com/LJmTLap.png",
		},
		Fields:      []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Платформа",
				Value:  "[Платформа](https://platform.001k.trade/)",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Заполни форму",
				Value:  "[Форма](https://forms.gle/as1vGdFkANdqKFUe8)",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Методичка",
				Value:  "[Методичка](https://paper.dropbox.com/doc/D3BhdvMwZOMiBKQwl2bVI)",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "20% на Spot",
				Value:  "[Binance Spot](https://www.binance.com/ru/register?ref=EZRRJ46M)",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "10% на Futures",
				Value:  "[Binance Futures](https://www.binance.com/ru/futures/ref/37763047)",
				Inline: false,
			},
		},
	}
	return embed
}

func pingUser(id string) string {
	return fmt.Sprintf("<@%v>",id)
}
