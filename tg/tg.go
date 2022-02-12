package tg

import (
	"001.AI/bot"
	"001.AI/config"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var tgBot *tgbotapi.BotAPI

func init() {
	tgBot = start()
	if tgBot == nil {
		panic("telegram error!\n")
	}
}

func start() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.GetTelegramToken())
	if err != nil {
		log.Fatal(err)
		return nil
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	go startRoutine()
	return bot

}

func startRoutine() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 3
	updates := tgBot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			printToDs(fmt.Sprintf(
				"{%s} [%s] %s", update.Message.Chat.UserName, update.Message.From.UserName, update.Message.Text))
		}
	}
}

func printToDs(text string) {
	bot.PrintFromTG(text)
}

func CheckTG() bool {
	return true
}
