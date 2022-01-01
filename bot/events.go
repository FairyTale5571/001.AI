
package bot

import (
	"001.AI/config"
	"001.AI/logger"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func onUserConnected(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	user := u.Member.User
	logger.PrintLog("New user connected %v#%v | ID: %v",user.Username, user.Discriminator, user.ID )
	sendPrivateMessage(user.ID,
		"Привет," + pingUser(user.ID) + "!\n" +
			"Это бот пана Киевского, царя криптовалютного мира!\n" +
			"Если ты хочешь быть таким же классным как Пан Сергей или успешным как Тёма\n" +
			"Тогда заполни эту форму **ссылка тут будет** и мы начнем обучение как только так сразу" +
			"")
}

func onUserDisconnected(s *discordgo.Session, u *discordgo.GuildMemberRemove) {
	user := u.Member.User
	logger.PrintLog("User disconnected %v#%v | ID: %v",user.Username, user.Discriminator, user.ID )
	sendPrivateMessage(user.ID,
		"Ей, " + pingUser(user.ID) + "! Ты куда?\n" +
			"Это бот пана Киевского, царя криптовалютного мира!\n" +
			"Если у тебя что то не получалось, то возвращайся, мы тебе поможем!" +
			"")
}

func onMessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.GetPrefix()) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		var vars []string
		var content string
		inputSplit := strings.Split(m.Content, " ")
		for idx := range inputSplit {
			if idx == 0 {
				content = inputSplit[idx]
			}else{
				vars = append(vars, inputSplit[idx])
			}
		}
		switch content {
			case "!help":
				printSimpleMessage(m.ChannelID, "Привет," + pingUser(m.Author.ID) + "!" +
					" Это бот пана Киевского, царя криптовалютного мира!")
		case "!s":
			sendPrivateMessage(m.Author.ID,
				"Привет," + pingUser(m.Author.ID) + "!\n" +
				"Это бот пана Киевского, царя криптовалютного мира!\n" +
				"Если ты хочешь быть таким же классным как Пан Сергей или успешным как Тёма\n" +
				"Тогда заполни эту форму **ссылка тут будет** и мы начнем обучение как только так сразу" +
				"")
		}
	}
}
