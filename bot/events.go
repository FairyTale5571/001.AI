package bot

import (
	"001.AI/logger"
	"github.com/bwmarrin/discordgo"
)

func onUserConnected(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	user := u.Member.User
	logger.PrintLog("New user connected %v#%v | ID: %v",user.Username, user.Discriminator, user.ID )
}

func onUserDisconnected(s *discordgo.Session, u *discordgo.GuildMemberRemove) {
	user := u.Member.User
	logger.PrintLog("User disconnected %v#%v | ID: %v",user.Username, user.Discriminator, user.ID )
}

func onMessageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {

}
