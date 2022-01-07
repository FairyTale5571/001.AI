package bot

import (
	"001.AI/database"
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

func setVerifiedRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	role := i.ApplicationCommandData().Options[0].RoleValue(s,"")
	database.SetVerifiedRole(i.GuildID,role.ID)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Роль "+ role.Mention() +" установлена",
			Flags: 1 << 6,
		},
	})
}