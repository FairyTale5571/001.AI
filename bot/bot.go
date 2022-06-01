package bot

import (
	"001.AI/config"
	"001.AI/logger"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
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

	// даем необходимые полномочия
	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	// добавляем обработчики событий
	s.AddHandler(onUserConnected)    // обработчик новых пользователей
	s.AddHandler(onUserDisconnected) // обработчик ливнувших пользователей
	s.AddHandler(onMessageHandle)    // Обработчик сообщений
	s.AddHandler(onCommandsCall)     // обработчик / команд
	s.AddHandler(onReactMessage)

	// Проверям работает ли бот
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.PrintLog("Bot is up!")
	})

	// открываем сессию
	err = s.Open()
	if err != nil {
		logger.PrintLog("Cannot open the session: %v", err)
		return
	}
	defer s.Close() // закроем сессию при завершении
	for _, elem := range s.State.Guilds {
		log.Printf("Guild: %s\n", elem.ID)
		go addRemoveCommands(elem.ID)
	}
	startRoutine()
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
