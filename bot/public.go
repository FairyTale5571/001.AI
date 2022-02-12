package bot

import (
	"001.AI/database"
	"001.AI/embed"
	"001.AI/logger"
	"fmt"
)

func CreateNewForm(body map[string]string) {
	embed := embed.NewEmbed()
	embed.SetTitle("Новая заявка на регистрацию")
	embed.SetColor(0xFF0000)

	var description string
	sorted := sortMap(body)
	for _, v := range sorted {
		description += fmt.Sprintf("**%s**: ", v)
		description += fmt.Sprintf("%s\n", body[v])
	}
	embed.SetDescription(description)

	guilds := s.State.Guilds
	for _, guild := range guilds {
		channels, err := database.GetFormsChannelId(guild.ID)
		if err != nil {
			logger.PrintLog("get rule channel error: %s\n", err.Error())
			continue
		}
		if len(channels) == 0 {
			logger.PrintLog("forms channel not configured\n")
			continue
		}
		for _, elem := range channels {
			s.ChannelMessageSendEmbed(elem, embed.MessageEmbed)
		}
	}
}

func PrintFromTG(text string) {
	s.ChannelMessageSend("933046667287068732", text)
}
