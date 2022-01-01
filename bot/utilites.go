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

func printSimpleMessage(c string,m string) (string,error) {
	msg, err := s.ChannelMessageSend(c, m)
	if err != nil {
		logger.PrintLog("error print message")
		return "", err
	}
	return msg.ID, nil
}

func generateEmbed() *discordgo.MessageEmbed {
	
	embed := &discordgo.MessageEmbed{
		URL:         "https://platform.001k.trade/",
		Type:        discordgo.EmbedTypeImage,
		Title:       "",
		Description: "",
		Timestamp:   "",
		Color:       0x9300FF,
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      nil,
		Fields:      []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{

			},
		},
	}
	return embed
}

func pingUser(id string) string {
	return fmt.Sprintf("<@%v>",id)
}