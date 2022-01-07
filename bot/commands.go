package bot

import "github.com/bwmarrin/discordgo"

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:"print-rules",
			Description: "Распечатать правила и кнокпку",
			Version:       "1.0",
		},
		{
			Name:          "set-verified-role",
			Description:   "Установить роль для принявших правила",
			Version:       "1.0",
			Options:       []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role-option",
					Description: "Role option",
					Required:    true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"print-rules": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			printRules(s,i)
		},
		"set-verified-role": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			setVerifiedRole(s,i)
		},
	}
)
