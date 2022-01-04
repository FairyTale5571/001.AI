package bot

import (
	"001.AI/logger"
	"github.com/bwmarrin/discordgo"
)

func isAdmin(user *discordgo.User, channel string) bool {
	_, err := s.UserChannelPermissions(user.ID, channel)
	if err != nil {
		logger.PrintLog("cant check permission %s\n",err.Error())
		return false
	}
	return false
}