package bot

import (
	"001.AI/logger"
	"encoding/csv"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		// help-001
		{
			Name:        "help-001",
			Description: "Как мной пользоваться",
			Version:     "1.0",
		},
		// join
		{
			Name:        "join",
			Description: "Подключись к голосовому каналу",
			Version:     "1.0",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "voice",
					Description: "Укажите канал",
					Required:    true,
				},
			},
		},
		// disconnect
		{
			Name:        "disconnect",
			Description: "Отключиться от голосового канала",
			Version:     "1.0",
		},
		// start-record
		{
			Name:        "start-record",
			Description: "Начать запись",
			Version:     "1.0",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name-record",
					Description: "Введите название записи",
					Required:    false,
				},
			},
		},
		// stop-record
		{
			Name:        "stop-record",
			Description: "Остановить запись",
			Version:     "1.0",
		},
		// add-ticker
		{
			Name:        "add-ticker",
			Description: "Добавить тикер для ежедневного отчета",
			Version:     "1.1",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "ticker",
					Description: "Введите тикер",
					Required:    true,
				},
			},
		},
		// ticker
		{
			Name:        "ticker",
			Description: "Получить значение прямо сейчас",
			Version:     "1.1",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "ticker",
					Description: "Введите тикер",
					Required:    true,
				},
			},
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
		// set-channel-forms
		{
			Name:        "set-total-members",
			Description: "Установить канал для отображения количества пользователей",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel-option",
					Description: "Channel option",
					Required:    true,
				},
			},
		},
		// send-welcome
		{
			Name:        "send-welcome",
			Description: "send-welcome",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User Option",
					Required:    true,
				},
			},
		},
		// clear
		{
			Name:        "clear",
			Description: "удалить сообщения",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "count-messages",
					Description: "Сколько удалить",
					Required:    true,
				},
			},
		},
		{
			Name:        "get-users",
			Description: "Получить список пользователей",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"get-users":              getUsers,
		"join":                   joinVoice,
		"disconnect":             disconnectVoice,
		"start-record":           startRecord,
		"stop-record":            stopRecord,
		"clear":                  clearMessages,
		"add-ticker":             addTicker,
		"ticker":                 printTicker,
		"help-001":               help,
		"set-total-members":      setMembersChannel,
		"set-channel-forms":      setFormChannel,
		"set-verified-role":      setVerifiedRole,
		"remove-verified-role":   removeVerifiedRole,
		"set-welcome-channel":    setWelcomeChannel,
		"remove-welcome-channel": removeWelcomeChannel,
		"send-welcome": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			m := i.ApplicationCommandData().Options[0].UserValue(nil)
			sendPrivateEmbedMessage(m.ID, generateWelcomeEmbed(m))
		},
	}
)

func getUsers(s *discordgo.Session, i *discordgo.InteractionCreate) {

	err := os.Remove("users.csv")
	if err != nil {
		logger.PrintLog(err.Error())
	}

	f, err := os.OpenFile("users.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logger.PrintLog("get open file error: %s\n", err.Error())
		return
	}
	defer f.Close()

	defer sendUsersFile(i.ChannelID)

	w := csv.NewWriter(f)
	err = w.Write([]string{"id", "name", "discriminator", "joined_at"})
	if err != nil {
		logger.PrintLog("get write error: %s\n", err.Error())
		return
	}

	after := ""
	for {
		users, err := s.GuildMembers(i.GuildID, after, 1000)
		if err != nil {
			logger.PrintLog("get users error: %s\n", err.Error())
			break
		}
		for _, v := range users {
			if err := w.Write([]string{v.User.ID, v.User.Username, v.User.Discriminator, v.JoinedAt.Format("02-01-2006 15:01-05")}); err != nil {
				logger.PrintLog("get write error: %s\n", err.Error())
				return
			}
		}
		after = users[len(users)-1].User.ID
		if len(users) != 1000 {
			break
		}
	}
	w.Flush()

}

func sendUsersFile(channelId string) {

	f, err := os.OpenFile("users.csv", os.O_RDONLY, 0666)
	if err != nil {
		logger.PrintLog("get open file error: %s\n", err.Error())
		return
	}
	defer f.Close()

	_, err = s.ChannelFileSend(channelId, "users.csv", f)
	if err != nil {
		logger.PrintLog("get send error: %s\n", err.Error())
		return
	}
}
