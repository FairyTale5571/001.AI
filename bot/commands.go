package bot

import "github.com/bwmarrin/discordgo"

var (
	commands = []*discordgo.ApplicationCommand{
		// print rules
		{
			Name:        "help-001",
			Description: "Как мной пользоваться",
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
		// set-channel-forms
		{
			Name:        "set-channel-forms",
			Description: "Установить канал для логов с формы регистрации",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel-option",
					Description: "Channel option",
					Required:    true,
				},
			},
		},
		// set-channel-crypto
		{
			Name:        "set-channel-crypto",
			Description: "Установить канал для детекта цены BTC-USDT",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel-option",
					Description: "Channel option",
					Required:    true,
				},
			},
		},
		{
			Name:        "create-ticket",
			Description: "Create Ticket",
		},
		{
			Name:        "close-ticket",
			Description: "Close ticket",
		},
		{
			Name:        "add-user-to-ticket",
			Description: "add user to ticket",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User Option",
					Required:    true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"help-001": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			help(s, i)
		},
		"set-channel-forms": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			setFormChannel(s, i)
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
