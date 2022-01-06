package bot

import "github.com/bwmarrin/discordgo"

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:"print-rules",
			Description: "Распечатать правила и кнокпку",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"print-rules": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			printRules(s,i)
		},
	}
)
