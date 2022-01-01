package main

import (
	"001.AI/bot"
	"001.AI/config"
	"fmt"
	"log"
)

func main() {
	// читаем конфиг с yml файла
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("\n   ____  ____ ___  ___    ____\n  / __ \\/ __ <  / /   |  /  _/\n / / / / / / / / / /| |  / /  \n/ /_/ / /_/ / / / ___ |_/ /   \n\\____/\\____/_(_)_/  |_/___/   \n                              \n")
	log.Printf("001k AI will be started in few seconds\n")
	log.Printf("v%s\n",config.GetVersion())
	bot.Start()
	return
}