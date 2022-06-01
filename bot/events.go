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
	collectUser(u.Member)
	logger.PrintLog("New user connected %v#%v | ID: %v", user.Username, user.Discriminator, user.ID)
	text := fmt.Sprintf("😀 Пользвователь %s#%s присоединился %s", user.Username, user.Discriminator, pingUser(user.ID))
	sendConnectMessage(u.GuildID, text)
	if u.GuildID != "534573451877416981" {
		return
	}
	sendPrivateEmbedMessage(u.User.ID, generateWelcomeEmbed(u.User))
	database.SetConnectLog(u.GuildID, user.ID, user.Username, user.Discriminator, "connected")
}

func onUserDisconnected(s *discordgo.Session, u *discordgo.GuildMemberRemove) {
	user := u.Member.User
	logger.PrintLog("User disconnected %v#%v | ID: %v", user.Username, user.Discriminator, user.ID)
	text := fmt.Sprintf("😟 Пользвователь %s#%s отсоединился %s", user.Username, user.Discriminator, pingUser(user.ID))
	sendConnectMessage(u.GuildID, text)
	if u.GuildID != "534573451877416981" {
		return
	}
	database.SetConnectLog(u.GuildID, user.ID, user.Username, user.Discriminator, "disconnected")
}

func onCommandsCall(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
		if h, ok := ticketCommands[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	case discordgo.InteractionMessageComponent:

		if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
			h(s, i)
		}
	}
}

func onReactMessage(s *discordgo.Session, i *discordgo.MessageReactionAdd) {
	reactionHandlers := map[string]func(*discordgo.Session, *discordgo.MessageReactionAdd){
		"✅": addRole,
	}
	if h, ok := reactionHandlers[i.MessageReaction.Emoji.Name]; ok {
		h(s, i)
	}
}

func addRole(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	if e.MessageReaction.GuildID != "981627353857937509" {
		return
	}
	if e.MessageReaction.Emoji.Name == "✅" {
		if e.MessageReaction.UserID == s.State.User.ID {
			return
		}
		giveVerifiedRoles(e.GuildID, e.Member.User)
	}
}

func onMessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	// if private message
	if m.GuildID == "" {
		sendPrivateMessageError(m)
	}
	if strings.HasPrefix(m.Content, config.GetPrefix()) {
		if m.Author.ID == s.State.User.ID {
			fmt.Printf("i am author\n")
			return
		}

		var vars []string
		var content string
		inputSplit := strings.Split(m.Content, " ")
		for idx := range inputSplit {
			if idx == 0 {
				content = inputSplit[idx]
			} else {
				vars = append(vars, inputSplit[idx])
			}
		}
		switch content {
		case "!print-rule":
			printRules(m.ChannelID)
			if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
				logger.PrintLog("cant delete message %s\n", err.Error())
			}
		case "!asdvc":
			printRules2(m.ChannelID)
			if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
				logger.PrintLog("cant delete message %s\n", err.Error())
			}
		case "!tickers":
			printPrices(m.GuildID, m.ChannelID)
		case "!w":
			for _, member := range getAllUsers(m.GuildID) {
				collectUser(member)
			}
		case "!e":
			sendPrivateEmbedMessage(m.Author.ID, generateWelcomeEmbed(m.Author))
			return
		}
	}
}
