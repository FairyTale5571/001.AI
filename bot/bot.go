package bot

import (
	"001.AI/config"
	"001.AI/logger"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
)

// указатель на бота, используется только внутри пакета bot
var s *discordgo.Session

// Start функция старта бота
func Start() {
	var err error

	// создаем бота
	s, err = discordgo.New("Bot " + config.GetToken())
	if err != nil {
		logger.PrintLog(err.Error())
		return
	}

	s.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages

	s.AddHandler(onUserConnected)
	s.AddHandler(onUserDisconnected)
	s.AddHandler(onMessageHandle)
	s.AddHandler(onCommandsCall)
	s.AddHandler(onReactMessage)
	s.AddHandler(onGuildCreate)

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.PrintLog("Bot is up!")
	})

	err = s.Open()
	if err != nil {
		logger.PrintLog("Cannot open the session: %v", err)
		return
	}
	defer func(s *discordgo.Session) {
		err := s.Close()
		if err != nil {
			logger.PrintLog("Cannot close the session: %v", err)
		}
	}(s) // закроем сессию при завершении

	logger.PrintLog("Start goroutines")
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	// Конец работы
	logger.PrintLog("Gracefully shutdown\n************************************************************************\n\n")
}

func addRemoveCommands(guildId string) {
	logger.PrintLog("Init commands...")

	cmd, err := s.ApplicationCommands(s.State.User.ID, guildId)
	if err != nil {
		logger.PrintLog(err.Error())
	}

	insert := true
	for _, v := range commands {
		for _, elem := range cmd {
			if elem == v {
				insert = false
				break
			}
		}
		if !insert {
			continue
		}
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildId, v)
		if err != nil {
			logger.PrintLog("Cannot create '%v' command: %v", v.Name, err)
		}
		logger.PrintLog("Command %v created", v.Name)

	}
	s.ApplicationCommandBulkOverwrite(s.State.User.ID, guildId, commands)

	logger.PrintLog("Init commands finished")
}
