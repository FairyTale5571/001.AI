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

	// даем необходимые полномочия
	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	// добавляем обработчики событий
	s.AddHandler(onUserConnected)		// обработчик новых пользователей
	s.AddHandler(onUserDisconnected)	// обработчик ливнувших пользователей
	s.AddHandler(onMessageHandle)		// Обработчик сообщений

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
	defer s.Close() 	// закроем при завершении

	// Удаляем и тут же их добавляем, потому что дискорд принимает изменения очень долго
	//AddRemoveCommands()

	logger.PrintLog("Start goroutines")
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	// Конец работы
	logger.PrintLog("Gracefully shutdown\n************************************************************************\n\n")

}