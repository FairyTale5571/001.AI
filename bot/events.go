
package bot

import (
	"001.AI/config"
	"001.AI/database"
	"001.AI/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func onUserConnected(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	user := u.Member.User
	logger.PrintLog("New user connected %v#%v | ID: %v",user.Username, user.Discriminator, user.ID)
	sendPrivateEmbedMessage(u.User.ID, generateWelcomeEmbed(u.User))
	text := fmt.Sprintf("😀 Пользвователь %s#%s присоединился %s",user.Username, user.Discriminator, pingUser(user.ID))
	sendMessage("927519630396891137",text)
	database.SetConnectLog(u.GuildID,user.ID,user.Username,user.Discriminator,"connected")
}

func onUserDisconnected(s *discordgo.Session, u *discordgo.GuildMemberRemove) {
	user := u.Member.User
	logger.PrintLog("User disconnected %v#%v | ID: %v",user.Username, user.Discriminator, user.ID)

	text := fmt.Sprintf("😟 Пользвователь %s#%s отсоединился %s",user.Username, user.Discriminator, pingUser(user.ID))
	sendMessage("927519630396891137",text)
	database.SetConnectLog(u.GuildID,user.ID,user.Username,user.Discriminator,"disconnected")

}

func onCommandsCall(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	case discordgo.InteractionMessageComponent:

		if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
			h(s, i)
		}
	}
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
		case "!e":
			sendPrivateEmbedMessage(m.Author.ID, generateWelcomeEmbed(m.Author))
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
