package bot

import "github.com/bwmarrin/discordgo"

var (
	commands = []*discordgo.ApplicationCommand{
		// print rules
		{
			Name:        "print-rules",
			Description: "Распечатать правила и кнокпку",
			Version:     "1.0",
		},
		// set verified role
		{
			Name:        "set-verified-role",
			Description: "Установить роль для принявших правила",
			Version:     "1.0",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role-option",
					Description: "Role option",
					Required:    true,
				},
			},
		},
		// remove verified role
		{
			Name:        "remove-verified-role",
			Description: "Удалить роль для принявших правила",
			Version:     "1.0",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role-option",
					Description: "Role option",
					Required:    true,
				},
			},
		},
		// set welcome channel
		{
			Name:        "set-welcome-channel",
			Description: "Установить этот канал в качестве логов",
			Version:     "1.0",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel-option",
					Description: "Channel option",
					Required:    true,
				},
			},
		},
		// remove welcome channel
		{
			Name:        "remove-welcome-channel",
			Description: "Удалить этот канал в качестве логов",
			Version:     "1.0",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel-option",
					Description: "Channel option",
					Required:    true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"print-rules": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			printRules(s, i)
		},
		"set-verified-role": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			setVerifiedRole(s, i)
		},
		"remove-verified-role": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			removeVerifiedRole(s, i)
		},
		"set-welcome-channel": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			setWelcomeChannel(s, i)
		},
		"remove-welcome-channel": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			removeWelcomeChannel(s, i)
		},
	}
)
